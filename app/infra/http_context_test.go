package infra

import (
	"errors"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type dummyInputData struct {
	ID   int
	Name string
}

type dummyResponseBody struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestBind(t *testing.T) {
	testCases := []struct {
		name         string
		body         string
		expectedErr  error
		expectedData *dummyInputData
	}{
		{
			name:         "normal",
			body:         `{"id":1,"name":"test"}`,
			expectedErr:  nil,
			expectedData: &dummyInputData{ID: 1, Name: "test"},
		},
		{
			name:         "invalid request",
			body:         `{"id":"test","name":1}`,
			expectedErr:  errors.New("invalid request"),
			expectedData: &dummyInputData{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			body := strings.NewReader(tc.body)
			r := httptest.NewRequest("Post", "/dummies", body)
			w := httptest.NewRecorder()

			context := NewHttpContext(w, r)
			data := new(dummyInputData)

			err := context.Bind(&data)
			if !reflect.DeepEqual(tc.expectedErr, err) {
				t.Errorf("context.Bind expected = %v, got = %v", tc.expectedErr, err)
			}
			if !reflect.DeepEqual(tc.expectedData, data) {
				t.Errorf("context.Bind expected = %v, got = %v", tc.expectedData, data)
			}
		})
	}
}

func TestJSON(t *testing.T) {
	testCases := []struct {
		name         string
		body         interface{}
		code         int
		expectedErr  error
		expectedBody string
	}{
		{
			name:         "normal",
			body:         dummyResponseBody{ID: 1, Name: "test"},
			code:         http.StatusOK,
			expectedErr:  nil,
			expectedBody: "{\"id\":1,\"name\":\"test\"}",
		},
		{
			name:         "json.Marshal Error",
			body:         math.Inf(0),
			code:         http.StatusOK,
			expectedErr:  errors.New("internal server error"),
			expectedBody: "",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := httptest.NewRequest("Get", "/dummy", nil)
			w := httptest.NewRecorder()
			context := NewHttpContext(w, r)

			err := context.JSON(tc.code, tc.body)
			if !reflect.DeepEqual(tc.expectedErr, err) {
				t.Errorf("context.JSON expected = %v, got = %v", tc.expectedErr, err)
			}

			resp := w.Result()
			if resp.StatusCode != tc.code {
				t.Errorf("context.JSON StatusCode expected = %v, got = %v", tc.code, resp.StatusCode)
			}

			body, _ := ioutil.ReadAll(resp.Body)
			if string(body) != tc.expectedBody {
				t.Errorf("context.JSON ResponseBody expected = %v, got = %v", tc.expectedBody, string(body))
			}
		})
	}
}
