package datasource

// JSON object that specifies a course row.
type Course struct {
	Id string `json:"id"`

	// Dept is the abbreviated course department code (ME, CS, AA)
	Dept string `json:"dept"`

	DeptLongname string `json:"deptLongname"`

	// CourseNumber is the course number (e.g. 101 in "ME101").
	CourseNumber uint32 `json:"courseNumber"`

	DeptAndNumber string `json:"deptAndNumber"`

	CourseTitle string `json:"courseTitle"`

	CourseDescription string `json:"courseDescription"`

	Level string `json:"level"`

	Terms []string `json:"terms"`

	Units string `json:"units"`

	LastOffered string `json:"lastOffered"`

	Instructors []Instructor `json:"instructors"`

	UGReqs []string `json:"ugReqs"`

	Schedule []ClassSchedule `json:"schedule"`
}

type Instructor struct {
	// Name is the name of the instructor
	Name string `json:"name"`

	IsPI string `json:"isPI"`

	// ProfileURl is the link to the instructor's profile
	ProfileUrl string `json:"profileUrl"`
}

type ClassSchedule struct {
	Term string `json:"term"`

	ClassType string `json:"classType"`

	ClassTime string `json:"classTime"`

	ClassInstructors string `json:"termInstructors"`

	Notes string `json:"notes"`
}
