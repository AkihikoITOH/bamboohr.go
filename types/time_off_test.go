package types_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/AkihikoITOH/bamboohr.go/types"
)

var dummyTimeOffsData = []byte(
	`[
    {
      "id": 111,
      "type": "timeOff",
      "employeeId": 54321,
      "name": "Jessy Doe",
      "start": "2018-02-01",
      "end": "2021-01-31"
    },
    {
      "id": 222,
      "type": "timeOff",
      "employeeId": 12345,
      "name": "John Doe",
      "start": "2018-11-07",
      "end": "2019-03-31"
    },
    {
      "id": 123,
      "type": "holiday",
      "start": "2019-03-08",
      "end": "2019-03-08"
    }
  ]`)

func TestNewTimeOffsFromJSON(t *testing.T) {
	_, err := types.NewTimeOffsFromJSON(dummyTimeOffsData)

	assert.Assert(t, err == nil)
}
