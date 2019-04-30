package bamboohr

import (
	"fmt"

	"github.com/AkihikoITOH/bamboohr.go/types"
)

// EstimateFutureTimeOffBalances fetches a list of balances of various time offs for the employee at the given date.
// For more details, refer to https://www.bamboohr.com/api/documentation/time_off.php
func (api *API) EstimateFutureTimeOffBalances(employeeID, endDate string) ([]*types.TimeOffBalance, error) {
	path := fmt.Sprintf("/employees/%s/time_off/calculator?end=%s", employeeID, endDate)

	data, err := api.get(path)
	if err != nil {
		return nil, err
	}

	return types.NewTimeOffBalancesFromJSON(data)
}
