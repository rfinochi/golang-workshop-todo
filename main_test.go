package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompleteApi(t *testing.T) {
	router := SetupRouter()

	performGetItems(router, t, "", true, 0)
	performGetItem(router, t, 0, "", false)

	performPostItem(router, t, "POST", `{"id":1,"title":"Test_1","isdone":true}`)
	performGetItems(router, t, "Test_1", true, 1)
	performGetItem(router, t, 1, "Test_1", true)

	performPostItem(router, t, "PUT", `{"id":2,"title":"Test_2","isdone":true}`)
	performGetItems(router, t, "Test_1", true, 2)
	performGetItem(router, t, 2, "Test_2", true)

	performDeleteItem(router, t, 2)
	performGetItems(router, t, "Test_1", true, 1)

	performPatchItem(router, t, 1, `{"id":1,"title":"Test_3","isdone":false}`)
	performGetItems(router, t, "Test_3", false, 1)
	performGetItem(router, t, 1, "Test_3", false)
}

func performPostItem(r http.Handler, t *testing.T, method string, payload string) {
	request := performRequest(r, method, "/todo/", payload)

	assert.Equal(t, http.StatusCreated, request.Code)

	var response map[string]string

	err := json.Unmarshal([]byte(request.Body.String()), &response)
	value, exists := response["message"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "OK", value)
}

func performGetItem(r http.Handler, t *testing.T, id int, title string, isdone bool) {
	request := performRequest(r, "GET", fmt.Sprintf("/todo/%v", id), "")

	assert.Equal(t, http.StatusOK, request.Code)

	var response Item

	err := json.Unmarshal([]byte(request.Body.String()), &response)

	assert.Nil(t, err)
	assert.Equal(t, response.ID, id)
	assert.Equal(t, response.Title, title)
	assert.Equal(t, response.IsDone, isdone)
}

func performGetItems(r http.Handler, t *testing.T, title string, isdone bool, length int) {
	request := performRequest(r, "GET", "/todo/", "")

	assert.Equal(t, http.StatusOK, request.Code)

	var response []Item

	err := json.Unmarshal([]byte(request.Body.String()), &response)

	assert.Nil(t, err)
	assert.Equal(t, len(response), length)
	if length > 0 {
		assert.Equal(t, response[0].ID, 1)
		assert.Equal(t, response[0].Title, title)
		assert.Equal(t, response[0].IsDone, isdone)
	}
}

func performDeleteItem(r http.Handler, t *testing.T, id int) {
	request := performRequest(r, "DELETE", fmt.Sprintf("/todo/%v", id), "")

	assert.Equal(t, http.StatusOK, request.Code)

	var response map[string]string

	err := json.Unmarshal([]byte(request.Body.String()), &response)
	value, exists := response["message"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "OK", value)
}

func performPatchItem(r http.Handler, t *testing.T, id int, payload string) {
	request := performRequest(r, "PATCH", fmt.Sprintf("/todo/%v", id), payload)

	assert.Equal(t, http.StatusOK, request.Code)

	var response map[string]string

	err := json.Unmarshal([]byte(request.Body.String()), &response)
	value, exists := response["message"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "OK", value)
}

func performRequest(r http.Handler, method string, path string, payload string) *httptest.ResponseRecorder {
	var req *http.Request

	if method == "POST" || method == "PATCH" || method == "PUT" {
		req, _ = http.NewRequest(method, path, bytes.NewBuffer([]byte(payload)))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, _ = http.NewRequest(method, path, nil)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}
