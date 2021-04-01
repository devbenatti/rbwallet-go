package uuid

import (
	"github.com/devbenatti/rbwallet-go/model/valueObject"
	uuid "github.com/nu7hatch/gouuid"
)

var _ UuidGenerator = (*UuidAdapter)(nil)

type UuidAdapter struct {
}

func (ud *UuidAdapter) Generate() valueObject.Uuid {
	uuid, _ := uuid.NewV4()

	return valueObject.NewUuid(uuid.String())
}
