package user

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
	"test-monitoring/domain"
	"testing"
	"time"
)

func mockRepositorySetup() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("Failed to create sqlmock.")
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic("Failed to open gorm db.")
	}

	return gormDB, mock
}

func Test_Should_Create_User_With_Mock_Db(t *testing.T) {
	db, mock := mockRepositorySetup()
	repo := NewUserRepository(db)

	// Дано
	user := domain.User{Name: "John Doe", Age: 30, CreatedDate: time.Now()}

	// Когда
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users"`).
		WithArgs(user.Name, user.Age, user.CreatedDate).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	result, err := repo.CreateUser(user)

	// То
	assert.Nil(t, err)
	assert.NotNil(t, result.ID)
}

func Test_Should_Return_Err_When_Invoke_Create_User_With_Mock_Db(t *testing.T) {
	db, mock := mockRepositorySetup()
	repo := NewUserRepository(db)

	// Дано
	user := domain.User{Name: "John Doe", Age: 30, CreatedDate: time.Now()}
	gormErr := errors.New("Unexpected Error")
	unexpectedErr := domain.NewUnexpectedError(gormErr.Error())

	// Когда
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users"`).
		WithArgs(user.Name, user.Age, user.CreatedDate).
		WillReturnError(gormErr)
	mock.ExpectCommit()

	_, err := repo.CreateUser(user)

	// То
	assert.NotNil(t, err)
	assert.Equal(t, unexpectedErr.Code, err.Code)
	assert.True(t, strings.Contains(err.Message, unexpectedErr.Message), "Should contains Unexpected Error")
}

func Test_Should_Get_User_By_Id_With_Mock_Db(t *testing.T) {
	db, mock := mockRepositorySetup()
	repo := NewUserRepository(db)

	// Дано
	user := domain.User{ID: 1, Name: "John Doe", Age: 30, CreatedDate: time.Now()}

	// Когда
	expectedSQL := "SELECT (.+) FROM \"users\" WHERE id =(.+)"
	mock.ExpectQuery(expectedSQL).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age", "created_date"}).
			AddRow(user.ID, user.Name, user.Age, user.CreatedDate))

	result, err := repo.GetUserById(user.ID)

	// То
	assert.Nil(t, err)
	assert.Equal(t, user.Name, result.Name)
}

func Test_Should_Return_Not_Found_Error_When_Invoke_Get_User_By_Id_With_Mock_Db(t *testing.T) {
	db, mock := mockRepositorySetup()
	repo := NewUserRepository(db)

	// Дано
	var id uint = 1
	expectedError := domain.NewNotFoundError("User not found, ID: 1")

	// Когда
	expectedSQL := "SELECT (.+) FROM \"users\" WHERE id =(.+)"
	mock.ExpectQuery(expectedSQL).WillReturnError(gorm.ErrRecordNotFound)

	_, err := repo.GetUserById(id)

	// То
	assert.NotNil(t, err)
	assert.Equal(t, expectedError.Message, err.Message)
	assert.Equal(t, expectedError.Code, err.Code)
}

func Test_Should_Return_Unexpected_Error_When_Invoke_Get_User_By_Id_With_Mock_Db(t *testing.T) {
	db, mock := mockRepositorySetup()
	repo := NewUserRepository(db)

	// Дано
	var id uint = 1
	expectedError := domain.NewUnexpectedError("Unexpected Err")

	// Когда
	expectedSQL := "SELECT (.+) FROM \"users\" WHERE id =(.+)"
	mock.ExpectQuery(expectedSQL).WillReturnError(gorm.ErrNotImplemented)

	_, err := repo.GetUserById(id)

	// То
	assert.NotNil(t, err)
	assert.Equal(t, expectedError.Code, err.Code)
}

func Test_Should_Update_User_With_Mock_Db(t *testing.T) {
	db, mock := mockRepositorySetup()
	repo := NewUserRepository(db)

	// Дано
	user := domain.User{ID: 1, Name: "Edit User", Age: 29, CreatedDate: time.Now()}

	// Когда
	updUserSQL := "UPDATE \"users\" SET .+"
	mock.ExpectBegin()
	mock.ExpectExec(updUserSQL).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	updateUser, err := repo.UpdateUser(user)

	// То
	assert.Nil(t, err)
	assert.Equal(t, user.Name, updateUser.Name)
}

func Test_Should_Return_Unexpected_Err_When_Invoke_Update_User_With_Mock_Db(t *testing.T) {
	db, mock := mockRepositorySetup()
	repo := NewUserRepository(db)

	// Дано
	user := domain.User{ID: 1}
	gormErr := errors.New("Unexpected Error")
	unexpectedErr := domain.NewUnexpectedError(gormErr.Error())

	// Когда
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"users\" SET .+").
		WillReturnError(gormErr)
	mock.ExpectCommit()

	_, err := repo.UpdateUser(user)

	// То
	assert.NotNil(t, err)
	assert.Equal(t, unexpectedErr.Code, err.Code)
	assert.True(t, strings.Contains(err.Message, unexpectedErr.Message), "Should contains Unexpected Error")
}

func Test_Should_Delete_User_With_Mock_Db(t *testing.T) {
	db, mock := mockRepositorySetup()
	repo := NewUserRepository(db)

	// Дано
	user := domain.User{ID: 1}

	// Когда
	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"users\" WHERE (.+)$").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.DeleteUserById(user.ID)

	// То
	assert.Nil(t, err)
}

func Test_Should_Return_Unexpected_Err_When_Invoke_Delete_User_By_Id_With_Mock_Db(t *testing.T) {
	db, mock := mockRepositorySetup()
	repo := NewUserRepository(db)

	// Дано
	user := domain.User{ID: 1}
	gormErr := errors.New("Unexpected Error")
	unexpectedErr := domain.NewUnexpectedError(gormErr.Error())

	// Когда
	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"users\" WHERE (.+)$").
		WillReturnError(gormErr)
	mock.ExpectCommit()

	err := repo.DeleteUserById(user.ID)

	// То
	assert.NotNil(t, err)
	assert.Equal(t, unexpectedErr.Code, err.Code)
	assert.True(t, strings.Contains(err.Message, unexpectedErr.Message), "Should contains Unexpected Error")
}
