package pkg

import (
	"encoding/json"

	"github.com/zujisoft/sl/cc/pkg/interfaces"
	"github.com/zujisoft/sl/cc/pkg/tctx"
	"github.com/zujisoft/sl/cc/pkg/utils"
)

func (sl *SlContract) GetBatchHistory(ctx tctx.SlTransactionContext, id string) (interfaces.Batch, error) {
	if hashId, err := utils.HashId(id); err != nil {
		panic(err.Error())
	} else {
		if baHisBts, err := ctx.GetStub().GetState(*hashId); err != nil {
			panic(err.Error())
		} else {
			var batchHistory interfaces.Batch
			if err := json.Unmarshal(baHisBts, &batchHistory); err != nil {
				panic(err.Error())
			}
			return batchHistory, nil
		}
	}

}
