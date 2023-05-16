package pkg

import (
	"encoding/json"
	"errors"

	"github.com/zujisoft/sl/cc/pkg/constants"
	"github.com/zujisoft/sl/cc/pkg/helpers"
	"github.com/zujisoft/sl/cc/pkg/interfaces"
	"github.com/zujisoft/sl/cc/pkg/tctx"
	"github.com/zujisoft/sl/cc/pkg/utils"
)

func (sl *SlContract) RequestBatchTransfer(ctx tctx.SlTransactionContext, id string) interfaces.Batch {
	// Get MSP ID of client
	if mspId, err := ctx.GetClientIdentity().GetMSPID(); err != nil {
		panic(err.Error())
	} else {
		if hashId, err := utils.HashId(id); err != nil {
			panic(err.Error())
		} else {
			// Check hash ID
			batchBytes, err := ctx.GetStub().GetState(*hashId)
			if err != nil {
				panic(err.Error())
			}
			if batchBytes != nil {
				panic(errors.New("batch id is existed"))
			}
			// Interface Batch
			batch := &interfaces.Batch{
				CurrentOwnerOrgId: mspId,
				State:             constants.Request,
			}

			batchJson, err := json.Marshal(batch)
			if err != nil {
				panic(err)
			}

			if err := ctx.GetStub().PutState(*hashId, batchJson); err != nil {
				panic(err)
			}

			return *batch
		}
	}
}

func (sl *SlContract) AcceptBatchTransfer(ctx tctx.SlTransactionContext, id string) interfaces.Batch {
	batch, err := helpers.SetStateForBatchTransfer(ctx, id, constants.Accept)
	if err != nil {
		panic(err)
	}
	return *batch
}

func (sl *SlContract) RefuseBatchTransfer(ctx tctx.SlTransactionContext, id string) interfaces.Batch {
	batch, err := helpers.SetStateForBatchTransfer(ctx, id, constants.Refuse)
	if err != nil {
		panic(err)
	}
	return *batch
}
