package app_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"tender-service/internal/app"
)

func TestCreateTender(t *testing.T) {
	a, err := app.New(context.Background())
	require.NoError(t, err)
	defer a.Close(context.Background(), 5*time.Second)

	tenderData := map[string]interface{}{
		"Title":           "Тендер 1",
		"description":     "Описание тендера",
		"serviceType":     "Construction",
		"status":          "Open",
		"organizationId":  1,
		"creatorUsername": "user1",
	}

	body, err := json.Marshal(tenderData)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/tenders", bytes.NewBuffer(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, req)

	t.Logf("Response Body: %s", recorder.Body.String())

	assert.Equal(t, http.StatusCreated, recorder.Code)

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "Тендер 1", response["name"])
}

func TestGetTenders(t *testing.T) {
	a, err := app.New(context.Background())
	require.NoError(t, err)
	defer a.Close(context.Background(), 5*time.Second)

	username := "user1"
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/tenders/my?username=%s", username), nil)
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response []map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.NotEmpty(t, response)
}

func TestEditTender(t *testing.T) {
	a, err := app.New(context.Background())
	require.NoError(t, err)
	defer a.Close(context.Background(), 5*time.Second)

	tenderEdit := map[string]interface{}{
		"name":        "Обновленный Тендер 1",
		"description": "Обновленное описание",
	}

	body, err := json.Marshal(tenderEdit)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPatch, "/api/v1/tenders/1/edit", bytes.NewBuffer(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "Обновленный Тендер 1", response["name"])
	assert.Equal(t, "Обновленное описание", response["description"])
}

func TestCreateOffer(t *testing.T) {
	a, err := app.New(context.Background())
	require.NoError(t, err)
	defer a.Close(context.Background(), 5*time.Second)

	offerData := map[string]interface{}{
		"name":            "Предложение 1",
		"description":     "Описание предложения",
		"status":          "Submitted",
		"tenderId":        1,
		"organizationId":  1,
		"creatorUsername": "user1",
	}

	body, err := json.Marshal(offerData)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/offers/new", bytes.NewBuffer(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "Предложение 1", response["name"])
}

func TestEditOffer(t *testing.T) {
	a, err := app.New(context.Background())
	require.NoError(t, err)
	defer a.Close(context.Background(), 5*time.Second)

	offerEdit := map[string]interface{}{
		"name":        "Обновленное Предложение 1",
		"description": "Обновленное описание",
	}

	body, err := json.Marshal(offerEdit)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPatch, "/api/v1/offers/1/edit", bytes.NewBuffer(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "Обновленное Предложение 1", response["name"])
}
