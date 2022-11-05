# Solana in Go
# ?????golang
# Solana?

This blurb is a summary of [the article on Wikipedia](https://en.wikipedia.org/wiki/Solana_(blockchain_platform)).

> Solana is a peer to peer blockchain.  Peers run a network client called a Validator.  People move money by cryptographically signing financial transactions.  Network clients called Validators put transactions into bundles called blocks and distribute those blocks.

The monetory unit for Solana is SOL with Lamports being the smallest unit.

```go
const sgo.LAMPORTS_PER_SOL uint64 = 1000000000
```

Funds are stored [in accounts](https://docs.solana.com/developing/programming-model/accounts).  Accounts are identified by 256-bit public key.  Most public keys are derived from private keys and are on the [ed25519 curve](https://www.cryptopp.com/wiki/Ed25519).


# Setup

## Named Package

*Package* defines a group of structs, functions, and constants. An example is:

```go
package solana
```

Packages with the `_test` suffix hold tests (see testing section under Implementation).

```go
package solana_test
```

## Importing Packages

Import imports essential packages to enable the code to run.

```go
import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/gagliardetto/bin"
	sgo "github.com/gagliardetto/solana-go"
)
```
* _context_ is a package to managing goroutines
* _log_ is a package for printing to stderr
* _testing_ is a package to support testing in golang
* _time_ is for tracking time, `Sleep` 


### Deconflicting packages

If two packages import under the same name, change one package's name manually.

```go
import (
	sgo "github.com/gagliardetto/solana-go"
	anothersgo "github.com/anotherguy/solana-go"
)

func main(){
	// not solana.NewRandomPrivateKey()
	key,err:=sgo.NewRandomPrivateKey()

	history,err:=anothersgo.GetHistory()
```



# Stories

This program is to implement a wallet.

A user (human) will store funds in this wallet.  The funds live in the Solana blockchain.

The user will do the following with the wallet:
1. Open an account to store funds
1. Deposit funds
1. Hold funds **securely**
1. Show balance
1. Transfer funds
1. Show transaction history

In this document, `Alice` , `Bob`, and `Eve` will represent humans.

## Open an Account (and Hold funds)

Alice needs an account in Solana.

To do this, she will:

```go
wallet, err := oursolana.CreateWallet()
```
Inside that function, the following code is called:

```go
key, err := sgo.NewRandomPrivateKey()
```

Alice needs to keep this `key` secure.



## Deposit Funds from Airdrop

Alice will deposit *5000000* LAMPORTS (`uint64`) [into her account](https://docs.solana.com/developing/programming-model/accounts).  Since she is expirementing with Solana, she would like to get free [fake-money from a faucet](https://101blockchains.com/crypto-faucet/).

```go
_, err = rpcClient.RequestAirdrop(
    ctx, // context b/c we are making network request
    wallet.PublicKey(),  // Alice's "account number"
    5000000, // monetary unit= LAMPORTS
    sgorpc.CommitmentFinalized, // probabilty of the transaction being reversed (finalized=very low)
)
```
The airdrop function transfers money from a "faucet" account to Alice's account.  Airdrops are rate-limited.

## Show Balance

Alice will check the balance of LAMPORTS in her account.

To do this, she will:

```go
amount, err := wallet.Balance(ctx)
```
> she will call the Balance function in her wallet instance.

Alice's wallet asks the validator for the balance via RPC (remote procedure call):

```go
result, err := rpcClient.GetBalance(
	ctx, 
	wallet.PublicKey(), 
    sgorpc.CommitmentFinalized,
)
```

## Transfer Funds

Alice will tranfer half of her funds from her account to Bob's account. In order to do this Alice will need her private key and Bob's public key.  She will:
1. Create a transfer instruction
1. Assemble a transaction
1. Sign the transaction with her private key
1. Send transaction


### Transfer Instruction

```go
transferInstruction:=sgosys.NewTransferInstructionBuilder()
```

### Assemble Transaction

```go
txBuilder := sgo.NewTransactionBuilder()
//....
txBuilder.AddInstruction(transferInstruction.Build())
//....
tx, err := txBuilder.Build()
```

> ????? txfee
 
### Sign Transaction

```go
_, err = tx.Sign(func(pubkey sgo.PublicKey) *sgo.PrivateKey{
	if wallet.PublicKey().Equals(pubkey){
		key:=wallet.Key()
		return &key
	} else {
		return nil
	}
})
```

> A callback is a variable that contains a function definition. See example below:

```go
myCallback:=func(a int,b int) int{
	return a+b
}
```
### Send Transaction

Alice will send her funds from her accont to Bob's account.

To do this, she will:

```go
signature, err := rpcClient.SendTransaction(ctx, tx)
```
* *signature* in Solana means transaction ID, not the cryptographic signature

To print the signature, she will:
```go
//prints Alice's signature as a string
log.Printf("AliceSig=%s", sig.String()) 
```

## Show Transaction History 

tbd



# Implementation

* wallet struct and method
* testing
* error handling (risk management)
* logging



```go
func TestWallet(t *testing.T) {
	w, err := solana.CreateWallet()
	if err != nil {
		t.Fatal(err)

	}
	log.Printf("key=%s", w.PrintKey())
	t.Fatal("fail me")
}
```

## Error Handling

For testing:

### t.Fatal

If or when Alice encounters an error t.Fatal is used to stop the test from continuing.

```go
if a > 2{
	t.Fatal(err)
}
```

## Logging

[Logging](https://pkg.go.dev/log) is used for validating if a program is running correctly.

1. Print strings `%s`
1. Print integers `%d`
1. Print floats ????

prints Alice's signature as a string 
```go		
log.Printf("AliceSig=%s", sig.String()) 
```