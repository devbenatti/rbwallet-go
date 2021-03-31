package driven

import (
	"github.com/devbenatti/rbwallet-go/model"
	uuid "github.com/nu7hatch/gouuid"
)

var _ UuidGenerator = (*UuidAdapter)(nil)

type UuidAdapter struct {
}

func (ud *UuidAdapter) Generate() model.Uuid {
	uuid, _ := uuid.NewV4()

	return model.NewUuid(uuid.String())
}
