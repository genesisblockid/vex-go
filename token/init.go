package token

import eos "github.com/genesisblockid/vex-go"

func init() {
	eos.RegisterAction(AN("vex.token"), ActN("transfer"), Transfer{})
	eos.RegisterAction(AN("vex.token"), ActN("issue"), Issue{})
	eos.RegisterAction(AN("vex.token"), ActN("create"), Create{})
}
