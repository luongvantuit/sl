package pkg

import (
	"github.com/zujisoft/sl/cc/pkg/helpers"
	"github.com/zujisoft/sl/cc/pkg/interfaces"
	"github.com/zujisoft/sl/cc/pkg/tctx"
)

func (sl *SlContract) BlockProduct(ctx tctx.SlTransactionContext, key, id string) interfaces.Product {
	if err := helpers.RequireIsRegulatoryDepartment(ctx, &key); err != nil {
		panic(err)
	}
	product, err := helpers.SetStateForProduct(ctx, id, Block)
	if err != nil {
		panic(err)
	}

	return *product
}

func (sl *SlContract) UnblockProduct(ctx tctx.SlTransactionContext, key, id string) interfaces.Product {
	if err := helpers.RequireIsRegulatoryDepartment(ctx, &key); err != nil {
		panic(err)
	}
	product, err := helpers.SetStateForProduct(ctx, id, Block)
	if err != nil {
		panic(err)
	}

	return *product
}
