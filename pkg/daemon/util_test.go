package apd

import (
	"testing"

	"github.com/SkycoinProject/skycoin/src/cipher"
)

func TestBroadCastPubKey(t *testing.T) {
	type args struct {
		Packet      []byte
		BroadCastIP string
		Port        int
	}
	getPacketByte := func(packet Packet) []byte {
		d, err := serialize(packet)
		if err != nil {
			t.Fatalf("failed to serialize packet")
		}

		return d
	}

	key, _ := cipher.GenerateKeyPair()
	IP := "127.0.0.1:8080"
	packet := Packet{
		PublicKey: key.Hex(),
		IP:        IP,
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"success",
			args{
				getPacketByte(packet),
				"255.255.255.255",
				3000,
			},
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := BroadCast(tt.args.BroadCastIP, tt.args.Port, tt.args.Packet)
			if (err != nil) != tt.wantErr {
				t.Errorf("BroadCastPubKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
