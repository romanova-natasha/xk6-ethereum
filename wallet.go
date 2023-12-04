package ethereum

import (
	"encoding/hex"

	"github.com/umbracle/ethgo/wallet"
	"go.k6.io/k6/js/modules"
)

type Wallet struct{}

type Key struct {
	PrivateKey string
	Address    string
}

func init() {
	wallet := Wallet{}

	modules.Register("k6/x/ethereum/wallet", &wallet)
}

// GenerateKey key creates a random key
func (w *Wallet) GenerateKey() (*Key, error) {
	k, err := wallet.GenerateKey()
	if err != nil {
		return nil, err
	}
	pk, err := k.MarshallPrivateKey()
	if err != nil {
		return nil, err
	}
	pks := hex.EncodeToString(pk)

	return &Key{
		PrivateKey: pks,
		Address:    k.Address().String(),
	}, err
}

func (*Wallet) NewWalletKeyFromPrivateKey(privateKey string) (*wallet.Key, error) {
	pk, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}

	w, err = wallet.NewWalletFromPrivateKey(pk)
	if err != nil {
		return nil, err
	}
	
	return w.Key, err
}
