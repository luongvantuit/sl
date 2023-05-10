package state

type Rule struct {
	Id                   string      `json:"id"`
	ProductTypeName      string      `json:"productTypeName"`
	JsonValue            interface{} `json:"jsonValue"`
	IssuerOrgId          string      `json:"issuerOrgId"`
	State                string      `json:"state"`
	CurrentDisablerOrgId string      `json:"currentDisablerOrgId"`
}
