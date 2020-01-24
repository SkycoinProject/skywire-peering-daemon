package apd

import (
	"testing"

	"github.com/SkycoinProject/skycoin/src/cipher"
)

func TestBroadCastPubKey(t *testing.T) {
	type args struct {
		PublicKey   string
		BroadCastIP string
		Port        int
	}
	publicKey := func() string {
		key, _ := cipher.GenerateKeyPair()
		return key.Hex()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"success",
			args{
				publicKey(),
				"255.255.255.255",
				3000,
			},
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := BroadCastPubKey(tt.args.PublicKey, tt.args.BroadCastIP, tt.args.Port)
			if (err != nil) != tt.wantErr {
				t.Errorf("BroadCastPubKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
