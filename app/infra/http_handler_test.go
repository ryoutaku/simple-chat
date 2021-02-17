package infra

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/ryoutaku/simple-chat/app/interface/adapter"
)

func newFakeHttpHandler(code int, message string) httpHandler {
	return func(c adapter.HttpContext) *adapter.HttpError {
		switch code {
		case 500:
			panic("panic test")
		case 400, 403, 404:
			return adapter.NewHttpError(message, code)
		default:
			return nil
		}
	}
}

func TestRun(t *testing.T) {
	testCases := []struct {
		name         string
		handler      httpHandler
		expectedBody string
		expectedCode int
	}{
		{
			name:         "normal",
			handler:      newFakeHttpHandler(200, ""),
			expectedBody: "",
			expectedCode: 200,
		},
		{
			name:         "recover panic",
			handler:      newFakeHttpHandler(500, ""),
			expectedBody: "Internal Server Error\n",
			expectedCode: 500,
		},
		{
			name:         "error handling 400",
			handler:      newFakeHttpHandler(400, "test error 400"),
			expectedBody: "test error 400\n",
			expectedCode: 400,
		},
		{
			name:         "error handling 403",
			handler:      newFakeHttpHandler(403, "test error 403"),
			expectedBody: "test error 403\n",
			expectedCode: 403,
		},
		{
			name:         "error handling 404",
			handler:      newFakeHttpHandler(404, "test error 404"),
			expectedBody: "test error 404\n",
			expectedCode: 404,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := httptest.NewRequest("Post", "/dummy", nil)
			w := httptest.NewRecorder()
			tc.handler.run(w, r)

			resp := w.Result()
			if resp.StatusCode != tc.expectedCode {
				t.Errorf("%v: context.JSON StatusCode expected = %v, got = %v", tc.name, tc.expectedCode, resp.StatusCode)
			}

			body, _ := ioutil.ReadAll(resp.Body)
			if string(body) != tc.expectedBody {
				t.Errorf("%v: context.JSON ResponseBody expected = %v, got = %v",
					tc.name, tc.expectedBody, string(body))
			}
		})
	}
}
