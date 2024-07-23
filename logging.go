package vex

import (
	"fmt"

	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

var zlog, tracer = logging.PackageLogger("vex-go", "github.com/genesisblockid/vex-go")

type logStringerFunc func() string

func (f logStringerFunc) String() string { return f() }

func typeField(field string, v interface{}) zap.Field {
	return zap.Stringer(field, logStringerFunc(func() string {
		return fmt.Sprintf("%T", v)
	}))
}
