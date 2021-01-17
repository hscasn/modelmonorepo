package settings

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/testingtools"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	router := chi.NewRouter()
	New(router)
	s := httptest.NewServer(router)
	defer s.Close()

	// Settings
	res, _, err := testingtools.HTTPRequest(s.URL, "GET", "/loglevel")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("should get OK status")
	}
	loglevels := []string{"", "debug", "info", "warn", "error", "fatal"}
	for _, p := range loglevels {
		path := fmt.Sprintf("/loglevel/%s", p)
		res, _, err = testingtools.HTTPRequest(s.URL, "PUT", path)
		if err != nil {
			t.Error(err)
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("should get OK status for level '%s'", p)
		}
	}
}
