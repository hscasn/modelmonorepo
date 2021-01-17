package apirequest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Parse the request body into a target struct
func Parse(r *http.Request, target interface{}) error {
	type reqBody struct {
		number  int64
		text    string
		boolean bool
		decimal float64
		date    time.Time
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("Could not parse request body: %w", err)
	}
	return nil
}
