// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package relayer

import (
	"math/big"
	"testing"
)

// TestExtractAmount tests extractAmountTransferred to extract the total amount
// transferred during the transfer event
func TestExtractAmountTransferred(t *testing.T) {
	// init instance of Message
	msg := &Message{
		Payload: []interface{}{
			big.NewInt(10).Bytes(), // 10 tokens
		},
	}

	payloadAmount, err := msg.extractAmountTransferred()
	if err != nil {
		t.Fatalf("could not extract amount transferred: %v", err)
	}

	if payloadAmount < 1 {
		t.Fatal("amount less than 1")
	}
}
