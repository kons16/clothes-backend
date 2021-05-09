package MySQL

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

//func Test_Create(t *testing.T) {
//	mockDB, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer mockDB.Close()
//
//	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO users (id, name, submit_id, year, sex, password_hash) VALUES (?, ?, ?, ?, ?, ?)`)).
//		WithArgs(1, "test_name", "10", 2000, 0, "password_hash").
//		WillReturnResult(sqlmock.NewResult(1, 1))
//
//	user := &entity.User{
//		ID:           1,
//		Name:         "test_name",
//		SubmitID:     "10",
//		Year:         2000,
//		Sex:          0,
//		PasswordHash: "password_hash",
//	}
//
//	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
//	mockUserRepository := NewUserRepository(sqlxDB)
//
//	// Act
//	userId, err := mockUserRepository.Create(user)
//	if err != nil {
//		t.Errorf("error was not expected while updating stats: %s", err)
//	}
//
//	fmt.Println(userId)
//
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//	}
//}

func Test_FindUserBySubmitID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	columns := []string{"id", "password_hash"}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, password_hash FROM users WHERE submit_id = ?`)).
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "password_hash"))

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	mockUserRepository := NewUserRepository(sqlxDB)

	// Act
	user, err := mockUserRepository.FindUserBySubmitID("1")
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, "password_hash", user.PasswordHash)
	assert.Equal(t, 1, user.ID)
}
