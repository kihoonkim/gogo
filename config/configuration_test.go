package config

import (
	"testing"
	"os"
	"fmt"
)

func TestGetConfiguration(t *testing.T) {
	// mock env variable
	os.Setenv("PORT", "1234")

	configuration := GetConfiguration()

	if configuration.Port != 1234 {
		fmt.Errorf("Expected %d, but Actual %d", 1234, configuration.Port)
	}
}