package interfaces

type Product struct {
	Name                string `json:"name"`
	ProductTypeId       string `json:"productTypeId"`
	IssuerOrgId         string `json:"issuerOrgId"`
	State               string `json:"state"`
	CurrentBlockerOrgId string `json:"currentBlockerOrgId"`
	ApproverOrgId       string `json:"approverOrgId"`
	RefuserOrgId        string `json:"refuserOrgId"`
}
