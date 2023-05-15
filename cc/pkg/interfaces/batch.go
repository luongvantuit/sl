package interfaces

type Batch struct {
	Id                   string   `json:"id"`
	ProductId            string   `json:"productId"`
	BatchIngredientIds   []string `json:"batchIngredientIds"`
	Params               []string `json:"params"`
	State                string   `json:"state"`
	CurrentOwnerOrgId    string   `json:"currentOwnerOrgId"`
	CurrentBlockerOrgId  string   `json:"currentBlockerOrgId"`
	CurrentReceiverOrgId string   `json:"currentReceiverOrgId"`
	OutputBatchId        string   `json:"outputBatchId"`
}
