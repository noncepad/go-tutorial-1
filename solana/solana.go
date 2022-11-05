package solana

import (
	"context"
	"log"

	sgo "github.com/gagliardetto/solana-go"
	sgosys "github.com/gagliardetto/solana-go/programs/system"
	sgorpc "github.com/gagliardetto/solana-go/rpc"
)

type Wallet struct {
	rpc *sgorpc.Client
	key sgo.PrivateKey
}

func LoadWallet(url string, privateKey string) (*Wallet, error) {

	key, err := sgo.PrivateKeyFromBase58(privateKey)
	if err != nil {
		return nil, err

	}

	log.Printf("key=%s", key.String())

	w := new(Wallet)
	w.key = key
	w.rpc = sgorpc.New(url)
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

func (w Wallet) Key() sgo.PrivateKey {
	return w.key
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

func (w Wallet) Balance(ctx context.Context) (balance uint64, err error) {
	rpcClient := w.rpc

	result, err := rpcClient.GetBalance(ctx, w.PublicKey(), sgorpc.CommitmentFinalized)
	if err != nil {
		return
	}
	balance = result.Value

	return
}

func (w Wallet) Transfer(ctx context.Context, lamportsBeingTransfered uint64, recipientPublicKey sgo.PublicKey) (err error) {
	var tx *sgo.Transaction // default value is nil, so we don't need to set

	keyMap := make(map[string]sgo.PrivateKey) // map from publickey to privatekey

	// create instruction to transfer money
	transferInstruction := sgosys.NewTransferInstructionBuilder()
	transferInstruction.SetFundingAccount(w.PublicKey())
	keyMap[w.PublicKey().String()] = w.Key()
	transferInstruction.SetRecipientAccount(recipientPublicKey)
	transferInstruction.SetLamports(lamportsBeingTransfered)

	// set up tx
	txBuilder := sgo.NewTransactionBuilder()
	txBuilder.SetFeePayer(w.PublicKey())
	keyMap[w.PublicKey().String()] = w.Key()                                 // this line is not necessary as it is a duplicate, but it is added here for brevity
	result, err := w.rpc.GetLatestBlockhash(ctx, sgorpc.CommitmentFinalized) //idication trasation done recently
	if err != nil {
		return
	}
	txBuilder.SetRecentBlockHash(result.Value.Blockhash)

	// add instructions
	txBuilder.AddInstruction(transferInstruction.Build())

	tx, err = txBuilder.Build()
	if err != nil {
		return
	}
	_, err = tx.Sign(func(pubkey sgo.PublicKey) *sgo.PrivateKey {
		privkey, present := keyMap[pubkey.String()]
		if present {
			return &privkey
		} else {
			return nil
		}
	})
	if err != nil {
		return
	}

	if tx != nil {

		var sig sgo.Signature
		sig, err = w.rpc.SendTransaction(ctx, tx)

		if err != nil {
			return
		}
		log.Printf("AliceSig=%s", sig.String())
		return

	} else {
		return
	}
}
