package solana

import (
	"log"

	sgo "github.com/gagliardetto/solana-go"
)

type Wallet struct {
	key sgo.PrivateKey
}

func LoadWallet(privateKey string) (*Wallet, error) {

	key, err := sgo.PrivateKeyFromBase58(privateKey)
	if err != nil {
		return nil, err

	}

	log.Printf("key=%s", key.String())

	w := new(Wallet)
	w.key = key
	return w, nil
}

func CreateWallet() (*Wallet, error) {
	key, err := sgo.NewRandomPrivateKey()
	if err != nil {
		return nil, err

	}

	log.Printf("key=%s", key.String())

	w := new(Wallet)
	w.key = key
	return w, nil

}

func (w Wallet) PrintKey() string {
	return w.key.String()
}

// put the public key here
func (w Wallet) Address() string {
	return w.key.PublicKey().String()
}

func (w Wallet) PublicKey() sgo.PublicKey {
	return w.key.PublicKey()
}
