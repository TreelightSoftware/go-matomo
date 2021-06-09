package matomo

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

// Send is a helper function that reads the siteID from the configuration file rather than requiring the user
// to provide it.
func Send(params *Parameters) error {
	if config.Domain == "" || config.SiteID == "" {
		return errors.New("either domain or site id are not provided")
	}
	return SendToSite(config.SiteID, params)
}

// SendToSite sends the parameters to Matomo instance. Matomo wants all of the data in the query string, regardless of whether
// GET or POST is used.
func SendToSite(siteID string, params *Parameters) error {
	if config.Domain == "" {
		return errors.New("the domain was not provided")
	}
	data := params.encode()
	// set the required parameters
	data["idsite"] = siteID
	data["rec"] = config.Rec

	client := resty.New()
	resp, err := client.R().SetQueryParams(data).Get(config.Domain + "/matomo.php")
	if err != nil {
		return err
	}
	statusCode := resp.StatusCode()
	if statusCode != http.StatusOK && statusCode != http.StatusNoContent {
		return fmt.Errorf("invalid status code returned: %d, body was: %+v", statusCode, string(resp.Body()))
	}
	return nil
}
