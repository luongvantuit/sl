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

func (sl *SlContract) RegisterBatch(ctx tctx.SlTransactionContext, id, productId string, batchIngredientIds, params []string) interfaces.Batch {
	mspId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		panic(err)
	}
	batchHashId, err := utils.HashId(id)
	if err != nil {
		panic(err)
	}
	batchBytes, err := ctx.GetStub().GetState(*batchHashId)
	if err != nil {
		panic(err)
	}
	if batchBytes != nil {
		panic(errors.New("batch id is existed"))
	}
	productHashId, err := utils.HashId(productId)
	if err != nil {
		panic(err)
	}
	productBytes, err := ctx.GetStub().GetState(*productHashId)
	if err != nil {
		panic(err)
	}

	if productBytes == nil {
		panic(errors.New("product not found"))
	}

	batch := &interfaces.Batch{
		Id:                 id,
		ProductId:          productId,
		BatchIngredientIds: batchIngredientIds,
		Params:             params,
		State:              constants.Request,
		CurrentOwnerOrgId:  mspId,
	}

	batchJson, err := json.Marshal(batch)
	if err != nil {
		panic(err)
	}

	if err := ctx.GetStub().PutState(*batchHashId, batchJson); err != nil {
		panic(err)
	}

	return *batch
}

func (sl *SlContract) BlockBatch(ctx tctx.SlTransactionContext, key, id string) {
	if err := helpers.RequireIsRegulatoryDepartment(ctx, &key); err != nil {
		panic(err)
	}

}

func (sl *SlContract) UnblockBatch(ctx tctx.SlTransactionContext) {

}
