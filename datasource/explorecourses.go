package datasource

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/rs/xid"
)

const endpoint = "https://explorecourses.stanford.edu/"
const dept_endpoint = endpoint + "?view=xml-20140630"

func ecRawSearch(search string, courses []Course) (int64, []Course, error) {
	courseEndpoint := endpoint + search
	var count int64

	collector := colly.NewCollector(
		colly.AllowedDomains("explorecourses.stanford.edu"),
		colly.CacheDir("./cache"),
	)

	collector.OnHTML("div#resultsHeader", func(element *colly.HTMLElement) {
		resultsSummary := element.ChildText("h1.resultsSummary")
		splitSummary := strings.SplitN(resultsSummary, "of ", 2)
		splitCount := strings.SplitN(splitSummary[1], " ", 2)[0]
		numResults, err := strconv.ParseInt(splitCount, 10, 64)
		if err != nil {
			panic(err)
		}
		count = numResults
	})

	collector.OnHTML("div.courseInfo", func(element *colly.HTMLElement) {
		var dept string
		var number string
		var terms string
		var units string
		var lastOffered string
		var instructors []Instructor
		var ugReqs string
		var schedules []ClassSchedule

		parsed := parseCourseNumber(element.ChildText("h2 > span.courseNumber"))
		if len(parsed) > 1 {
			dept = parsed[0]
			number = parsed[1]
		} else {
			dept = parsed[0]
		}
		title := TrimAllWhiteSpaces(element.ChildText("h2 > span.courseTitle"))
		description := TrimAllWhiteSpaces(element.ChildText("div.courseDescription"))

		// course attributes
		element.ForEach("div.courseAttributes", func(_ int, a *colly.HTMLElement) {
			if strings.Contains(a.Text, "Instructors") {
				a.ForEach("a", func(_ int, instructor *colly.HTMLElement) {
					profileUrl := instructor.Attr("href")
					name := TrimAllWhiteSpaces(instructor.Text)

					var isPI string
					splitTag := strings.Split(name, " (")
					if len(splitTag) > 1 {
						name = splitTag[0]
						if strings.Contains(splitTag[1], "PI") {
							isPI = "Y"
						} else {
							isPI = "N"
						}
					}
					instructors = append(instructors, Instructor{
						Name:       name,
						IsPI:       isPI,
						ProfileUrl: profileUrl,
					})
				})
			} else if strings.Contains(a.Text, "Last offered") {
				lastOffered = TrimAllWhiteSpaces(a.Text)
			} else {
				attributes := strings.Split(a.Text, "| ")
				m := make(map[string]string)
				// making a map of each of the attributes in the string
				for _, i := range attributes {
					parts := strings.SplitN(i, ": ", 2)
					if len(parts) > 1 {
						key := strings.TrimSpace(parts[0])
						value := strings.TrimSpace(parts[1])
						m[key] = value
					}
				}
				if t, found := m["Terms"]; found {
					terms = t
				}
				if u, found := m["Units"]; found {
					units = u
				}
				if ug, found := m["UG Reqs"]; found {
					ugReqs = ug
				}
			}
		})

		element.ForEach("div.sectionContainer", func(_ int, section *colly.HTMLElement) {
			var sectionType string
			var time string
			var instructors string
			var notes string

			sectionTerm := section.ChildText("h3.sectionContainerTerm")
			f := func(c rune) bool {
				return c == '\n' || c == '\t' || c == '|'
			}

			sectionDetails := strings.FieldsFunc(TrimWhiteSpaces(section.ChildText("li.sectionDetails")), f)
			for _, s := range sectionDetails {
				s = strings.TrimSpace(s)
				if s == "ISF" || s == "LEC" || s == "SEM" || s == "LBS" || s == "SEC" || s == "ACT" {
					sectionType = s
				} else if strings.Contains(s, "Notes:") {
					split := strings.Split(s, "Notes:")
					if len(split) > 1 {
						notes = split[1]
					}
				} else if strings.Contains(s, "Instructors:") {
					split := strings.Split(s, "Instructors:")
					if len(split) > 1 {
						instructors = split[1]
					}
				} else if strings.Contains(s, "AM") || strings.Contains(s, "PM") {
					time = s
				}
			}
			schedules = append(schedules, ClassSchedule{sectionTerm, sectionType, time, instructors, notes})

		})

		id := xid.New().String()

		courses = append(courses, Course{
			Id:                id,
			CourseDept:        dept,
			CourseNumber:      number,
			CourseTitle:       title,
			CourseDescription: description,
			Terms:             terms,
			Units:             units,
			LastOffered:       lastOffered,
			Instructors:       instructors,
			UGReqs:            ugReqs,
			Schedule:          schedules,
		})

	})

	collector.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	collector.Visit(courseEndpoint)

	return count, courses, nil
}

// TrimWhiteSpaces trims whitespace character such as double spaces
func TrimWhiteSpaces(str string) string {
	trimWhiteSpace := strings.NewReplacer("  ", "")
	return trimWhiteSpace.Replace(str)
}

// TrimWhiteSpaces trims all whitespace character such as line breaks or double spaces
func TrimAllWhiteSpaces(str string) string {
	trimWhiteSpace := strings.NewReplacer("\n", "", "\t", "", "  ", "")
	return trimWhiteSpace.Replace(str)
}

func parseCourseNumber(str string) []string {
	trimColon := strings.NewReplacer(":", "")
	trimmed := trimColon.Replace(str)
	splitNumber := strings.Split(trimmed, " ")
	return splitNumber
}

func getDepts() []string {
	var depts []string

	collector := colly.NewCollector(
		colly.AllowedDomains("explorecourses.stanford.edu"),
	)

	collector.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	collector.OnXML("//schools/school/*", func(e *colly.XMLElement) {
		name := e.ChildText("//@name")
		depts = append(depts, name)

	})

	collector.Visit(dept_endpoint)

	return depts
}

func ECGetCoursesByDepartment(pageSize, page uint, dept string) (count int64, courses []Course, err error) {
	search := "search?view=catalog&academicYear=&page=" + strconv.FormatUint(uint64(page), 10) + "&q=" + dept + "&filter-departmentcode-" + dept + "=on&filter-coursestatus-Active=on"
	count, thesecourses, err := ecRawSearch(search, courses)
	if err != nil {
		return
	}
	courses = append(courses, thesecourses...)
	return
}
