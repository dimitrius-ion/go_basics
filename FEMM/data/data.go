package data


type Things struct {
	Title  string
	Description string
	URL string
	Big bool
}

func GetAll() []Things {
	return listOfThings
}

func AddThing(t Things) {

	listOfThings = append(listOfThings, t)
}

var listOfThings = []Things{
	{
		Title: "Spoon", 
		Description: "an implement consisting of a small, shallow oval or round bowl on a long handle, used for eating, stirring, and serving food.",
		URL: "https://www.spoon.com",
		Big: false,
	},
	{
		Title: "Fork", 
		Description: "an implement with two or more prongs used for lifting food to the mouth or holding it when cutting.",
		URL: "https://www.fork.com",
		Big: true,
	},

}