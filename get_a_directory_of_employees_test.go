package bamboohr_test

import (
	"fmt"
	"net/http"
	"testing"

	"gotest.tools/assert"

	bamboohr "github.com/AkihikoITOH/bamboohr.go"
)

var employeeDirectoryData = []byte(
	`{
			"employees": [
				{
					"id": "99999",
					"firstName": "John",
					"lastName": "Doe"
				}
			]
		}`)

func Example_getADirectoryOfEmployees() {
	conf := bamboohr.NewConfig("api-domain", "api-key")
	api := bamboohr.NewAPI(conf)
	api.GetADirectoryOfEmployees()
}

func TestGetADirectoryOfEmployees(t *testing.T) {
	client := NewMockClient(employeeDirectoryData)
	api := &bamboohr.API{Config: mockAPIConfig, HTTPClient: client}
	employees, err := api.GetADirectoryOfEmployees()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	assert.Equal(t, client.Request.URL.String(), "https://api.bamboohr.com/api/gateway.php/test/v1/employees/directory")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	// Check employee object
	assert.Equal(t, employees.Employees[0].ID, 99999)
	assert.Equal(t, employees.Employees[0].FirstName, "John")
	assert.Equal(t, employees.Employees[0].LastName, "Doe")
}
