package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/devbenatti/rbwallet-go/driver/api/validator"
	dto "github.com/devbenatti/rbwallet-go/dto/transaction"
)

var _ Middleware = (*PayloadValidator)(nil)

type PayloadValidator struct {
	validator validator.Validator
}

func NewPayloadValidator() PayloadValidator {
	return PayloadValidator{
		validator: validator.NewJsonValidator(),
	}
}

func (p *PayloadValidator) Execute() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		data, _ := ioutil.ReadAll(r.Body)

		reader := ioutil.NopCloser(bytes.NewReader(data))
		reader1 := ioutil.NopCloser(bytes.NewReader(data))
		decoder := json.NewDecoder(reader)

		var t dto.CreateTransactionDTO

		_ = decoder.Decode(&t)

		_ = p.validator.Validate(t)

		r.Body = reader1
	})
}
