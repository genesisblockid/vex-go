package token

import eos "github.com/genesisblockid/vex-go"

func NewTransfer(from, to eos.AccountName, quantity eos.Asset, memo string) *eos.Action {
	return &eos.Action{
		Account: AN("vex.token"),
		Name:    ActN("transfer"),
		Authorization: []eos.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(Transfer{
			From:     from,
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

// Transfer represents the `transfer` struct on `vex.token` contract.
type Transfer struct {
	From     eos.AccountName `json:"from"`
	To       eos.AccountName `json:"to"`
	Quantity eos.Asset       `json:"quantity"`
	Memo     string          `json:"memo"`
}
