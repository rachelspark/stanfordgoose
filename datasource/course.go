package datasource

// JSON object that specifies a course row.
type Course struct {

	// CourseNumber is the course number (e.g. "ME101").
	CourseNumber string `json:"courseNumber"`

	CourseTitle string `json:"courseTitle"`

	CourseDescription string `json:"courseDescription"`

	Terms string `json:"terms"`

	Units string `json:"units"`

	LastOffered string `json:"lastOffered"`

	Instructors []Instructor `json:"instructors"`

	UGReqs string `json:"ugReqs"`

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
