package types_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/AkihikoITOH/bamboohr.go/types"
)

var dummyEmployeeData = []byte(
	`{
			"id": "99999",
			"address1": "Am Sandtorkai 37",
			"age": "27",
			"birthday": "03-07",
			"city": "Hamburg",
			"department": "Development",
			"displayName": "John Doe",
			"dateOfBirth": "2019-03-07",
			"workEmail": "test@example.com",
			"mobilePhone": "+491749999999",
			"location": "Hamburg",
			"lastChanged": "2019-02-26T20:46:32+00:00",
			"isPhotoUploaded": "true"
		}`)

func TestNewEmployeeFromJSON(t *testing.T) {
	_, err := types.NewEmployeeFromJSON(dummyEmployeeData)

	assert.Assert(t, err == nil)
}
