package types

import (
	"encoding/json"
)

// Employee contains all the fields of employees provided by the API.
// (https://www.bamboohr.com/api/documentation/employees.php)
type Employee struct {
	Address1                string  `json:"address1"`                    // The employee's first address line.
	Address2                string  `json:"address2"`                    // The employee's second address line.
	Age                     int     `json:"age,string"`                  // The employee's age. To change age, update dateOfBirth field.
	BestEmail               string  `json:"bestEmail"`                   // The employee's work email if set, otherwise their home email.
	Birthday                string  `json:"birthday"`                    // The employee's month and day of birth. To change birthday, update dateOfBirth field.
	City                    string  `json:"city"`                        // The employee's city.
	Country                 string  `json:"country"`                     // The employee's country.
	DateOfBirth             string  `json:"dateOfBirth"`                 // The date the employee was born.
	Department              string  `json:"department"`                  // The employee's CURRENT department.
	Division                string  `json:"division"`                    // The employee's CURRENT division.
	Eeo                     string  `json:"eeo"`                         // The employee's EEO job category. These are defined by the U.S. Equal Employment Opportunity Commission.
	EmployeeNumber          string  `json:"employeeNumber"`              // Employee number (assigned by your company).
	EmploymentHistoryStatus string  `json:"employmentHistoryStatus"`     // The employee's CURRENT employment status. Options are customized by account. Read-only starting with version 1.1; update using the employmentStatus table.
	Ethnicity               string  `json:"ethnicity"`                   // The employee's ethnicity.
	Exempt                  string  `json:"exempt"`                      // The FLSA Overtime status (Exempt or Non-exempt).
	FirstName               string  `json:"firstName"`                   // The employee's first name.
	FullName1               string  `json:"fullName1"`                   // The employee's first and last name. (e.g., John Doe). Read only.
	FullName2               string  `json:"fullName2"`                   // The employee's last and first name. (e.g., Doe, John). Read only.
	FullName3               string  `json:"fullName3"`                   // The employee's full name and their preferred name. (e.g., Doe, John Quentin (JDog)). Read only.
	FullName4               string  `json:"fullName4"`                   // The employee's full name without their preferred name, last name first. (e.g., Doe, John Quentin). Read only.
	FullName5               string  `json:"fullName5"`                   // The employee's full name without their preferred name, first name first. (e.g., John Quentin Doe). Read only.
	DisplayName             string  `json:"displayName"`                 // The employee's name displayed in a format configured by the user. Read only.
	Gender                  string  `json:"gender"`                      // The employee's gender (Male or Female).
	HireDate                string  `json:"hireDate"`                    // The date the employee was hired.
	OriginalHireDate        string  `json:"originalHireDate"`            // The date the employee was originally hired. Available starting with version 1.1.
	HomeEmail               string  `json:"homeEmail"`                   // The employee's home mmail address.
	HomePhone               string  `json:"homePhone"`                   // The employee's home phone number.
	ID                      int     `json:"id,string"`                   // The employee ID automatically assigned by BambooHR. Read only.
	JobTitle                string  `json:"jobTitle"`                    // The CURRENT value of the employee's job title, updating this field will create a new row in position history.
	LastChanged             string  `json:"lastChanged"`                 // The date and time that the employee record was last changed.
	LastName                string  `json:"lastName"`                    // The employee's last name.
	Location                string  `json:"location"`                    // The employee's CURRENT location.
	MaritalStatus           string  `json:"maritalStatus"`               // The employee's marital status (Single, Married, or Domestic Partnership).
	MiddleName              string  `json:"middleName"`                  // The employee's middle name.
	MobilePhone             string  `json:"mobilePhone"`                 // The employee's mobile phone number.
	PayChangeReason         string  `json:"payChangeReason"`             // The reason for the employee's last pay rate change.
	PayGroup                string  `json:"payGroup"`                    // The custom pay group that the employee belongs to.
	PayGroupID              int     `json:"payGroupId,string"`           // The ID value corresponding to the pay group that an employee belongs to.
	PayRate                 float32 `json:"payRate"`                     // The employee's CURRENT pay rate (e.g., $8.25).
	PayRateEffectiveDate    string  `json:"payRateEffectiveDate"`        // The day the most recent change was made.
	PayType                 string  `json:"payType"`                     // The employee's CURRENT pay type. ie: "hourly","salary","commission","exception hourly","monthly","weekly","piece rate","contract","daily","pro rata".
	PayPer                  string  `json:"payPer"`                      // The employee's CURRENT pay per. ie: "Hour", "Day", "Week", "Month", "Quarter", "Year".
	PaidPer                 string  `json:"paidPer"`                     // The employee's CURRENT pay per. ie: "Hour", "Day", "Week", "Month", "Quarter", "Year".
	PaySchedule             string  `json:"paySchedule"`                 // The employee's CURRENT pay schedule.
	PayScheduleID           int     `json:"payScheduleId,string"`        // The ID value corresponding to the pay schedule that an employee belongs to.
	PayFrequency            string  `json:"payFrequency"`                // The employee's CURRENT pay frequency. ie: "Weekly", "Every other week", "Twice a month", "Monthly", "Quarterly", "Twice a year", or "Yearly"
	IncludeInPayroll        bool    `json:"includeInPayroll,string"`     // Should employee be included in payroll (Yes or No)
	TimeTrackingEnabled     bool    `json:"timeTrackingEnabled,sring"`   // Should time tracking be enabled for the employee (Yes or No)
	PreferredName           string  `json:"preferredName"`               // The employee's preferred name.
	Ssn                     string  `json:"ssn"`                         // The employee's Social Security number.
	Sin                     string  `json:"sin"`                         // The employee's Canadian Social Insurance Number.
	State                   string  `json:"state"`                       // The employee's state/province.
	StateCode               string  `json:"stateCode"`                   // The 2 character abbreviation for the employee's state (US only). Read only.
	Status                  string  `json:"status"`                      // The employee's employee status (Active or Inactive).
	Supervisor              string  `json:"supervisor"`                  // The employee’s CURRENT supervisor. Read only.
	SupervisorID            string  `json:"supervisorId"`                // The 'employeeNumber' of the employee's CURRENT supervisor. Read only.
	SupervisorEId           int     `json:"supervisorEId,string"`        // The ID of the employee's CURRENT supervisor. Read only.
	TerminationDate         string  `json:"terminationDate"`             // The date the employee was terminated. Read-only starting with version 1.1; update using the employmentStatus table.
	WorkEmail               string  `json:"workEmail"`                   // The employee's work email address.
	WorkPhone               string  `json:"workPhone"`                   // The employee's work phone number, without extension.
	WorkPhonePlusExtension  string  `json:"workPhonePlusExtension"`      // The employee's work phone and extension. Read only.
	WorkPhoneExtension      string  `json:"workPhoneExtension"`          // The employee's work phone extension (if any).
	Zipcode                 string  `json:"zipcode"`                     // The employee's ZIP code.
	IsPhotoUploaded         string  `json:"isPhotoUploaded"`             // Whether a photo has been uploaded for the employee. Read only.
	AcaStatus               string  `json:"acaStatus"`                   // The employee's ACA (Affordable Care Act) Full-Time status. Options are yes, no, non-assessment
	StandardHoursPerWeek    int     `json:"standardHoursPerWeek,string"` // The number of hours the employee works in a standard week.
	BonusDate               string  `json:"bonusDate"`                   // The date of the last bonus.
	BonusAmount             float32 `json:"bonusAmount,string"`          // The amount of the most recent bonus.
	BonusReason             string  `json:"bonusReason"`                 // The reason for the most recent bonus.
	BonusComment            string  `json:"bonusComment"`                // Comment about the most recent bonus.
	CommissionDate          string  `json:"commissionDate"`              // The date of the last commission.
	CommissionAmount        float32 `json:"commissionAmount,string"`     // The amount of the most recent commission.
	CommissionComment       string  `json:"commissionComment"`           // Comment about the most recent commission.
	Nin                     string  `json:"nin"`                         // The employee's nin number
	NationalID              string  `json:"nationalID"`                  // The employee's national ID number
	Nationality             string  `json:"nationality"`                 // The employee's nationality
}

// NewEmployeeFromJSON parses the given JSON (as byte sequence) and returns a new Employee.
func NewEmployeeFromJSON(data []byte) (*Employee, error) {
	var employee Employee
	err := json.Unmarshal(data, &employee)
	return &employee, err
}
