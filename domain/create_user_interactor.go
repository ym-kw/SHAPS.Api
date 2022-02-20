package domain

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	"shaps.api/domain/exception"
	"shaps.api/entity"
	"shaps.api/infrastructure/external"
	"shaps.api/infrastructure/repository"
)

type CreateUserInteractor struct {
	UserRepository repository.UserRepositoryInterface
	StripeClient   *external.StripeClient
}

func NewCreateUserInteractor(
	r repository.UserRepositoryInterface,
	sc *external.StripeClient,
) *CreateUserInteractor {
	return &CreateUserInteractor{
		UserRepository: r,
		StripeClient:   sc,
	}
}

func (i *CreateUserInteractor) Excecute(c *gin.Context) exception.Wrapper {
	uid, exists := c.Get("userId")
	if !exists {
		e := exception.Wrapper{
			Code:    exception.NotFoundCode,
			Message: exception.NotFoundUserId,
			Err:     errors.New("not found userId"),
		}
		e.Error()
		return e
	}

	su, stripeErr := i.StripeClient.Customers.New(&stripe.CustomerParams{})
	if stripeErr != nil {
		e := exception.Wrapper{
			Code:    exception.InternalServerErrorCode,
			Message: exception.StripeError,
			Err:     stripeErr,
		}
		e.Error()
		return e
	}

	u := entity.User{
		ID:       uid.(string),
		StripeId: su.ID,
	}

	_, err := i.UserRepository.Create(u)

	return err
}
