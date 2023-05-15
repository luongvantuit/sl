package pkg

import (
	"encoding/json"
	"errors"

	"github.com/zujisoft/sl/cc/pkg/helpers"
	"github.com/zujisoft/sl/cc/pkg/interfaces"
	"github.com/zujisoft/sl/cc/pkg/tctx"
	"github.com/zujisoft/sl/cc/pkg/utils"
)

func (sl *SlContract) AddRule(ctx tctx.SlTransactionContext, key, id, productTypeId, ruleJsonValue string) interfaces.Rule {
	if err := helpers.RequireIsRegulatoryDepartment(ctx, &key); err != nil {
		panic(err)
	}
	hashId, err := utils.HashId(id)
	if err != nil {
		panic(err)
	}
	ruleBytes, err := ctx.GetStub().GetState(*hashId)
	if err != nil {
		panic(err)
	}
	if ruleBytes != nil {
		panic(errors.New("rule id is existed"))
	}

	mspId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		panic(err)
	}

	productTypeHashId, err := utils.HashId(productTypeId)
	if err != nil {
		panic(err)
	}
	productTypeJson, err := ctx.GetStub().GetState(*productTypeHashId)
	if err != nil {
		panic(err)
	}
	if productTypeJson == nil {
		panic(errors.New("product type not found"))
	}
	var productType interfaces.ProductType
	if err := json.Unmarshal(productTypeJson, &productType); err != nil {
		panic(err)
	}
	rule := &interfaces.Rule{
		Id:              id,
		ProductTypeName: productType.Name,
		JsonValue:       ruleJsonValue,
		IssuerOrgId:     mspId,
		State:           Enable,
	}
	ruleJson, err := json.Marshal(rule)
	if err != nil {
		panic(err)
	}
	if err := ctx.GetStub().PutState(*hashId, ruleJson); err != nil {
		panic(err)
	}

	return *rule
}

func (sl *SlContract) EnableRule(ctx tctx.SlTransactionContext, key, id string) *interfaces.Rule {
	if err := helpers.RequireIsRegulatoryDepartment(ctx, &key); err != nil {
		panic(err)
	}
	rule, err := helpers.SetStateOfRule(ctx, id, Enable)
	if err != nil {
		panic(err)
	}
	return rule
}

func (sl *SlContract) DisableRule(ctx tctx.SlTransactionContext, key, id string) *interfaces.Rule {
	if err := helpers.RequireIsRegulatoryDepartment(ctx, &key); err != nil {
		panic(err)
	}
	rule, err := helpers.SetStateOfRule(ctx, id, Disable)
	if err != nil {
		panic(err)
	}
	return rule
}
