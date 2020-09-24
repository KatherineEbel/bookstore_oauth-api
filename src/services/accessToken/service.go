package accessToken

import (
	"fmt"
	"strings"

	"github.com/KatherineEbel/bookstore_oauth-api/src/domain/accessToken"
	"github.com/KatherineEbel/bookstore_oauth-api/src/repository/db"
	"github.com/KatherineEbel/bookstore_oauth-api/src/repository/rest"
	"github.com/KatherineEbel/bookstore_oauth-api/src/utils/errors"
)

type IService interface {
	GetBId(string) (*accessToken.Token, *errors.RestError)
	Create(accessToken.Request) (*accessToken.Token, *errors.RestError)
	UpdateExpirationTime(accessToken.Token) *errors.RestError
}

type service struct {
	tokenRepo db.IDBRepository
	userRepo  rest.IRestUsersRepository
}

func Service(tr db.IDBRepository, ur rest.IRestUsersRepository) IService {
	return &service{tr, ur}
}

func (s *service) Create(r accessToken.Request) (*accessToken.Token, *errors.RestError) {
	if err := r.Validate(); err != nil {
		return nil, err
	}
	// TODO: Support client_credentials and password grant types
	u, err := s.userRepo.Login(r.UserName, r.Password)
	if err != nil {
		return nil, err
	}
	t := accessToken.GetNewAccessToken(u.Id)
	t.Generate()

	if err := s.tokenRepo.Create(t); err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *service) UpdateExpirationTime(t accessToken.Token) *errors.RestError {
	if err := t.Validate(); err != nil {
		return err
	}
	return s.tokenRepo.UpdateExpirationTime(t)
}

func (s *service) GetBId(id string) (*accessToken.Token, *errors.RestError) {
	id = strings.TrimSpace(id)
	if len(id) == 0 {
		return nil, errors.NewBadRequestError(fmt.Sprintf("invalid token id: %s", id))
	}
	return s.tokenRepo.GetById(id)
}
