package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	// default variabel nya i, namun agar lebih mudah dipahami diganti menjadi error
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, error interface{}) {
		fmt.Fprint(w, "Panic : ", error)
	}

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("Ups")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Panic : Ups", string(body))
}
