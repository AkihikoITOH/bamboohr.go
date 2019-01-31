package bamboohr_test

import (
	"testing"

	bamboohr "github.com/AkihikoITOH/bamboohr.go"
)

func TestNewEmployeeFromJSON(t *testing.T) {
	data := []byte(
		`{
      "id": "40941",
      "address1": "Choriner Str. 46",
      "age": "26",
      "birthday": "03-26",
      "city": "Berlin",
      "department": "Product Management",
      "displayName": "Carolin Adam",
      "dateOfBirth": "1992-03-26"
    }`)

	_, err := bamboohr.NewEmployeeFromJSON(data)

	if err != nil {
		t.Error(err.Error())
	}
}
