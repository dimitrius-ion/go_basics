package data

import "fmt"

type Instructor struct {
	Id                int
	FirstName     string
	LastName  string
	Score     float64
}

// factory function
func NewInstructor (FirstName string, LastName string) Instructor {
	return Instructor{FirstName: FirstName, LastName: LastName}
}

func (i Instructor) Print() string {
	return fmt.Sprintf("%d: %s %s", i.Id, i.FirstName, i.LastName)
}
