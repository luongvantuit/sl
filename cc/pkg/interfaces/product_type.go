package interfaces

type ProductType struct {
	Name                       string   `json:"name" required:"true"`
	Type                       string   `json:"type"`
	ProductTypeIngredientNames []string `json:"productTypeIngredientNames"`
	IssuerOrgId                string   `json:"issuerOrgId"`
	State                      string   `json:"state" default:"normal"`
	CurrentBlockerOrgId        string   `json:"currentDisablerOrgId"`
}
