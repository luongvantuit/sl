package state

type Product struct {
	Name                string `json:"name"`
	ProductTypeName     string `json:"productTypeName"`
	IssuerOrgId         string `json:"issuerOrgId"`
	State               string `json:"state"`
	CurrentBlockerOrgId string `json:"currentBlockerOrgId"`
	ApproverOrgId       string `json:"approverOrgId"`
	RefuserOrgId        string `json:"refuserOrgId"`
}
