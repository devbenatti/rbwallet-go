package infra

import "github.com/devbenatti/rbwallet-go/model"

type UuidGenerator interface {
	Generate() model.Uuid
}
