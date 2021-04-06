package action

import (
	"encoding/json"
	"net/http"

	"github.com/devbenatti/rbwallet-go/driver/api/action"
	"github.com/devbenatti/rbwallet-go/model/valueObject"
	"github.com/gorilla/mux"
)

var _ action.Action = (*CreateAccount)(nil)

type FindAccount struct {
	accountBehavior
}

func NewFindAccount() FindAccount {
	return FindAccount{}
}

func (ca *FindAccount) Execute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ca.configure()

		vars := mux.Vars(r)

		id := valueObject.NewUuid(vars["id"])

		result, err := ca.service.FindOne(id)

		w.Header().Add("Content-Type", "application/json")

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(action.FormatJSONError(err.Error()))
			return
		}

		_ = json.NewEncoder(w).Encode(result)
	}
}
