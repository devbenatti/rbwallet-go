package action

import (
	"encoding/json"
	"net/http"

	"github.com/devbenatti/rbwallet-go/driver/api/action"
	dto "github.com/devbenatti/rbwallet-go/dto/transaction"
)

var _ action.Action = (*CreateTransaction)(nil)

type CreateTransaction struct {
	transactionBehavior
}

func NewCreateTransaction() CreateTransaction {
	return CreateTransaction{}
}

func (ca *CreateTransaction) Execute() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ca.configure()

		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)

		var t dto.CreateTransactionDTO

		err := decoder.Decode(&t)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		id := ca.uuid.Generate()

		t.Code = id.String()

		ca.service.Create(t)

		response := map[string]string{
			"code": id.String(),
		}

		w.Header().Add("Content-Type", "application/json")

		_ = json.NewEncoder(w).Encode(response)
	})
}
