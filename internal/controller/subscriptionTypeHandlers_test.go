package controller

import (
	"SystemSubscription/internal/entity"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestAPIGetSubscriptionTypesHandlerWithNoErr(t *testing.T) {
	serverWithNoErr := GetTestHTTPServer(NewUseCaseTestWithNoErr())

	test := struct {
		name           string
		expectedStatus int
		expectedBody   []entity.SubscriptionType
	}{
		name:           "success",
		expectedStatus: http.StatusOK,
		expectedBody: []entity.SubscriptionType{
			{
				ID:       1,
				TypeName: "1 second",
				Duration: 1,
				Price:    10,
			},
			{
				ID:       2,
				TypeName: "3 seconds",
				Duration: 3,
				Price:    30,
			},
			{
				ID:       3,
				TypeName: "6 seconds",
				Duration: 6,
				Price:    60,
			},
			{
				ID:       4,
				TypeName: "12 seconds",
				Duration: 12,
				Price:    120,
			},
		},
	}

	t.Run(test.name, func(t *testing.T) {
		request, err := http.NewRequest("GET", "/subscription-types", nil)
		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder() //создание объекта иммитирующего http.ResponseWriter
		router := mux.NewRouter()
		router.HandleFunc("/subscription-types", serverWithNoErr.APIGetSubscriptionTypesHandler)
		router.ServeHTTP(responseRecorder, request)

		if status := responseRecorder.Code; status != test.expectedStatus {
			t.Errorf("handler returned wrong status code: got %v want %v", status, test.expectedStatus)
		}
		var responseRecorderBody []entity.SubscriptionType
		err = json.NewDecoder(responseRecorder.Body).Decode(&responseRecorderBody)
		if err != nil {
			t.Errorf("handler returned unexpected body: got %v want %v", responseRecorderBody, test.expectedBody)
		}

	})
}

func TestAPIGetSubscriptionTypesHandlerWithErr(t *testing.T) {
	serverWithErr := GetTestHTTPServer(NewUseCaseTestWithErr())

	test := struct {
		name           string
		expectedStatus int
		expectedBody   string
	}{
		name:           "error",
		expectedStatus: http.StatusInternalServerError,
		expectedBody:   "test\n",
	}

	t.Run(test.name, func(t *testing.T) {
		request, err := http.NewRequest("GET", "/subscription-types", nil)
		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder() //создание объекта иммитирующего http.ResponseWriter
		router := mux.NewRouter()
		router.HandleFunc("/subscription-types", serverWithErr.APIGetSubscriptionTypesHandler)
		router.ServeHTTP(responseRecorder, request)

		if status := responseRecorder.Code; status != test.expectedStatus {
			t.Errorf("handler returned wrong status code: got %v want %v", status, test.expectedStatus)
		}

		if responseRecorder.Body.String() != test.expectedBody {
			t.Errorf("handler returned unexpected body: got %v want %v", responseRecorder.Body.String(), test.expectedBody)
		}
	})
}
