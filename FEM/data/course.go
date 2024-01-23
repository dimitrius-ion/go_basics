package data
import "fmt"

type Duration float64

type Course struct {
	Id          int
	Name        string
	Description string
	Instructor  Instructor
}

func (c Course) SingUp() bool {
	fmt.Println("SingUp for", c.Name)
	return true
}

func (c Course) String () string {
	return fmt.Sprintf("-----%v ------\n  %d (%v %v)", c.Name,  c.Id, c.Instructor.FirstName , c.Instructor.LastName )
}
