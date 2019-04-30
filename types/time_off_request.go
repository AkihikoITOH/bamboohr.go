package types

import (
	"encoding/json"
)

// TimeOffStatus is a part of TimeOffRequest.
type TimeOffStatus struct {
	LastChanged         string `json:"lastChanged"`
	LastChangedByUserID int    `json:"lastChangedByUserId,string"`
	Status              string `json:"status"`
}

// TimeOffType is a part of TimeOffRequest.
type TimeOffType struct {
	ID   int    `json:"id,string"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

// TimeOffAmount is a part of TimeOffRequest.
type TimeOffAmount struct {
	Unit   string  `json:"unit"`
	Amount float32 `json:"amount,string"`
}

// TimeOffActions is a part of TimeOffRequest.
type TimeOffActions struct {
	View    bool `json:"view"`
	Edit    bool `json:"edit"`
	Cancel  bool `json:"cancel"`
	Approve bool `json:"approve"`
	Deny    bool `json:"deny"`
	Bypass  bool `json:"bypass"`
}

// TimeOffRequest represents a time off request on BambooHR.
// For more details, refer to https://www.bamboohr.com/api/documentation/employees.php
type TimeOffRequest struct {
	ID         int               `json:"id,string"`
	EmployeeID int               `json:"employeeId,string"`
	Status     TimeOffStatus     `json:"status"`
	Name       string            `json:"name"`
	Start      string            `json:"start"`
	End        string            `json:"end"`
	Created    string            `json:"created"`
	Type       TimeOffType       `json:"type"`
	Amount     TimeOffAmount     `json:"amount"`
	Actions    TimeOffActions    `json:"actions"`
	Dates      map[string]string `json:"dates"`
	// Notes      map[string]string `json:"notes"` // notes field is either a map or list
}

// NewTimeOffRequestsFromJSON parses the given JSON (as byte sequence) and returns a new list of TimeOffRequest objects.
func NewTimeOffRequestsFromJSON(data []byte) ([]*TimeOffRequest, error) {
	requests := make([]*TimeOffRequest, 0)
	err := json.Unmarshal(data, &requests)
	return requests, err
}
