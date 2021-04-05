package action

import (
	"encoding/json"
	"net/http"

	"github.com/devbenatti/rbwallet-go/driven/database"
	dao "github.com/devbenatti/rbwallet-go/driven/database/dao/account"
	"github.com/devbenatti/rbwallet-go/driven/uuid"
	"github.com/devbenatti/rbwallet-go/driver/api/action"
	dto "github.com/devbenatti/rbwallet-go/dto/account"
	service "github.com/devbenatti/rbwallet-go/service/account"
)

var _ action.Action = (*CreateAccount)(nil)

type CreateAccount struct {
	uuid    uuid.UuidGenerator
	dao     dao.AccountDAO
	service service.AccountService
}

func NewCreateAccount() CreateAccount {
	return CreateAccount{}
}

type account struct {
	DocumentIdentifier string `json:"documentIdentifier"`
}

func (ca *CreateAccount) Execute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ca.configure()

		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)

		var acc account

		err := decoder.Decode(&acc)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		id := ca.uuid.Generate()

		accountDTO := dto.NewAccountDTO(id.Val(), acc.DocumentIdentifier)

		ca.service.Create(accountDTO)
	}
}

func (ca *CreateAccount) configure() {
	ca.uuid = uuid.NewUuidGenerator()
	db := database.GetConnection()
	ca.dao = dao.NewAccountDAO(db)
	ca.service = service.NewAccountService(ca.dao)
}
