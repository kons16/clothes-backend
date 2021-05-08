package handler

import (
	"github.com/golang/mock/gomock"
	"github.com/kons16/team7-backend/domain/repository/mock_repository"
	"github.com/kons16/team7-backend/usecase"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_FindUserBySession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSession := mock_repository.NewMockSession(ctrl)
	sessionUseCase := usecase.NewSessionUseCase(mockSession)
	sessionHandler := NewSessionHandler(sessionUseCase)

	mockSession.EXPECT().CheckBySession("session_id").Return(1)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/is_login", nil)
	rec := httptest.NewRecorder()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "sessionID=session_id")

	// Act
	sessionHandler.FindUserBySession(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"is_login\":\"true\"}", rec.Body.String())
}

func Test_Logout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSession := mock_repository.NewMockSession(ctrl)
	sessionUseCase := usecase.NewSessionUseCase(mockSession)
	sessionHandler := NewSessionHandler(sessionUseCase)

	mockSession.EXPECT().Logout("true").Return(true)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/is_login", nil)
	rec := httptest.NewRecorder()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "isLoggedIn=true")

	// Act
	sessionHandler.Logout(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"is_logout\":\"true\"}", rec.Body.String())
}
