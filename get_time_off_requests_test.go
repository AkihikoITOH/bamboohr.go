package bamboohr_test

import (
	"fmt"
	"net/http"
	"testing"

	"gotest.tools/assert"

	bamboohr "github.com/AkihikoITOH/bamboohr.go"
)

var timeOffRequestsData = []byte(
	`[
			{
				"id": "1111",
				"employeeId": "1234",
				"status": {
					"lastChanged": "2017-11-26",
					"lastChangedByUserId": "2222",
					"status": "approved"
				},
				"name": "John Doe",
				"start": "2018-01-01",
				"end": "2018-01-05",
				"created": "2017-11-23",
				"type": {
					"id": "1",
					"name": "Vacation",
					"icon": "time-off-luggage"
				},
				"amount": {
					"unit": "days",
					"amount": "4"
				},
				"actions": {
					"view": true,
					"edit": false,
					"cancel": false,
					"approve": false,
					"deny": false,
					"bypass": false
				},
				"dates": {
					"2018-01-02": "1",
					"2018-01-03": "1",
					"2018-01-04": "1",
					"2018-01-05": "1"
				},
				"notes": {
					"employee": "Ok"
				}
			},
			{
				"id": "3333",
				"employeeId": "9876",
				"status": {
					"lastChanged": "2018-01-19",
					"lastChangedByUserId": "4444",
					"status": "approved"
				},
				"name": "Jessy Doe",
				"start": "2018-01-01",
				"end": "2018-01-21",
				"created": "2017-08-24",
				"type": {
					"id": "1",
					"name": "Vacation",
					"icon": "time-off-luggage"
				},
				"amount": {
					"unit": "days",
					"amount": "11"
				},
				"actions": {
					"view": true,
					"edit": false,
					"cancel": false,
					"approve": false,
					"deny": false,
					"bypass": false
				},
				"dates": {
					"2018-01-02": "1",
					"2018-01-03": "1",
					"2018-01-04": "1",
					"2018-01-05": "1",
					"2018-01-08": "1",
					"2018-01-09": "1",
					"2018-01-15": "1",
					"2018-01-16": "1",
					"2018-01-17": "1",
					"2018-01-18": "1",
					"2018-01-19": "1"
				},
				"notes": []
		  }
		]`)

func Example_getTimeOffRequests() {
	conf := bamboohr.NewConfig("api-domain", "api-key")
	api := bamboohr.NewAPI(conf)
	filters := map[string]string{
		"start":  "2018-01-01",
		"end":    "2018-01-02",
		"status": "approved",
	}
	api.GetTimeOffRequests(filters)
}

func TestGetTimeOffRequests(t *testing.T) {
	client := NewMockClient(timeOffRequestsData)
	api := &bamboohr.API{Config: mockAPIConfig, HTTPClient: client}
	filters := map[string]string{
		"start":  "2018-01-01",
		"end":    "2018-01-02",
		"type":   "1",
		"status": "approved",
	}
	requests, err := api.GetTimeOffRequests(filters)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	assert.Equal(t, client.Request.URL.String(), "https://api.bamboohr.com/api/gateway.php/test/v1/time_off/requests?start=2018-01-01&end=2018-01-02&type=1&status=approved")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	// Check employee object
	assert.Equal(t, len(requests), 2)
	assert.Equal(t, requests[0].ID, 1111)
}
