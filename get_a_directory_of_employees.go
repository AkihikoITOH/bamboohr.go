package bamboohr

import "github.com/AkihikoITOH/bamboohr.go/types"

// GetADirectoryOfEmployees fetches a list of employees.
// For more details, refer to https://www.bamboohr.com/api/documentation/employees.php
func (api *API) GetADirectoryOfEmployees() (*types.EmployeeDirectory, error) {
	path := "/employees/directory"

	data, err := api.get(path)
	if err != nil {
		return nil, err
	}

	return types.NewEmployeeDirectoryFromJSON(data)
}
