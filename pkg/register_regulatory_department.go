package pkg

import (
	"errors"

	"github.com/zujisoft/sl/cc/pkg/tctx"
)

func (sl *SlContract) RegisterRegulatoryDepartment(ctx tctx.SlTransactionContext, key string) {
	mspId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		panic(err.Error())
	}

	mspIdOfRDBytes, err := ctx.GetStub().GetPrivateData("RDMSPID", key)
	if err != nil {
		panic(err.Error())
	}

	if mspIdOfRDBytes != nil {
		panic(errors.New("rule regulatory depart is registered"))
	}

	if err := ctx.GetStub().PutPrivateData("RDMSPID", key, []byte(mspId)); err != nil {
		panic(err.Error())
	}
}
