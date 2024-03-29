package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rfinochi/golang-workshop-todo/pkg/common"
	"github.com/rfinochi/golang-workshop-todo/pkg/models"
)

const apiToken string = "85ba6be3-b2d5-4c15-aae5-d4878dfa203c"

func TestCompleteAPIInMemory(t *testing.T) {
	os.Setenv(common.RepositoryEnvVarName, common.RepositoryMemory)
	os.Setenv(common.APITokenEnvVarName, apiToken)

	app := &application{
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
	app.initModels()
	app.initRouter()
	app.addAPIRoutes()

	doAllAPIRequests(t, app)
}

func TestCompleteAPIInMongo(t *testing.T) {
	os.Setenv(common.RepositoryEnvVarName, common.RepositoryMongo)
	os.Setenv(common.APITokenEnvVarName, apiToken)

	app := &application{
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
	app.initModels()
	app.initRouter()
	app.addAPIRoutes()

	doAllAPIRequests(t, app)
}

func TestConnectionErrorMongo(t *testing.T) {
	os.Setenv(common.RepositoryEnvVarName, common.RepositoryMongo)
	os.Setenv(common.RepositoryMongoURIEnvVarName, "mongodb://bad:99999")
	os.Setenv(common.APITokenEnvVarName, apiToken)

	app := &application{
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
	app.initModels()
	app.initRouter()
	app.addAPIRoutes()

	doError(app.router, t, "GET", "/api/1", "", http.StatusInternalServerError)
	doError(app.router, t, "GET", "/api/", "", http.StatusInternalServerError)
	doError(app.router, t, "PUT", "/api/", `{"id":1,"title":"Test_1","isdone":true}`, http.StatusInternalServerError)
	doError(app.router, t, "POST", "/api/", `{"id":1,"title":"Test_1","isdone":true}`, http.StatusInternalServerError)
	doError(app.router, t, "PATCH", "/api/1", `{"id":1,"title":"Test_1","isdone":true}`, http.StatusInternalServerError)
	doError(app.router, t, "DELETE", "/api/1", "", http.StatusInternalServerError)
}

func TestSwagger(t *testing.T) {
	os.Setenv(common.RepositoryEnvVarName, common.RepositoryMemory)
	os.Setenv(common.APITokenEnvVarName, apiToken)

	app := &application{
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
	app.initModels()
	app.initRouter()
	app.addSwaggerRoutes()

	request := doRequesWithAPIToken(app.router, "GET", "/api-docs", "", "")

	assert.Equal(t, 301, request.Code)
	assert.Greater(t, len(request.Body.String()), 0)
}

func TestBadRequestError(t *testing.T) {
	os.Setenv(common.RepositoryEnvVarName, common.RepositoryMemory)
	os.Setenv(common.APITokenEnvVarName, apiToken)

	app := &application{
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
	app.initModels()
	app.initRouter()
	app.addAPIRoutes()

	doError(app.router, t, "PUT", "/api/", "BAD", http.StatusBadRequest)
	doError(app.router, t, "POST", "/api/", "BAD", http.StatusBadRequest)
	doError(app.router, t, "PATCH", "/api/1", "BAD", http.StatusBadRequest)
}

func TestNotFoundError(t *testing.T) {
	os.Setenv(common.RepositoryEnvVarName, common.RepositoryMemory)
	os.Setenv(common.APITokenEnvVarName, apiToken)

	app := &application{
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
	app.initModels()
	app.initRouter()
	app.addAPIRoutes()

	doError(app.router, t, "GET", "/api/bad", "", http.StatusNotFound)
	doError(app.router, t, "DELETE", "/api/bad", "", http.StatusNotFound)
	doError(app.router, t, "PATCH", "/api/bad", "", http.StatusNotFound)

	doError(app.router, t, "GET", "/api/-1", "", http.StatusNotFound)
	doError(app.router, t, "DELETE", "/api/-1", "", http.StatusNotFound)
	doError(app.router, t, "PATCH", "/api/-1", "", http.StatusNotFound)

	doError(app.router, t, "GET", "/api/4", "", http.StatusNotFound)
}

func TestPageNotFoundError(t *testing.T) {
	os.Setenv(common.RepositoryEnvVarName, common.RepositoryMemory)
	os.Setenv(common.APITokenEnvVarName, apiToken)

	app := &application{
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
	app.initModels()
	app.initRouter()

	doError(app.router, t, "GET", "/api/1", "", http.StatusNotFound)
}

func TestRequestWithoutAPIToken(t *testing.T) {
	os.Setenv(common.RepositoryEnvVarName, common.RepositoryMemory)
	os.Setenv(common.APITokenEnvVarName, apiToken)

	app := &application{
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
	app.initModels()
	app.initRouter()
	app.addAPIRoutes()

	doErrorWithAPIToken(app.router, t, "GET", "/api/1", "", "", http.StatusUnauthorized)
	doErrorWithAPIToken(app.router, t, "GET", "/api/1", "", "X", http.StatusUnauthorized)
}

func TestInternalServerError(t *testing.T) {
	os.Setenv(common.RepositoryEnvVarName, common.RepositoryGoogle)
	os.Setenv(common.APITokenEnvVarName, apiToken)

	app := &application{
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
	app.initModels()
	app.initRouter()
	app.addAPIRoutes()

	doError(app.router, t, "POST", "/api/", `{"id":1,"title":"Test_1","isdone":true}`, http.StatusInternalServerError)
	doError(app.router, t, "PUT", "/api/", `{"id":1,"title":"Test_1","isdone":true}`, http.StatusInternalServerError)
	doError(app.router, t, "GET", "/api/", "", http.StatusInternalServerError)
	doError(app.router, t, "GET", "/api/1", "", http.StatusInternalServerError)
	doError(app.router, t, "PATCH", "/api/1", `{"id":1,"title":"Test_1","isdone":true}`, http.StatusInternalServerError)
	doError(app.router, t, "DELETE", "/api/1", "", http.StatusInternalServerError)
}

func TestValidations(t *testing.T) {
	os.Setenv(common.RepositoryEnvVarName, common.RepositoryMemory)
	os.Setenv(common.APITokenEnvVarName, apiToken)

	app := &application{
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
	app.initModels()
	app.initRouter()
	app.addAPIRoutes()

	doPostItem(app.router, t, "POST", `{"id":8,"title":"Test_87","isdone":true}`)

	doError(app.router, t, "POST", "/api/", `{"id":8,"title":"Test_8","isdone":true}`, http.StatusConflict)
	doError(app.router, t, "PUT", "/api/", `{"id":8,"title":"Test_8","isdone":true}`, http.StatusConflict)
	doError(app.router, t, "PATCH", "/api/9", `{"id":9,"title":"Test_9","isdone":true}`, http.StatusNotFound)
	doError(app.router, t, "DELETE", "/api/9", "", http.StatusNotFound)
}

func doAllAPIRequests(t *testing.T, a *application) {
	doCleanUp(a.router, t)

	doGetItems(a.router, t, "", true, 0)

	doPostItem(a.router, t, "POST", `{"id":1,"title":"Test_1","isdone":true}`)
	doGetItems(a.router, t, "Test_1", true, 1)
	doGetItem(a.router, t, 1, "Test_1", true)

	doPostItem(a.router, t, "PUT", `{"id":2,"title":"Test_2","isdone":true}`)
	doGetItems(a.router, t, "Test_1", true, 2)
	doGetItem(a.router, t, 2, "Test_2", true)

	doDeleteItem(a.router, t, 2)
	doGetItems(a.router, t, "Test_1", true, 1)

	doPatchItem(a.router, t, 1, `{"id":1,"title":"Test_3","isdone":false}`)
	doGetItems(a.router, t, "Test_3", false, 1)
	doGetItem(a.router, t, 1, "Test_3", false)
}

func doGetItem(r http.Handler, t *testing.T, id int, title string, isdone bool) {
	request := doRequest(r, "GET", fmt.Sprintf("/api/%v", id), "")

	assert.Equal(t, http.StatusOK, request.Code)

	var response models.Item

	err := json.Unmarshal(request.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, response.ID, id)
	assert.Equal(t, response.Title, title)
	assert.Equal(t, response.IsDone, isdone)
}

func doGetItems(r http.Handler, t *testing.T, title string, isdone bool, length int) {
	request := doRequest(r, "GET", "/api/", "")

	assert.Equal(t, http.StatusOK, request.Code)

	var response []models.Item

	err := json.Unmarshal(request.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, len(response), length)
	if length > 0 {
		assert.Equal(t, response[0].ID, 1)
		assert.Equal(t, response[0].Title, title)
		assert.Equal(t, response[0].IsDone, isdone)
	}
}

func doPostItem(r http.Handler, t *testing.T, method string, payload string) {
	request := doRequest(r, method, "/api/", payload)

	assert.Equal(t, http.StatusCreated, request.Code)

	var response map[string]string

	err := json.Unmarshal(request.Body.Bytes(), &response)
	value, exists := response["message"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "OK", value)
}

func doPatchItem(r http.Handler, t *testing.T, id int, payload string) {
	request := doRequest(r, "PATCH", fmt.Sprintf("/api/%v", id), payload)

	assert.Equal(t, http.StatusOK, request.Code)

	var response map[string]string

	err := json.Unmarshal(request.Body.Bytes(), &response)
	value, exists := response["message"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "OK", value)
}

func doDeleteItem(r http.Handler, t *testing.T, id int) {
	request := doRequest(r, "DELETE", fmt.Sprintf("/api/%v", id), "")

	assert.Equal(t, http.StatusOK, request.Code)

	var response map[string]string

	err := json.Unmarshal(request.Body.Bytes(), &response)
	value, exists := response["message"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "OK", value)
}

func doCleanUp(r http.Handler, t *testing.T) {
	request := doRequest(r, "GET", "/api/", "")

	assert.Equal(t, http.StatusOK, request.Code)

	var response []models.Item
	s := request.Body.String()
	err := json.Unmarshal([]byte(s), &response)

	assert.Nil(t, err)

	for i := 0; i < len(response); i++ {
		doDeleteItem(r, t, response[i].ID)
	}
}

func doError(r http.Handler, t *testing.T, method string, path string, payload string, errorCode int) {
	request := doRequest(r, method, path, payload)

	assert.Equal(t, errorCode, request.Code)
}

func doErrorWithAPIToken(r http.Handler, t *testing.T, method string, path string, token string, payload string, errorCode int) {
	request := doRequesWithAPIToken(r, method, path, token, payload)

	assert.Equal(t, errorCode, request.Code)
}

func doRequest(r http.Handler, method string, path string, payload string) *httptest.ResponseRecorder {
	return doRequesWithAPIToken(r, method, path, apiToken, payload)
}

func doRequesWithAPIToken(r http.Handler, method string, path string, token string, payload string) *httptest.ResponseRecorder {
	var req *http.Request

	if method == "POST" || method == "PATCH" || method == "PUT" {
		req, _ = http.NewRequest(method, path, bytes.NewBuffer([]byte(payload)))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, _ = http.NewRequest(method, path, nil)
	}

	req.Header.Set(common.APITokenHeaderName, token)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}
