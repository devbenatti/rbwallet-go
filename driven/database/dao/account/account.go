package dao

import (
	"database/sql"

	dto "github.com/devbenatti/rbwallet-go/dto/account"
	"github.com/devbenatti/rbwallet-go/model/valueObject"
)

var _ AccountDAO = (*Account)(nil)

type Account struct {
	db *sql.DB
}

func NewAccountDAO(conn *sql.DB) *Account {
	return &Account{
		db: conn,
	}
}

func (a *Account) Create(ac dto.AccountDTO) {
	defer a.db.Close()

	insert, err := a.db.Prepare("INSERT INTO account(id, document_identifier) VALUES(?,?)")
	if err != nil {
		panic(err)
	}

	id := ac.ID()

	_, err = insert.Exec(id.Val(), ac.DocumentIdentifier())

	if err != nil {
		panic(err)
	}
}

func (a *Account) FindByID(id valueObject.Uuid) *dto.AccountDTO {
	defer a.db.Close()

	var accountID, documentIdentifier string

	query := a.db.QueryRow("SELECT * FROM account where id = ?", id.Val())

	err := query.Scan(&accountID, &documentIdentifier)

	if err != nil {
		return nil
	}

	result := dto.NewAccountDTO(accountID, documentIdentifier)

	return &result
}
