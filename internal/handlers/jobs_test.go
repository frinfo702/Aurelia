package handlers_test

import (
	"Aurelia/internal/handlers"
	"Aurelia/internal/models"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var db *sql.DB

func setupTestDatabase(ctx context.Context, container testcontainers.Container) error {
	mappedPort, _ := container.MappedPort(ctx, "5432")
	host, _ := container.Host(ctx)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, mappedPort.Port(), "testuser", "testpass", "testdb")

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open db: %v", err)
	}

	// Wait for DB to be ready
	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		return fmt.Errorf("db not ready: %v", err)
	}

	// Execute schema
	schemaSQL, err := os.ReadFile("../../migrations/001_init_schema.sql")
	if err != nil {
		return fmt.Errorf("failed to read schema: %v", err)
	}
	if _, err := db.ExecContext(ctx, string(schemaSQL)); err != nil {
		return fmt.Errorf("failed to execute schema: %v", err)
	}

	// Execute test data
	testDataSQL, err := os.ReadFile("../../migrations/test_data.sql")
	if err != nil {
		return fmt.Errorf("failed to read test data: %v", err)
	}
	if _, err := db.ExecContext(ctx, string(testDataSQL)); err != nil {
		return fmt.Errorf("failed to execute test data: %v", err)
	}

	return nil
}

func TestMain(m *testing.M) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image: "postgres:14",
		Env: map[string]string{
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpass",
			"POSTGRES_DB":       "testdb",
		},
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}

	if err := setupTestDatabase(ctx, container); err != nil {
		panic(err)
	}

	// ここでグローバルDB接続を設定
	handlers.SetDB(db)

	code := m.Run()

	db.Close()
	container.Terminate(ctx)
	os.Exit(code)
}

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
			req := httptest.NewRequest("GET", "/api/jobs/"+strconv.Itoa(tt.jobID), nil)

			// set path parameter
			req = mux.SetURLVars(req, map[string]string{
				"id": strconv.Itoa(tt.jobID),
			})

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
