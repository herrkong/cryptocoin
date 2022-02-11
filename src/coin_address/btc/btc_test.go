package btc

import (
	"testing"

	"github.com/btcsuite/btcd/btcec"
)

func TestBtc_GetBitcoinAddress(t *testing.T) {
	type args struct {
		pubkey *btcec.PublicKey
		chain  string
	}
	tests := []struct {
		name        string
		b           *Btc
		args        args
		wantAddress string
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Btc{}
			gotAddress, err := b.GetBitcoinAddress(tt.args.pubkey, tt.args.chain)
			if (err != nil) != tt.wantErr {
				t.Errorf("Btc.GetBitcoinAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAddress != tt.wantAddress {
				t.Errorf("Btc.GetBitcoinAddress() = %v, want %v", gotAddress, tt.wantAddress)
			}
		})
	}
}
