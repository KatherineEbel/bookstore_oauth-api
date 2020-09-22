package rest

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"
	_ "github.com/mercadolibre/golang-restclient/rest"

	"github.com/KatherineEbel/bookstore_oauth-api/src/domain/users"
	"github.com/KatherineEbel/bookstore_oauth-api/src/utils/errors"
)

var (
	restClient *rest.RequestBuilder
)

func init() {
	restClient = &rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 10 * time.Millisecond,
	}
}

type IRestUsersRepository interface {
	Login(string, string) (*users.User, *errors.RestError)
}

type restUsersRepository struct {
}

func (r *restUsersRepository) Login(e string, p string) (*users.User, *errors.RestError) {
	req := users.LoginRequest{
		Email:    e,
		Password: p,
	}
	return doLogin(req)
}

func Repository() *restUsersRepository {
	return &restUsersRepository{}
}

func doLogin(req users.LoginRequest) (*users.User, *errors.RestError) {
	res := restClient.Post("/users/login", req)
	if res == nil || res.Response == nil {
		return nil, errors.NewInternalServerError("response timeout for login")
	}
	if res.StatusCode > 299 {
		fmt.Println(res.String())
		var re errors.RestError
		err := json.Unmarshal(res.Bytes(), &re)
		if err != nil {
			return nil, errors.NewInternalServerError("unknown error type")
		}
		return nil, &re
	}
	var u users.User
	if err := json.Unmarshal(res.Bytes(), &u); err != nil {
		return nil, errors.NewInternalServerError("unknown user type")
	}
	return &u, nil
}
