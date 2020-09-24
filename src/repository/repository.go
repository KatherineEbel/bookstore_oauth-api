package repository

import (
	"github.com/KatherineEbel/bookstore_oauth-api/src/domain/accessToken"
	"github.com/KatherineEbel/bookstore_oauth-api/src/utils/errors"
)

type IRepository interface {
	GetById(string) (*accessToken.Token, *errors.RestError)
	Create(request accessToken.Request) *errors.RestError
	UpdateExpirationTime(token accessToken.Token) *errors.RestError
}
