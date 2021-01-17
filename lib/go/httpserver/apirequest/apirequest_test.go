package apirequest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

func mockRequestWithBody(body interface{}) *http.Request {
	mbody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPost, "", bytes.NewReader(mbody))
	if err != nil {
		panic(err)
	}
	return req
}

func mustStringify(body interface{}) string {
	b, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func TestParse_EmptyStruct(t *testing.T) {
	type reqBody struct{}

	want := reqBody{}
	request := mockRequestWithBody(want)
	target := reqBody{}
	wantErr := false

	err := Parse(request, &target)
	if (err != nil) != wantErr {
		t.Errorf("Parse() error = %v, wantErr %v", err, wantErr)
	}
	if target != want {
		t.Errorf("Parse() target = %v, want %v", target, want)
	}
}

func TestParse_DiverseFields(t *testing.T) {
	type reqBody struct {
		Number  int64     `json:"number"`
		Text    string    `json:"text"`
		Boolean bool      `json:"boolean"`
		Decimal float64   `json:"decimal"`
		Date    time.Time `json:"date"`
	}

	want := reqBody{
		Number:  10,
		Text:    "hello world",
		Boolean: true,
		Decimal: 11.12,
		Date:    time.Now().Round(time.Minute),
	}
	request := mockRequestWithBody(want)
	target := reqBody{}
	wantErr := false

	err := Parse(request, &target)
	if (err != nil) != wantErr {
		t.Errorf("Parse() error = %v, wantErr %v", err, wantErr)
	}
	if target != want {
		t.Errorf("Parse() target = %v, want %v", target, want)
	}
}

func TestParse_DiverseFieldsNullable_Filled(t *testing.T) {
	type reqBody struct {
		Number  *int       `json:"number"`
		Text    *string    `json:"text"`
		Boolean *bool      `json:"boolean"`
		Decimal *float64   `json:"decimal"`
		Date    *time.Time `json:"date"`
	}

	number := 10
	text := "hello world"
	boolean := true
	decimal := 11.12
	date := time.Now().Round(time.Minute)

	want := reqBody{
		Number:  &number,
		Text:    &text,
		Boolean: &boolean,
		Decimal: &decimal,
		Date:    &date,
	}
	request := mockRequestWithBody(want)
	target := reqBody{}
	wantErr := false

	err := Parse(request, &target)
	if (err != nil) != wantErr {
		t.Errorf("Parse() error = %v, wantErr %v", err, wantErr)
	}

	starget := mustStringify(target)
	swant := mustStringify(want)
	if starget != swant {
		t.Errorf("Parse() target = %v, want %v", starget, swant)
	}
}

func TestParse_DiverseFieldsNullable_Null(t *testing.T) {
	type reqBody struct {
		Number  *int       `json:"number"`
		Text    *string    `json:"text"`
		Boolean *bool      `json:"boolean"`
		Decimal *float64   `json:"decimal"`
		Date    *time.Time `json:"date"`
	}

	want := reqBody{}
	request := mockRequestWithBody(want)
	target := reqBody{}
	wantErr := false

	err := Parse(request, &target)
	if (err != nil) != wantErr {
		t.Errorf("Parse() error = %v, wantErr %v", err, wantErr)
	}

	starget := mustStringify(target)
	swant := mustStringify(want)
	if starget != swant {
		t.Errorf("Parse() target = %v, want %v", starget, swant)
	}
}
