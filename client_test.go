package matomo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	// if the domain and site id aren't set, then we
	// will skip this test
	Setup()
	if config.Domain == "" || config.SiteID == "" {
		// TODO: convert to a logger
		fmt.Println("No domain or no site id, skipping client test")
		t.Skip()
	}
	// send up some data; currently there is no way to validate the data was saved
	// using the SDK in an automated way (no data retrieval) so we rely on the
	// status code and lack of an error
	err := Send(&testAllParams)
	assert.Nil(t, err)
}
