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
      "name": "Vacation",
      "units": "days",
      "balance": "27.0",
      "end": "2019-03-30",
      "policyType": "manual",
      "usedYearToDate": "0.0"
    },
    {
      "timeOffType": "6",
      "name": "Ditched",
      "units": "days",
      "balance": "-1.0",
      "end": "2019-03-30",
      "policyType": "manual",
      "usedYearToDate": "1.0"
    },
    {
      "timeOffType": "2",
      "name": "Sick leave",
      "units": "days",
      "balance": "0.0",
      "end": "2019-03-30",
      "policyType": "manual",
      "usedYearToDate": "0.0"
    },
    {
      "timeOffType": "4",
      "name": "Parental leave",
      "units": "days",
      "balance": "0.0",
      "end": "2019-03-30",
      "policyType": "discretionary",
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
    }
  ]`)

func TestNewTimeOffBalancesFromJSON(t *testing.T) {
	_, err := types.NewTimeOffBalancesFromJSON(dummyTimeOffBalanceData)

	assert.Assert(t, err == nil)
}
