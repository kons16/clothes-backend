package MySQL

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/kons16/team7-backend/domain/entity"
	"github.com/kons16/team7-backend/domain/repository/mock_repository"
	"testing"
)

func Test_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_repository.NewMockUser(ctrl)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO users").WithArgs(1, "test_name", "10", 2000, 0, "password_hash", nil, nil).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	user := &entity.User{
		ID:           1,
		Name:         "test_name",
		SubmitID:     "10",
		Year:         2000,
		Sex:          0,
		PasswordHash: "password_hash",
	}

	// Act
	userId, err := mockUserRepository.Create(user)
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	fmt.Println(userId)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
