package helpers

import (
	"encoding/json"
	"errors"

	"github.com/zujisoft/sl/cc/pkg/interfaces"
	"github.com/zujisoft/sl/cc/pkg/tctx"
	"github.com/zujisoft/sl/cc/pkg/utils"
)

func IsRegulatoryDepartment(ctx tctx.SlTransactionContext, key *string) (bool, error) {
	if key == nil {
		return false, errors.New("require key")
	}
	mspId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return false, nil
	}
	// Get private data by key RDMSPID
	mspIdOfRD, err := ctx.GetStub().GetPrivateData("RDMSPID", *key)
	if err != nil {
		return false, err
	}
	if mspIdOfRD == nil {
		return false, errors.New("please register rule regulatory department")
	}
	return mspId == string(mspIdOfRD), nil
	// if mspId == string(mspIdOfRD) {
	// 	return true, nil
	// }
	// return false, errors.New("you is not regulatory department")
}

func RequireIsRegulatoryDepartment(ctx tctx.SlTransactionContext, key *string) error {
	if isRD, err := IsRegulatoryDepartment(ctx, key); err != nil {
		if !isRD {
			return errors.New("you is not regulatory department")
		}
	}
	return nil
}

func HaveRole(ctx tctx.SlTransactionContext, role string) (bool, error) {
	mspId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return false, err
	}
	hashSysId, err := utils.HashSysId(mspId)

	if err != nil {
		return false, err
	}

	orgBytes, err := ctx.GetStub().GetState(*hashSysId)
	if err != nil {
		return false, err
	}

	if orgBytes == nil {
		return false, errors.New("require was role set")
	}
	var org interfaces.RoleSet
	if err := json.Unmarshal(orgBytes, &org); err != nil {
		return false, err
	}

	for counter := 0; counter < len(org.Roles); counter++ {
		if org.Roles[counter] == role {
			return true, nil
		}
	}

	return false, nil
}
