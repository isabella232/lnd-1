package chainreg

import (
	"git-indra.lan/indra-labs/lnd/lnd/build"
	"github.com/btcsuite/btclog"
)

// Subsystem defines the logging code for this subsystem.
const Subsystem = "CHRE"

// log is a logger that is initialized with the btclog.Disabled logger.
var log btclog.Logger

// The default amount of logging is none.
func init() {
	UseLogger(build.NewSubLogger(Subsystem, nil))
}

// DisableLog disables all logging output.
func DisableLog() {
	UseLogger(btclog.Disabled)
}

// UseLogger uses a specified Logger to output package logging info.
func UseLogger(logger btclog.Logger) {
	log = logger
}
