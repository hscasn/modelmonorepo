package ready

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/apiresponse"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/health"
	"github.com/hscasn/modelmonorepo/lib/go/httpserver/testingtools"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	router := chi.NewRouter()
	checks := health.Checks{}
	New(router, checks)
	s := httptest.NewServer(router)
	defer s.Close()

	res, _, err := testingtools.HTTPRequest(s.URL, "GET", "/")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("should get OK status")
	}
}

type dummyCheckThatSucceeds struct{}

func (d *dummyCheckThatSucceeds) Ping() bool {
	return true
}

type dummyCheckThatFails struct{}

func (d *dummyCheckThatFails) Ping() bool {
	return false
}

func TestMakeDefaultNoServices(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("", "http://localhost", nil)
	h := controller(health.Checks{})

	h(w, r)

	res := &apiresponse.Response{}
	json.Unmarshal([]byte(w.Body.String()), res)

	if w.Code != 200 {
		t.Errorf("should have OK status code header")
	}
	if res.Status.AdditionalInfo != nil {
		t.Errorf("response should not include additional info")
	}
	resHealth := res.Result.(map[string]interface{})["status"].(string)
	if resHealth != health.HEALTHY {
		t.Errorf(
			"response should have healthy status, but is \"%s\"",
			res.Result.(string))
	}
}

func TestMakeDefaultTwoThatSucceed(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("", "http://localhost", nil)
	h := controller(health.Checks{
		"one": &dummyCheckThatSucceeds{},
		"two": &dummyCheckThatSucceeds{},
	})

	h(w, r)

	res := &apiresponse.Response{}
	json.Unmarshal([]byte(w.Body.String()), res)

	if w.Code != 200 {
		t.Errorf("should have OK status code header")
	}
	if res.Status.AdditionalInfo != nil {
		t.Errorf("response should not include additional info")
	}
	resHealth := res.Result.(map[string]interface{})["status"].(string)
	if resHealth != health.HEALTHY {
		t.Errorf(
			"response should have healthy status, but is \"%s\"",
			res.Result.(string))
	}
}

func TestMakeDefaultOneThatFails(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("", "http://localhost", nil)
	h := controller(health.Checks{
		"one": &dummyCheckThatSucceeds{},
		"two": &dummyCheckThatFails{},
	})

	h(w, r)

	res := &apiresponse.Response{}
	json.Unmarshal([]byte(w.Body.String()), res)

	if w.Code != 503 {
		t.Errorf("should have Service Unavailable status code header")
	}
	if res.Status.AdditionalInfo != nil {
		t.Errorf("response should not include additional info")
	}
	resHealth := res.Result.(map[string]interface{})["status"].(string)
	if resHealth != health.UNHEALTHY {
		t.Errorf(
			"response should have unhealthy status, but is \"%s\"",
			res.Result.(string))
	}
}

func TestMakeDefaultTwoThatFail(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("", "http://localhost", nil)
	h := controller(health.Checks{
		"one": &dummyCheckThatFails{},
		"two": &dummyCheckThatFails{},
	})

	h(w, r)

	res := &apiresponse.Response{}
	json.Unmarshal([]byte(w.Body.String()), res)

	if w.Code != 503 {
		t.Errorf("should have Service Unavailable status code header")
	}
	if res.Status.AdditionalInfo != nil {
		t.Errorf("response should not include additional info")
	}
	resHealth := res.Result.(map[string]interface{})["status"].(string)
	if resHealth != health.UNHEALTHY {
		t.Errorf(
			"response should have unhealthy status, but is \"%s\"",
			res.Result.(string))
	}
}
