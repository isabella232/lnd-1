//go:build gofuzz
// +build gofuzz

package lnwirefuzz

import (
	"git-indra.lan/indra-labs/lnd/lnd/lnwire"
)

// Fuzz_update_fulfill_htlc is used by go-fuzz.
func Fuzz_update_fulfill_htlc(data []byte) int {
	// Prefix with MsgUpdateFulfillHTLC.
	data = prefixWithMsgType(data, lnwire.MsgUpdateFulfillHTLC)

	// Pass the message into our general fuzz harness for wire messages!
	return harness(data)
}
