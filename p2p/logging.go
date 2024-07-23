package p2p

import (
	"fmt"

	"github.com/streamingfast/logging"
)

var zlog, _ = logging.PackageLogger("vex-go", "github.com/genesisblockid/vex-go/p2p")

// SyncLogger sync logger, should `defer SyncLogger()` when use p2p package
func SyncLogger() {
	err := zlog.Sync()
	if err != nil {
		fmt.Printf("unable to sync p2p logger: %s\n", err)
	}
}
