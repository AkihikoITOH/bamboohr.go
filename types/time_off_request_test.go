package types_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/AkihikoITOH/bamboohr.go/types"
)

var dummyTimeOffRequestsData = []byte(
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

func TestNewTimeOffRequestsFromJSON(t *testing.T) {
	_, err := types.NewTimeOffRequestsFromJSON(dummyTimeOffRequestsData)

	assert.Assert(t, err == nil)
}
