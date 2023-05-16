package pkg

import (
	"encoding/json"

	"github.com/zujisoft/sl/cc/pkg/helpers"
	"github.com/zujisoft/sl/cc/pkg/interfaces"
	"github.com/zujisoft/sl/cc/pkg/tctx"
	"github.com/zujisoft/sl/cc/pkg/utils"
)

func (sl *SlContract) AddRoleSet(ctx tctx.SlTransactionContext, key, orgId string, roles []string) bool {
	if err := helpers.RequireIsRegulatoryDepartment(ctx, &key); err != nil {
		panic(err)
	}
	hashSysId, err := utils.HashSysId(orgId)
	if err != nil {
		panic(err)
	}
	org := &interfaces.RoleSet{
		OrgId: orgId,
		Roles: roles,
	}
	orgJson, err := json.Marshal(org)
	if err != nil {
		panic(err)
	}
	if err := ctx.GetStub().PutState(*hashSysId, orgJson); err != nil {
		panic(err)
	}
	return true
}
