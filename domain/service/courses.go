package service

type CreateCoursesReq struct {
	Subject string `json:"subject"`
}

type CreateCourseDetailReq struct {
	CoursesUUID string `json:"courses_uuid"`
	SubCourses  string `json:"sub_courses"`
}
