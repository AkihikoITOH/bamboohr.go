package bamboohr

import (
	"fmt"
	"strings"

	"github.com/AkihikoITOH/bamboohr.go/types"
)

// WhosOut fetches a list of time offs, including holidays between the given dates.
// For more details, refer to https://www.bamboohr.com/api/documentation/time_off.php
func (api *API) WhosOut(start, end string) ([]*types.TimeOff, error) {
	path := "/time_off/whos_out"

	params := make([]string, 0)
	if len(start) > 0 {
		params = append(params, fmt.Sprintf("start=%s", start))
	}
	if len(end) > 0 {
		params = append(params, fmt.Sprintf("end=%s", end))
	}
	if len(params) > 0 {
		path += fmt.Sprintf("?%s", strings.Join(params, "&"))
	}

	data, err := api.get(path)
	if err != nil {
		return nil, err
	}

	return types.NewTimeOffsFromJSON(data)
}
