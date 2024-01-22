package main


func getUser (id int) (string, error) {
	//get user
	
	var error error = nil

	var ok bool
	if ok {
		user := "Max"
		return user, nil
	} else {
		return "", error
	}
}

func init() {
	
	user, err := getUser(123)
	print(user, err)

}