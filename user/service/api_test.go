package service_test

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/httpmon/user/config"
	"github.com/httpmon/user/mock"
	"github.com/httpmon/user/service"
	"github.com/labstack/echo/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterEmailEmpty(t *testing.T) {
	t.Parallel()

	cfg := config.Read()

	api := service.API{
		User:   mock.User{Info: map[string]string{}},
		URL:    mock.URL{Urls: map[string]int{}},
		Config: cfg.JWT,
	}

	e := echo.New()
	e.POST("/register", api.Register)

	registerationJSON := `{"Password":"1378"}`

	req := httptest.NewRequestWithContext(
		context.Background(), http.MethodPost, "/register", strings.NewReader(registerationJSON),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	require.NoError(t, err, "Cannot read body")

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	log.Println(string(body))
}

func Register(t *testing.T, api service.API) {
	t.Helper()

	e := echo.New()
	e.POST("/register", api.Register)

	registerationJSON := `{"Email":"raha.alvani@gmail.com",
							"Password":"1378"}`

	req := httptest.NewRequestWithContext(
		context.Background(), http.MethodPost, "/register", strings.NewReader(registerationJSON),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	require.NoError(t, err, "Cannot read body")

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	log.Println(string(body))
}

func Login(t *testing.T, api service.API) string {
	t.Helper()

	e := echo.New()
	e.POST("/login", api.Login)

	loginJSON := `{"Email":"raha.alvani@gmail.com",
							"Password":"1378"}`

	req := httptest.NewRequestWithContext(context.Background(), http.MethodPost, "/login", strings.NewReader(loginJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	resp := rec.Result()
	defer checkClose(resp)

	body, err := io.ReadAll(resp.Body)

	require.NoError(t, err, "Cannot read body")

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	log.Println(string(body))

	var token string

	if err := json.Unmarshal(body, &token); err != nil {
		t.Fatal(err)
	}

	return token
}

func Add(t *testing.T, token string, api service.API) {
	t.Helper()

	e := echo.New()
	e.POST("/url", api.Add)

	addJSON := `{"URL": "https://www.google.com", "Period": 2}`

	req := httptest.NewRequestWithContext(context.Background(), http.MethodPost, "/url", strings.NewReader(addJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", token)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	resp := rec.Result()
	defer checkClose(resp)

	body, err := io.ReadAll(resp.Body)

	require.NoError(t, err, "Cannot read body")

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	log.Println(string(body))
}

func TestAPI(t *testing.T) {
	t.Parallel()

	cfg := config.Read()

	api := service.API{
		User:   mock.User{Info: map[string]string{}},
		URL:    mock.URL{Urls: map[string]int{}},
		Config: cfg.JWT,
	}

	Register(t, api)
	token := Login(t, api)
	Add(t, token, api)
}

func checkClose(resp *http.Response) {
	if err := resp.Body.Close(); err != nil {
		log.Fatal(err)
	}
}
