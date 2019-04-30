package bamboohr_test

import (
	"testing"

	"gotest.tools/assert"

	bamboohr "github.com/AkihikoITOH/bamboohr.go"
)

func TestNewConfig(t *testing.T) {
	domain := "AkihikoITOH"
	key := "bamboohr-api-secret-key"

	config := bamboohr.NewConfig(domain, key)

	assert.Equal(t, config.APIDomain(), domain)
	assert.Equal(t, config.APIKey(), key)
}
