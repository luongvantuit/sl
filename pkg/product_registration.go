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

func (sl *SlContract) RequestProductRegistration(ctx tctx.SlTransactionContext, productTypeId, productId, productName string) interfaces.Product {
	mspId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		panic(err)
	}
	productTypeHashId, err := utils.HashId(productTypeId)
	if err != nil {
		panic(err)
	}
	productTypeBytes, err := ctx.GetStub().GetState(*productTypeHashId)
	if err != nil {
		panic(err)
	}
	if productTypeBytes == nil {
		panic(errors.New("product type not found"))
	}

	var productTypeJson interfaces.ProductType

	if err := json.Unmarshal(productTypeBytes, &productTypeJson); err != nil {
		panic(err)
	}

	if productTypeJson.Type == Primary {
		if is, err := helpers.HaveRole(ctx, Producer); err != nil {
			panic(err)
		} else {
			if !is {
				panic("role not valid")
			}
		}
	} else if productTypeJson.Type == Derived {
		if is, err := helpers.HaveRole(ctx, Manufacturer); err != nil {
			panic(err)
		} else {
			if !is {
				panic("role not valid")
			}
		}
	} else {
		panic("role not valid")
	}

	productHashId, err := utils.HashId(productId)
	if err != nil {
		panic(err.Error())
	}
	productBytes, err := ctx.GetStub().GetState(*productHashId)
	if err != nil {
		panic(err)
	}
	if productBytes != nil {
		panic(errors.New("product is existed"))
	}

	product := &interfaces.Product{
		IssuerOrgId:   mspId,
		State:         constants.Request,
		Name:          productName,
		ProductTypeId: productTypeId,
	}

	productJson, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	if err := ctx.GetStub().PutState(*productHashId, productJson); err != nil {
		panic(err)
	}
	return *product
}

func (sl *SlContract) AcceptProductRegistration(ctx tctx.SlTransactionContext, key, id string) interfaces.Product {
	if err := helpers.RequireIsRegulatoryDepartment(ctx, &key); err != nil {
		panic(err)
	}
	product, err := helpers.SetStateForProduct(ctx, id, constants.Accept)
	if err != nil {
		panic(err)
	}
	return *product
}

func (sl *SlContract) RefuseProductRegistration(ctx tctx.SlTransactionContext, key, id string) interfaces.Product {
	if err := helpers.RequireIsRegulatoryDepartment(ctx, &key); err != nil {
		panic(err)
	}
	product, err := helpers.SetStateForProduct(ctx, id, constants.Refuse)
	if err != nil {
		panic(err)
	}
	return *product
}
