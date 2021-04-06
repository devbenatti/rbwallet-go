package action

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devbenatti/rbwallet-go/driver/api/action"
	dto "github.com/devbenatti/rbwallet-go/dto/account"
)

var _ action.Action = (*CreateAccount)(nil)

type CreateAccount struct {
	accountBehavior
}

func NewCreateAccount() CreateAccount {
	return CreateAccount{}
}

func (ca *CreateAccount) Execute() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ca.configure()

		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)

		var acc dto.AccountDTO

		err := decoder.Decode(&acc)

		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		id := ca.uuid.Generate()

		accountDTO := dto.NewAccountDTO(id.String(), acc.DocumentIdentifier)

		ca.service.Create(accountDTO)

		response := map[string]string{
			"id": id.String(),
		}

		w.Header().Add("Content-Type", "application/json")

		_ = json.NewEncoder(w).Encode(response)
	})
}
