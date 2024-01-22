package data

import "fmt"

// struct kinda like class but without inheritance
type User struct {
	id int
	name string
	age int
}
func (u User) getUser() string {
	return fmt.Sprintf("User: %v",u.name)
}
func (u User) prityPrintUser() string {
	fmt.Printf("User: %v\n",u.name)
	return fmt.Sprintf("User: %v",u.name)
}

func init() {

	var u1 User
	u1 = User{name:"John", age: 44}
	u2 := User {33, "Jane",  33}


	u1.prityPrintUser()
	u2.prityPrintUser()

}