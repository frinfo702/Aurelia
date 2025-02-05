package handlers_test

import (
	"Aurelia/internal/domain/usecase"
	"Aurelia/internal/handlers"
	"Aurelia/internal/handlers/testdata"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

// To  avoid error "not used"
var jobHandler *handlers.JobHandler

func TestMain(m *testing.M) {
	// initialize mock repository
	mockRepo := new(testdata.MockJobRepository)

	// テストデータに基づくモックの動作設定(個別テストファイルで上書きする場合も)
	mockRepo.On("FindAll").Return(testdata.JobTestData, nil)
	mockRepo.On("FindByID", 1).Return(&testdata.JobTestData[0], nil)
	mockRepo.On("Insert", mock.Anything).Return(nil)

	// inject mock repository to usecase
	jobUseCase := usecase.NewJobUsecase(mockRepo)

	// inject usecase to handler
	jobHandler = handlers.NewJobHandler(jobUseCase)

	m.Run()
}

// helper function to create a common test case
func createRequest(method, url string, vars map[string]string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, nil)
	rec := httptest.NewRecorder()

	// set url parameters
	if vars != nil {
		req = mux.SetURLVars(req, vars)
	}
	return req, rec
}
