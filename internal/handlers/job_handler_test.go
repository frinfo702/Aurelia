package handlers_test

import (
	"Aurelia/internal/domain/models"
	"Aurelia/internal/handlers/testdata"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetJobsHandler
func TestGetJobsHandler(t *testing.T) {
	// create a request and response recorder
	req, rec := createRequest("GET", "/api/jobs", nil)

	// execute the handler
	jobHandler.GetJobsHandler(rec, req)

	// check the response status code
	assert.Equal(t, http.StatusOK, rec.Code)

	var jobs []models.Job
	err := json.NewDecoder(rec.Body).Decode(&jobs)
	assert.NoError(t, err)
	assert.Equal(t, testdata.JobTestData, jobs)
}

// TestGetJobByIDHandler
func TestGetJobByIDHandler(t *testing.T) {
	// create a request and response recorder
	req, rr := createRequest("GET", "/api/jobs/1", map[string]string{"id": "1"})

	// ハンドラーの実行
	jobHandler.GetJobByIDHandler(rr, req)

	// レスポンスの検証
	assert.Equal(t, http.StatusOK, rr.Code)

	var job models.Job
	err := json.NewDecoder(rr.Body).Decode(&job)
	assert.NoError(t, err)
	assert.Equal(t, testdata.JobTestData[0], job)
}
