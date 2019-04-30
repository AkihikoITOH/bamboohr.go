package bamboohr

import (
	"fmt"
	"strings"

	"github.com/AkihikoITOH/bamboohr.go/types"
)

// GetAnEmployee fetches an employee with the given employee number from the BambooHR API.
// For more details, refer to https://www.bamboohr.com/api/documentation/employees.php
func (api *API) GetAnEmployee(number string, fieldList []string) (*types.Employee, error) {
	path := fmt.Sprintf("/employees/%s", number)
	if len(fieldList) > 0 {
		path += fmt.Sprintf("?fields=%s", strings.Join(fieldList, ","))
	}

	data, err := api.get(path)
	if err != nil {
		return nil, err
	}

	return types.NewEmployeeFromJSON(data)
}
