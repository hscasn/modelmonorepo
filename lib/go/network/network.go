package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"cloud.google.com/go/compute/metadata"
	"github.com/sirupsen/logrus"
)

const cloudRunURLNeedle = "a.run.app"

// PostCloudRunCall is a POST call to a CloudRun service
func PostCloudRunCall(log *logrus.Entry, serviceURL, endpoint string, requestBody, responseBody interface{}) error {
	bodyb, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("error marshalling request body: %w", err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", serviceURL, endpoint), strings.NewReader(string(bodyb)))
	if err != nil {
		return fmt.Errorf("error creating request for post cloud run call: %w", err)
	}
	if strings.Contains(serviceURL, cloudRunURLNeedle) {
		// query the id_token with ?audience as the serviceURL
		tokenURL := fmt.Sprintf("/instance/service-accounts/default/identity?audience=%s", serviceURL)
		idToken, err := metadata.Get(tokenURL)
		if err != nil {
			return fmt.Errorf("metadata.Get: failed to query id_token: %w", err)
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", idToken))
	} else {
		log.Infof("detected non-CloudRun url: %s; authentication will be skipped", serviceURL)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error performing request for post cloud run call: %w", err)
	}
	bd, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body from post cloud run call: %w", err)
	}
	err = json.Unmarshal(bd, responseBody)
	if err != nil {
		return fmt.Errorf("error unmarshaling response body from post cloud run call: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("post cloud run call responded with status code: %d. body: %s", res.StatusCode, string(bd))
	}
	return nil
}
