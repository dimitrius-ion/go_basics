package data


type Things struct {
	Title  string
	Description string
	URL string
}

func GetAll() []Things {
	return listOfThings
}

var listOfThings = []Things{
	{
		Title: "Spoon", 
		Description: "an implement consisting of a small, shallow oval or round bowl on a long handle, used for eating, stirring, and serving food.",
		URL: "https://www.spoon.com",
	},
	{
		Title: "Fork", 
		Description: "an implement with two or more prongs used for lifting food to the mouth or holding it when cutting.",
		URL: "https://www.fork.com",
	},

}