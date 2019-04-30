package types

import (
	"encoding/json"
)

// TimeOff represents a fixed time off(s).
type TimeOff struct {
	ID         int    `json:"id"`
	Type       string `json:"type"`
	EmployeeID int    `json:"employeeID"`
	Name       string `json:"name"`
	Start      string `json:"start"`
	End        string `json:"end"`
}

// NewTimeOffsFromJSON parses the given JSON (as byte sequence) and returns a new list of TimeOff objects.
func NewTimeOffsFromJSON(data []byte) ([]*TimeOff, error) {
	var timeoffs []*TimeOff
	err := json.Unmarshal(data, &timeoffs)
	return timeoffs, err
}
