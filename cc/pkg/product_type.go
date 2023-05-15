package pkg

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/zujisoft/sl/cc/pkg/helpers"
	"github.com/zujisoft/sl/cc/pkg/interfaces"
	"github.com/zujisoft/sl/cc/pkg/tctx"
	"github.com/zujisoft/sl/cc/pkg/utils"
)

// Typ = Type primary & derived
func (sl *SlContract) AddProductType(ctx tctx.SlTransactionContext, key, id string, productTypeName, typ string, productTypeIngredientNames []string) interfaces.ProductType {
	if err := helpers.RequireIsRegulatoryDepartment(ctx, &key); err != nil {
		panic(err)
	}
	typeLower := strings.ToLower(typ)
	if !((typeLower == "primary" && len(productTypeIngredientNames) == 0) || (typeLower == "derived" && len(productTypeIngredientNames) >= 1)) {
		panic(errors.New("arguments is not valid"))
	}
	mspId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		panic(err)
	}

	hashId, err := utils.HashId(id)
	if err != nil {
		panic(err)
	}
	productType := &interfaces.ProductType{
		Type:                       typ,
		Name:                       productTypeName,
		ProductTypeIngredientNames: productTypeIngredientNames,
		IssuerOrgId:                mspId,
		State:                      Normal,
	}
	productTypeJson, err := json.Marshal(productType)
	if err != nil {
		panic(err)
	}
	if err := ctx.GetStub().PutState(*hashId, productTypeJson); err != nil {
		panic(err)
	}

	return *productType
}

func (sl *SlContract) BlockProductType(ctx tctx.SlTransactionContext, key, id string) interfaces.ProductType {
	if err := helpers.RequireIsRegulatoryDepartment(ctx, &key); err != nil {
		panic(err)
	}
	productType, err := helpers.SetStateForProductType(ctx, id, Block)
	if err != nil {
		panic(err)
	}
	return *productType
}

func (sl *SlContract) UnblockProductType(ctx tctx.SlTransactionContext, key, id string) interfaces.ProductType {
	productType, err := helpers.SetStateForProductType(ctx, id, Normal)
	if err != nil {
		panic(err)
	}
	return *productType
}
