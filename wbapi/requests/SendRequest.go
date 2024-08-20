package requests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gasuhwbab/wbbot/wbapi"
)

func sendRequest(wbc *wbapi.WbClient, req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", wbc.ApiToken))
	resp, err := wbc.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error, status code: %d", resp.StatusCode)
	}
	if err = json.NewDecoder(resp.Body).Decode(v); err != nil {
		return err
	}
	return nil
}
