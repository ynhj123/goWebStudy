package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestHelloRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello?name=test1003715284185477120", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var res Response
	errs := json.Unmarshal(w.Body.Bytes(), &res)
	if errs != nil {
		fmt.Println("json unmarshal error:", errs)
	}
	fmt.Println("res2 code:", res.Code)
	fmt.Println("res2 msg:", res.Msg)

	assert.Equal(t, 0, res.Code)
	assert.Equal(t, "success", res.Msg)
	jsons, errs := json.Marshal(res.Data)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("json data :", string(jsons))
}
