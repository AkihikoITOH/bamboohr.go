package bamboohr_test

import (
	"fmt"
	"net/http"
	"testing"

	"gotest.tools/assert"

	bamboohr "github.com/AkihikoITOH/bamboohr.go"
)

var dummyTimeOffBalanceData = []byte(
	`[
    {
      "timeOffType": "1",
      "name": "Urlaub",
      "units": "days",
      "balance": "27.0",
      "end": "2019-03-30",
      "policyType": "manual",
      "usedYearToDate": "0.0"
    },
    {
      "timeOffType": "6",
      "name": "Krankmeldung ohne Attest",
      "units": "days",
      "balance": "-1.0",
      "end": "2019-03-30",
      "policyType": "manual",
      "usedYearToDate": "1.0"
    },
    {
      "timeOffType": "2",
      "name": "Krankheit",
      "units": "days",
      "balance": "0.0",
      "end": "2019-03-30",
      "policyType": "manual",
      "usedYearToDate": "0.0"
    },
    {
      "timeOffType": "4",
      "name": "Elternzeit",
      "units": "days",
      "balance": "0.0",
      "end": "2019-03-30",
      "policyType": "discretionary",
      "usedYearToDate": "0.0"
    },
    {
      "timeOffType": "14",
      "name": "Kind krank bezahlt",
      "units": "days",
      "balance": "0.0",
      "end": "2019-03-30",
      "policyType": "discretionary",
      "usedYearToDate": "0.0"
    },
    {
      "timeOffType": "15",
      "name": "Projektausgleich",
      "units": "days",
      "balance": "0.0",
      "end": "2019-03-30",
      "policyType": "discretionary",
      "usedYearToDate": "0.0"
    },
    {
      "timeOffType": "5",
      "name": "Sonderurlaub",
      "units": "days",
      "balance": "0.0",
      "end": "2019-03-30",
      "policyType": "manual",
      "usedYearToDate": "0.0"
    },
    {
      "timeOffType": "7",
      "name": "Homeoffice",
      "units": "days",
      "balance": "-8.0",
      "end": "2019-03-30",
      "policyType": "manual",
      "usedYearToDate": "0.0"
    },
    {
      "timeOffType": "16",
      "name": "Sabbatical",
      "units": "days",
      "balance": "0.0",
      "end": "2019-03-30",
      "policyType": "discretionary",
      "usedYearToDate": "0.0"
    },
    {
      "timeOffType": "13",
      "name": "Unbezahlter Urlaub",
      "units": "days",
      "balance": "0.0",
      "end": "2019-03-30",
      "policyType": "discretionary",
      "usedYearToDate": "0.0"
    },
    {
      "timeOffType": "12",
      "name": "Bezahlte Freistellung",
      "units": "days",
      "balance": "0.0",
      "end": "2019-03-30",
      "policyType": "discretionary",
      "usedYearToDate": "0.0"
    },
    {
      "timeOffType": "9",
      "name": "Fehltag unentschuldigt",
      "units": "days",
      "balance": "0.0",
      "end": "2019-03-30",
      "policyType": "discretionary",
      "usedYearToDate": "0.0"
    },
    {
      "timeOffType": "8",
      "name": "Kind krank - unbezahlt",
      "units": "days",
      "balance": "0.0",
      "end": "2019-03-30",
      "policyType": "discretionary",
      "usedYearToDate": "0.0"
    },
    {
      "timeOffType": "11",
      "name": "Krank Arbeitsunfall",
      "units": "days",
      "balance": "0.0",
      "end": "2019-03-30",
      "policyType": "discretionary",
      "usedYearToDate": "0.0"
    },
    {
      "timeOffType": "3",
      "name": "Mutterschutz",
      "units": "days",
      "balance": "0.0",
      "end": "2019-03-30",
      "policyType": "discretionary",
      "usedYearToDate": "0.0"
    },
    {
      "timeOffType": "10",
      "name": "Urlaub ausbezahlt",
      "units": "days",
      "balance": "0.0",
      "end": "2019-03-30",
      "policyType": "discretionary",
      "usedYearToDate": "0.0"
    }
  ]`)

func Example_estimateFutureTimeOffBalances() {
	conf := bamboohr.NewConfig("api-domain", "api-key")
	api := bamboohr.NewAPI(conf)
	api.EstimateFutureTimeOffBalances("12345", "2019-03-30")
}

func TestEstimateFutureTimeOffBalances(t *testing.T) {
	client := NewMockClient(dummyTimeOffBalanceData)
	api := &bamboohr.API{Config: mockAPIConfig, HTTPClient: client}
	balances, err := api.EstimateFutureTimeOffBalances("12345", "2019-03-30")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Check request object
	assert.Equal(t, client.Request.Method, http.MethodGet)
	assert.Equal(t, client.Request.URL.String(), "https://api.bamboohr.com/api/gateway.php/test/v1/employees/12345/time_off/calculator?end=2019-03-30")
	assert.Equal(t, client.Request.Body, nil)
	SharedRequestTests(t, client.Request)

	assert.Equal(t, balances[0].Name, "Urlaub")
	assert.Equal(t, balances[0].Balance, float32(27.0))
}
