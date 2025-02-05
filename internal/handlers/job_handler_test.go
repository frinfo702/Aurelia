package handlers_test

import (
	"Aurelia/internal/domain/models"
	"Aurelia/internal/domain/usecase"
	"Aurelia/internal/handlers"
	"Aurelia/internal/handlers/testdata"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	// use go test container
)

// TestGetJobsHandler
func TestGetJobsHandler(t *testing.T) {
	testCase := []struct {
		name           string
		mockSetup      func(mockRepo *testdata.MockJobRepository)
		expectedStatus int
		expectedBody   []models.Job
	}{
		{
			name: "success",
			mockSetup: func(mockRepo *testdata.MockJobRepository) {
				mockRepo.On("FindAll").Return(testdata.JobTestData, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   testdata.JobTestData,
		},
		{
			name: "failed",
			mockSetup: func(mockRepo *testdata.MockJobRepository) {
				mockRepo.On("FindAll").Return([]models.Job(nil), errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   nil,
		},
		{
			name: "empty",
			mockSetup: func(mockRepo *testdata.MockJobRepository) {
				mockRepo.On("FindAll").Return([]models.Job{}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   []models.Job{},
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			// initialize mock repository
			mockRepo := new(testdata.MockJobRepository)
			tt.mockSetup(mockRepo)

			// initialize usecase and handler
			jobUsecase := usecase.NewJobUsecase(mockRepo)
			jobHandler := handlers.NewJobHandler(jobUsecase)

			// create a request and response recorder
			req, rec := createRequest("GET", "/api/jobs", nil)

			// execute the handler
			jobHandler.GetJobsHandler(rec, req)

			// check the response status code
			assert.Equal(t, tt.expectedStatus, rec.Code)

			// assert the response body
			if tt.expectedStatus == http.StatusOK {
				var jobs []models.Job
				err := json.NewDecoder(rec.Body).Decode(&jobs)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, jobs)
			} else {
				var errResp map[string]string
				err := json.NewDecoder(rec.Body).Decode(&errResp)
				assert.NoError(t, err)
				assert.Contains(t, errResp, "error")
			}

			// assert the mock behavior
			mockRepo.AssertExpectations(t)
		})
	}
}

// TestGetJobByIDHandler
func TestGetJobByIDHandler(t *testing.T) {
	testCase := []struct {
		name           string
		mockSetup      func(mockRepo *testdata.MockJobRepository)
		expectedStatus int
		expectedBody   models.Job
	}{
		{
			name: "success",
			mockSetup: func(mockRepo *testdata.MockJobRepository) {
				mockRepo.On("FindByID", 1).Return(&testdata.JobTestData[0], nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   testdata.JobTestData[0],
		},
		{
			name: "failed",
			mockSetup: func(mockRepo *testdata.MockJobRepository) {
				mockRepo.On("FindByID", 1).Return(nil, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   models.Job{},
		},
		{
			name: "not found",
			mockSetup: func(mockRepo *testdata.MockJobRepository) {
				mockRepo.On("FindByID", 999).Return((*models.Job)(nil), errors.New("not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   models.Job{},
		},
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {

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
		})
	}
}

// func TestCreateJobHandler(t *testing.T) {
// 	testCase := []struct {
// 		name           string
// 		mockSetup      func(mockRepo *testdata.MockJobRepository)
// 		expectedStatus int
// 		expectedBody   map[string]string
// 	}{
// 		{
// 			name: "success",
// 			mockSetup: func(mockRepo *testdata.MockJobRepository) {
// 				mockRepo.On("Insert", mock.Anything).Return(nil)
// 			},
// 			expectedStatus: http.StatusCreated,
// 			expectedBody:   map[string]string{"message": "job created"},
// 		},
// 	}

// 	for _, tt := range testCase {
// 		t.Run(tt.name, func(t *testing.T) {

// 			// create request
// 			req, rr := createRequest("POST", "/api/jobs", nil)

// 			// execute handler
// 			jobHandler.CreateJobHandler(rr, req)

// 			// check response status code
// 			assert.Equal(t, tt.expectedStatus, rr.Code)

// 			// assert response body
// 			var resp map[string]string
// 			err := json.NewDecoder(rr.Body).Decode(&resp)
// 			assert.NoError(t, err)
// 			assert.Equal(t, tt.expectedBody, resp)

// 		})
// 	}
// }
