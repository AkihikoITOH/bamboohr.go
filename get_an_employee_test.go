package bamboohr_test

import (
	"fmt"
	"net/http"
	"testing"

	"gotest.tools/assert"

	bamboohr "github.com/AkihikoITOH/bamboohr.go"
)

var employeeData = []byte(
	`{
			"id": "99999",
			"firstName": "John",
			"lastName": "Doe"
		}`)

func Example_getAnEmployee() {
	conf := bamboohr.NewConfig("api-domain", "api-key")
	api := bamboohr.NewAPI(conf)
	api.GetAnEmployee("99999", []string{"firstName", "lastName"})
}

func TestGetAnEmployee(t *testing.T) {
	client := NewMockClient(employeeData)
	api := &bamboohr.API{Config: mockAPIConfig, HTTPClient: client}
	employee, err := api.GetAnEmployee("99999", []string{"firstName", "lastName"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	assert.Equal(t, client.Request.URL.String(), "https://api.bamboohr.com/api/gateway.php/test/v1/employees/99999?fields=firstName,lastName")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	// Check employee object
	assert.Equal(t, employee.ID, 99999)
	assert.Equal(t, employee.FirstName, "John")
	assert.Equal(t, employee.LastName, "Doe")
}
