package helpers

import (
	"encoding/json"
	"errors"

	"github.com/zujisoft/sl/cc/pkg/interfaces"
	"github.com/zujisoft/sl/cc/pkg/tctx"
	"github.com/zujisoft/sl/cc/pkg/utils"
)

func SetStateOfRule(ctx tctx.SlTransactionContext, id, state string) (*interfaces.Rule, error) {
	hashId, err := utils.HashId(id)
	if err != nil {
		return nil, err
	}

	ruleBytes, err := ctx.GetStub().GetState(*hashId)

	if err != nil {
		return nil, err
	}
	if ruleBytes == nil {
		return nil, errors.New("rule not found")
	}

	var rule interfaces.Rule

	if err := json.Unmarshal(ruleBytes, &rule); err != nil {
		return nil, err
	}

	rule.State = state

	ruleJson, err := json.Marshal(rule)

	if err != nil {
		return nil, err
	}

	if err := ctx.GetStub().PutState(*hashId, ruleJson); err != nil {
		return nil, err
	}

	return &rule, nil
}
