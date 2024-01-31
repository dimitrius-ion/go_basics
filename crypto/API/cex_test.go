package api_test
import (
	"testing"
)

// _test.go will not compile


func TestAPICall(t *testing.T) {
	// Test the API call
	// ...
     _,err := api.GetRate("")
     if err == nil {
		 t.Errorf("Error: %v", err)
	 }
}