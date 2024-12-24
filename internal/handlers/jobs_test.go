package handlers_test

import (
	"Aurelia/internal/handlers"
	"Aurelia/internal/models"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB

func TestGetJobsHandler(t *testing.T) {
	tests := []struct {
		name           string
		jobType        string
		expectedCount  int
		expectedStatus int
	}{
		{
			name:           "Get all jobs",
			jobType:        "",
			expectedCount:  3,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Get intern jobs",
			jobType:        "intern",
			expectedCount:  2,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Get fulltime jobs",
			jobType:        "fulltime",
			expectedCount:  1,
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := handlers.NewJobHandler(db)

			req := httptest.NewRequest("GET", "/api/jobs?type="+tt.jobType, nil)
			w := httptest.NewRecorder()

			handler.GetJobsHandler(w, req)

			resp := w.Result()
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			body, _ := io.ReadAll(resp.Body)
			var jobs []models.Job
			err := json.Unmarshal(body, &jobs)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCount, len(jobs))
		})
	}
}

func TestGetJobDetailHandler(t *testing.T) {
	tests := []struct {
		name           string
		jobID          int
		expectedJob    *models.Job
		expectedStatus int
	}{
		{
			name:  "Valid job ID returns correct job",
			jobID: 1,
			expectedJob: &models.Job{
				JobID:          1,
				CompanyID:      1,
				HiringType:     "intern",
				TechnologyType: "React",
				IncomeRange:    300000,
				JobTag:         "Backend",
				Requirements:   "Go experience",
				UsedTechnology: "Go, Docker",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Non-existence job ID",
			jobID:          999,
			expectedJob:    nil,
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/api/jobs/detail?id="+strconv.Itoa(tt.jobID), nil)
			w := httptest.NewRecorder()
			handler := handlers.NewJobHandler(db)
			handler.GetJobDetailHandler(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d; got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedJob != nil {
				var gotJob models.Job
				if err := json.NewDecoder(w.Body).Decode(&gotJob); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				if !reflect.DeepEqual(gotJob, *tt.expectedJob) {
					t.Errorf("expected job %+v; got %+v", tt.expectedJob, gotJob)
				}
			}
		})
	}
}
