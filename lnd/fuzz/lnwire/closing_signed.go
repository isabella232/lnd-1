//go:build gofuzz
// +build gofuzz

package lnwirefuzz

import (
	"git-indra.lan/indra-labs/lnd/lnd/lnwire"
)

// Fuzz_closing_signed is used by go-fuzz.
func Fuzz_closing_signed(data []byte) int {
	// Prefix with MsgClosingSigned.
	data = prefixWithMsgType(data, lnwire.MsgClosingSigned)

	// Pass the message into our general fuzz harness for wire messages!
	return harness(data)
}
