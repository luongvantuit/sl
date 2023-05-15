package helpers

import (
	"encoding/json"
	"errors"

	"github.com/zujisoft/sl/cc/pkg/interfaces"
	"github.com/zujisoft/sl/cc/pkg/tctx"
	"github.com/zujisoft/sl/cc/pkg/utils"
)

func SetStateForProductType(ctx tctx.SlTransactionContext, id, state string) (*interfaces.ProductType, error) {
	hashId, err := utils.HashId(id)
	if err != nil {
		return nil, err
	}
	productTypeBytes, err := ctx.GetStub().GetState(*hashId)
	if err != nil {
		return nil, err
	}
	if productTypeBytes != nil {
		return nil, errors.New("product type not found")
	}
	var productType interfaces.ProductType
	if err := json.Unmarshal(productTypeBytes, &productType); err != nil {
		return nil, err
	}
	productType.State = state
	productTypeJson, err := json.Marshal(productType)
	if err != nil {
		return nil, err
	}
	if err := ctx.GetStub().PutState(*hashId, productTypeJson); err != nil {
		return nil, err
	}
	return &productType, nil
}
