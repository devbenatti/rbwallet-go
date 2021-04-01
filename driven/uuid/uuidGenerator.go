package uuid

import "github.com/devbenatti/rbwallet-go/model/valueObject"

type UuidGenerator interface {
	Generate() valueObject.Uuid
}
