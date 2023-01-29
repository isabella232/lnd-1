package kvdb

import (
	"git-indra.lan/indra-labs/lnd/lnd/kvdb/postgres"
	"github.com/btcsuite/btclog"
)

// log is a logger that is initialized as disabled.  This means the package will
// not perform any logging by default until a logger is set.
var log = btclog.Disabled

// UseLogger uses a specified Logger to output package logging info.
func UseLogger(logger btclog.Logger) {
	log = logger

	postgres.UseLogger(log)
}
