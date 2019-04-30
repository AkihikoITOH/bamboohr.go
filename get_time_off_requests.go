package bamboohr

import (
	"fmt"
	"strings"

	"github.com/AkihikoITOH/bamboohr.go/types"
)

// GetTimeOffRequests fetches time off requests from the BambooHR API.
// For more details, refer to https://www.bamboohr.com/api/documentation/time_off.php
func (api *API) GetTimeOffRequests(filters map[string]string) ([]*types.TimeOffRequest, error) {
	path := "/time_off/requests"
	if filters != nil && len(filters) > 0 {
		params := []string{}
		for k, v := range filters {
			params = append(params, fmt.Sprintf("%s=%s", k, v))
		}
		path += fmt.Sprintf("?%s", strings.Join(params, "&"))
	}

	data, err := api.get(path)
	if err != nil {
		return nil, err
	}

	return types.NewTimeOffRequestsFromJSON(data)
}
