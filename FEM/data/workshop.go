package data
import "time"

type Workshop struct {
	Course Course
	Date time.Time
}

// embedding, cant add Course in constractor need a factory
type Workshop_Embedding struct {
	Course
	Date time.Time
}

func NewWorkshop_Embedding (name string, instructor Instructor, course_name string) Workshop_Embedding{
	w := Workshop_Embedding {}
	w.Name = name
	w.Instructor = instructor
	w.Course.Name = course_name
	return w

}