package datasource

import (
	"log"
	"sort"
	"sync"

	"github.com/schollz/progressbar/v3"
	"golang.org/x/exp/slices"
)

// Download a paginated set of courses over the network, with specified concurrency.
func paginatedDownload(
	callback func(uint, uint, string) (int64, []Course, error),
	pageSize uint,
	concurrency uint,
	dept string,
) []Course {
	totalCount, _, err := callback(pageSize, 1, dept)
	if err != nil {
		log.Fatalf("failed to get courses: %v", err)
	}

	var mu sync.Mutex // Protects the courses list.
	var courses []Course
	bar := progressbar.Default(totalCount)

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, concurrency)
	for i := uint(0); i < uint(totalCount); i += pageSize {
		i := i
		semaphore <- struct{}{}
		wg.Add(1)
		go func() {
			defer func() { <-semaphore }()
			defer wg.Done()
			_, data, err := callback(pageSize, i/pageSize, dept)
			if err != nil {
				log.Fatalf("failed to get courses: %v", err)
			}
			mu.Lock()
			defer mu.Unlock()
			courses = append(courses, data...)
			bar.Add(len(data))
		}()
	}
	wg.Wait()

	initialLen := len(courses)
	sort.Slice(courses, func(i, j int) bool {
		return courses[i].Id < courses[j].Id
	})
	courses = slices.CompactFunc(courses, func(a, b Course) bool {
		return a.Id == b.Id
	})

	log.Printf("read %v out of %v total courses (%v before compaction)",
		len(courses), totalCount, initialLen)
	return courses
}

// Download course data from Explore courses
func DownloadCourses() []Course {
	log.Printf("starting to download from Explore Courses")

	depts := getDepts()
	var courses []Course

	for name := range depts {
		courses = append(courses, paginatedDownload(ECGetCoursesByDepartment, 10, 32, name)...)
	}

	return courses
}
