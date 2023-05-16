package helpers

import (
	"encoding/json"
	"errors"

	"github.com/zujisoft/sl/cc/pkg/interfaces"
	"github.com/zujisoft/sl/cc/pkg/tctx"
	"github.com/zujisoft/sl/cc/pkg/utils"
)

func SetStateForBatchTransfer(ctx tctx.SlTransactionContext, id, state string) (*interfaces.Batch, error) {
	hashId, err := utils.HashId(id)
	if err != nil {
		return nil, err
	}
	batchBytes, err := ctx.GetStub().GetState(*hashId)
	if err != nil {
		return nil, err
	}
	if batchBytes != nil {
		return nil, errors.New("batch not found")
	}
	var batch interfaces.Batch
	if err := json.Unmarshal(batchBytes, &batch); err != nil {
		return nil, err
	}
	batch.State = state
	batchJson, err := json.Marshal(batch)
	if err != nil {
		return nil, err
	}
	if err := ctx.GetStub().PutState(*hashId, batchJson); err != nil {
		return nil, err
	}
	return &batch, nil
}
