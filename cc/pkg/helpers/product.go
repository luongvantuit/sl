package helpers

import (
	"encoding/json"
	"errors"

	"github.com/zujisoft/sl/cc/pkg/constants"
	"github.com/zujisoft/sl/cc/pkg/interfaces"
	"github.com/zujisoft/sl/cc/pkg/tctx"
	"github.com/zujisoft/sl/cc/pkg/utils"
)

func SetStateForProduct(ctx tctx.SlTransactionContext, id, state string) (*interfaces.Product, error) {
	mspId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return nil, err
	}
	hashId, err := utils.HashId(id)
	if err != nil {
		return nil, err
	}
	productBytes, err := ctx.GetStub().GetState(*hashId)
	if err != nil {
		return nil, err
	}
	if productBytes != nil {
		return nil, errors.New("product not found")
	}
	var product interfaces.Product
	if err := json.Unmarshal(productBytes, &product); err != nil {
		return nil, err
	}

	if state != constants.Block {
		if state == constants.Unblock {
			if product.State == constants.Block {
				if product.ApproverOrgId != "" {
					product.State = constants.Accept
				} else if product.RefuserOrgId != "" {
					product.State = constants.Refuse
				} else {
					product.State = constants.Request
				}
				product.CurrentBlockerOrgId = ""
			}
		} else {
			if product.CurrentBlockerOrgId != "" {
				return nil, errors.New("product current is not available")
			}

			if product.ApproverOrgId != "" {
				return nil, errors.New("product current is accepted")
			}

			if product.RefuserOrgId == "" {
				return nil, errors.New("product current is refused")
			}

			product.State = state

			if state == constants.Accept {
				product.ApproverOrgId = mspId
			} else if state == constants.Refuse {
				product.RefuserOrgId = mspId
			}
		}
	} else {
		product.State = constants.Block
		product.CurrentBlockerOrgId = mspId
	}

	productJson, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}
	if err := ctx.GetStub().PutState(*hashId, productJson); err != nil {
		return nil, err
	}
	return &product, nil
}
