package token

import eos "github.com/genesisblockid/vex-go"

func NewCreate(issuer eos.AccountName, maxSupply eos.Asset) *eos.Action {
	return &eos.Action{
		Account: AN("vex.token"),
		Name:    ActN("create"),
		Authorization: []eos.PermissionLevel{
			{Actor: AN("vex.token"), Permission: PN("active")},
		},
		ActionData: eos.NewActionData(Create{
			Issuer:        issuer,
			MaximumSupply: maxSupply,
		}),
	}
}

// Create represents the `create` struct on the `vex.token` contract.
type Create struct {
	Issuer        eos.AccountName `json:"issuer"`
	MaximumSupply eos.Asset       `json:"maximum_supply"`
}
