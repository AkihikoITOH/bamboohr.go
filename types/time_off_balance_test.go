package types_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/AkihikoITOH/bamboohr.go/types"
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

func TestNewTimeOffBalancesFromJSON(t *testing.T) {
	_, err := types.NewTimeOffBalancesFromJSON(dummyTimeOffBalanceData)

	assert.Assert(t, err == nil)
}
