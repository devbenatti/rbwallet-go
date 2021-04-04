package action

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/devbenatti/rbwallet-go/driven/uuid"
	"github.com/devbenatti/rbwallet-go/driver/webApi/action"
	dto "github.com/devbenatti/rbwallet-go/dto/account"
	service "github.com/devbenatti/rbwallet-go/service/account"
)

var _ action.Action = (*CreateAccount)(nil)

type CreateAccount struct {
	uuid    uuid.UuidGenerator
	service service.AccountService
}

func NewCreateAccount(uuid uuid.UuidGenerator, s service.AccountService) CreateAccount {
	return CreateAccount{
		uuid:    uuid,
		service: s,
	}
}

type account struct {
	DocumentIdentifier string `json:"documentIdentifier"`
}

func (ca CreateAccount) Handle() action.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)

		var acc account

		err := decoder.Decode(&acc)

		if err != nil {
			log.Fatal(err)
		}

		id := ca.uuid.Generate()

		accountDTO := dto.NewAccountDTO(id.Val(), acc.DocumentIdentifier)

		ca.service.Create(accountDTO)
	}
}
