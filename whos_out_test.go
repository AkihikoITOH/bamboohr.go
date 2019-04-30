package bamboohr_test

import (
	"fmt"
	"net/http"
	"testing"

	"gotest.tools/assert"

	bamboohr "github.com/AkihikoITOH/bamboohr.go"
)

var dummyTimeOffsData = []byte(
	`[
    {
      "id": 111,
      "type": "timeOff",
      "employeeId": 54321,
      "name": "Jessy Doe",
      "start": "2018-02-01",
      "end": "2019-01-31"
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

func Example_whosOut() {
	conf := bamboohr.NewConfig("api-domain", "api-key")
	api := bamboohr.NewAPI(conf)
	api.WhosOut("2018-01-01", "2019-03-31")
}

func TestWhosOut(t *testing.T) {
	client := NewMockClient(dummyTimeOffsData)
	api := &bamboohr.API{Config: mockAPIConfig, HTTPClient: client}
	timeoffs, err := api.WhosOut("", "2019-03-31")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	assert.Equal(t, client.Request.URL.String(), "https://api.bamboohr.com/api/gateway.php/test/v1/time_off/whos_out?end=2019-03-31")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, len(timeoffs), 3)
	assert.Equal(t, timeoffs[0].Name, "Jessy Doe")
	assert.Equal(t, timeoffs[0].ID, 111)
}
