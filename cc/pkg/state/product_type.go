package state

type ProductType struct {
	Name                       string   `json:"name"`
	Type                       string   `json:"type"`
	ProductTypeIngredientNames []string `json:"productTypeIngredientNames"`
	IssuerOrgId                string   `json:"issuerOrgId"`
	State                      string   `json:"state"`
	CurrentBlockerOrgId        string   `json:"currentDisablerOrgId"`
}
