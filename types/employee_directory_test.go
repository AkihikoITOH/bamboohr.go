package types_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/AkihikoITOH/bamboohr.go/types"
)

var dummyEmployeeDirectoryData = []byte(
	`{
			"fields": [
        {
          "id": "displayName",
          "type": "text",
          "name": "Display Name"
        },
        {
          "id": "firstName",
          "type": "text",
          "name": "First Name"
        },
        {
          "id": "lastName",
          "type": "text",
          "name": "Last Name"
        },
        {
          "id": "gender",
          "type": "text",
          "name": "Gender"
        },
        {
          "id": "jobTitle",
          "type": "list",
          "name": "Job Title"
        },
        {
          "id": "workPhone",
          "type": "text",
          "name": "Work Phone"
        },
        {
          "id": "workPhoneExtension",
          "type": "text",
          "name": "Work Extension"
        },
        {
          "id": "skypeUsername",
          "type": "text",
          "name": "Skype Username"
        },
        {
          "id": "facebook",
          "type": "text",
          "name": "Facebook URL"
        }
	    ],
	    "employees": [
        {
          "id": "123",
          "displayName": "John Doe",
          "firstName": "John",
          "lastName": "Doe",
          "gender": "Male",
          "jobTitle": "Customer Service Representative",
          "workPhone": "555-555-5555",
          "workPhoneExtension": null,
          "skypeUsername": "JohnDoe",
          "facebook": "JohnDoeFacebook"
        }
	    ]
		}`)

func TestNewEmployeeDirectoryFromJSON(t *testing.T) {
	_, err := types.NewEmployeeDirectoryFromJSON(dummyEmployeeDirectoryData)

	assert.Assert(t, err == nil)
}
