package matomo

import (
	"fmt"
	"os"
)

type Configuration struct {
	Domain string
	SiteID int64 // if not provided, will be required in the call
	Rec    int64 // must always be set to 1
}

var config *Configuration

func Setup() {
	if config != nil {
		return
	}
	config = &Configuration{}
	config.Domain = envHelper("MATOMO_DOMAIN", "")
	if config.Domain == "" {
		// TODO: convert to logger
		fmt.Printf("\n----------------------------\nERROR: MATOMO_DOMAIN was not set, so events will not be tracked\n----------------------------\n")
	}

	config.Rec = 1

}

func envHelper(key, defaultValue string) string {
	found := os.Getenv(key)
	if found == "" {
		found = defaultValue
	}
	return found
}
