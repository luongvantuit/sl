package pkg

import ctxPkg "github.com/zujisoft/sl/cc/pkg/context"

const (
	TopRole = "RegulatoryDepartment"
)

func (sl *SlContract) addRoleSet(ctx ctxPkg.SlTransactionContext, role string) {
	var mspId, err = ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		panic(err.Error())
	}

}
