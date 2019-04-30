package types

import (
	"encoding/json"
)

// TimeOffBalance represents a balance of a certain type of time off at a certain point of time.
type TimeOffBalance struct {
	TimeOffType    int     `json:"timeOffType,string"`
	Name           string  `json:"name"`
	Units          string  `json:"units"`
	Balance        float32 `json:"balance,string"`
	PolicyType     string  `json:"policyType"`
	UsedYearToDate float32 `json:"usedYearToDate,string"`
	End            string  `json:"end"`
}

// NewTimeOffBalancesFromJSON parses the given JSON (as byte sequence) and returns a new TimeOffBalance.
func NewTimeOffBalancesFromJSON(data []byte) ([]*TimeOffBalance, error) {
	var balances []*TimeOffBalance
	err := json.Unmarshal(data, &balances)
	return balances, err
}
