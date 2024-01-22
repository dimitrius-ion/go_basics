package data
import "fmt"

// alias of int
type integer = int 

type arrayOfStrings = []string

type json = map[string]string

// new type 
// type json map[string]string


type distance float64          // miles
type distanceKm float64

func (miles distance) toKmt () distanceKm {
	return distanceKm(miles * 1.6)
}
func (km distanceKm) toMiles () distance {
	return distance(km / 1.6)
}

func init() {
	d := distance(54.6)
	km := d.toKmt()
	km.toMiles()
	fmt.Printf("Distance: %.2f miles === %.2f km",d, km)
}