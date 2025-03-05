package controller

/*
func TestAPIGetSubscriptionTypesHandlerWithNoErr(t *testing.T) {
	serverWithNoErr := GetTestHTTPServer(NewUseCaseTestWithNoErr())

	test := struct {
		name           string
		expectedStatus int
		expectedBody   []entity.SubscriptionType
	}{
		name:           "success",
		expectedStatus: http.StatusOK,
		expectedBody:   []entity.SubscriptionType{},
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

		if responseRecorder.Body != test.expectedBody {
			t.Errorf("handler returned unexpected body: got %v want %v", responseRecorder.Body, test.expectedBody)
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
*/
