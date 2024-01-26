package api
import (
	"github.com/dimitrius-ion/go_basics/crypto/datatypes"
	"fmt"
	"net/http"
	"strings"
	"io"
)

const apiURL = "https://cex.io/api/ticker/%s/USD"

func GetRate (currency string) (*datatypes.Rate, error) {
	upper_currency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiURL, upper_currency))
	
	// handle error
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", res.StatusCode)
	}
	defer res.Body.Close()
	// read body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	json := string(bodyBytes)
	fmt.Println(json)
	rate := datatypes.Rate{Currency: upper_currency, Price: 0.0}
	return &rate, nil
	
}