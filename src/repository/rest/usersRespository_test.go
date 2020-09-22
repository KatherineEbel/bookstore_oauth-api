package rest

import (
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	_ "github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer() // Changed mockup.go in library to make work in current go version
	os.Exit(m.Run())
}

func TestLogin_Timeout(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email": "email@foo.com", "password": "password"}`,
		RespHTTPCode: 0,
		RespBody:     `{}`,
	})
	repo := restUsersRepository{}
	user, err := repo.Login("email@foo.com", "password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)
}

func TestLogin_InvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email": "email@foo.com", "password": "password"}`,
		RespHTTPCode: http.StatusInternalServerError,
		RespBody:     `{"message": "invalid login credentials", "code": "404"}`,
	})
	repo := restUsersRepository{}
	user, err := repo.Login("email@foo.com", "password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)

}

func TestLogin_InvalidCredentials(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email": "email@foo.com", "password": "password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials", "code": 404}`,
	})
	repo := restUsersRepository{}
	user, err := repo.Login("email@foo.com", "password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Code)

}

func TestLogin_InvalidJSONResponse(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email": "email@foo.com", "password": "password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "1", "first_name": "Jane", "last_name": "Doe", "email": "email@foo.com""}`,
	})
	repo := restUsersRepository{}
	user, err := repo.Login("email@foo.com", "password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)
}

func TestLogin_NoError(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email": "email@foo.com", "password": "password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 1, "first_name": "Jane", "last_name": "Doe", "email": "email@foo.com"}`,
	})
	repo := restUsersRepository{}
	user, err := repo.Login("email@foo.com", "password")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, "email@foo.com", user.Email)
}
