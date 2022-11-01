package solana_test

import (
	"context"
	"log"
	"testing"
	"time"

	sgo "github.com/gagliardetto/solana-go"
	sgorpc "github.com/gagliardetto/solana-go/rpc"
	"github.com/noncepad/go-tutorial-1/solana"
)

func TestWallet(t *testing.T) {
	w, err := solana.CreateWallet()
	if err != nil {
		t.Fatal(err)

	}
	log.Printf("key=%s", w.PrintKey())
	t.Fatal("fail me")
}

func TestAddress(t *testing.T) {
	w, err := solana.CreateWallet()
	if err != nil {
		t.Fatal(err)

	}

	ctx := context.Background()
	rpcClient := sgorpc.New("https://api.devnet.solana.com")

	_, err = rpcClient.RequestAirdrop(ctx, w.PublicKey(), sgo.LAMPORTS_PER_SOL/2, sgorpc.CommitmentFinalized)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	bh, err := rpcClient.GetBalance(ctx, w.PublicKey(), sgorpc.CommitmentFinalized)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("address=%s balance=%d", w.Address(), bh.Value)
	t.Fatal("fail me")
}

func TestLoadWallet(t *testing.T) {
	myPrivateKey := "1mBuFj84ek75zucKeBD2xCYexrDujzqg1KMfTL8K3s6eqUD7QZBduWgQTmWk4xzfhCThyA85P5aSXsrVmMFCWSy"
	w, err := solana.LoadWallet(myPrivateKey)
	if err != nil {
		t.Fatal(err)

	}

	ctx := context.Background()
	// ctx so we do not wait forever
	rpcClient := sgorpc.New("https://api.devnet.solana.com")
	//remote procedure call

	_, err = rpcClient.RequestAirdrop(
		ctx,
		w.PublicKey(),
		sgo.LAMPORTS_PER_SOL/2,
		sgorpc.CommitmentFinalized,
	)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	bh, err := rpcClient.GetBalance(ctx, w.PublicKey(), sgorpc.CommitmentFinalized)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("walletaddress=%s balance=%d", w.Address(), bh.Value)
	t.Fatal("fail me")
}

func TestTransferWallet(t *testing.T) {
	// Transfer from Alice to Bob
	BobPrivateKey := "3Kj5JrzKfLEjJW9xkG4eYB2vGjBbTdG5u6TcAmh7xaftVuFwGpUidDdNpAKCLecojPM4A28RhyjEEyq6zbEUYqu3"
	AlicePrivateKey := "1mBuFj84ek75zucKeBD2xCYexrDujzqg1KMfTL8K3s6eqUD7QZBduWgQTmWk4xzfhCThyA85P5aSXsrVmMFCWSy"
	_, err := solana.LoadWallet(BobPrivateKey)
	if err != nil {
		t.Fatal(err)

	}
	aw, err := solana.LoadWallet(AlicePrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	// type here....
	ctx := context.Background()
	rpcClient := sgorpc.New("https://api.devnet.solana.com")

	// fetch Alice's balance, then send half of her balance to Bob
	{
		result, err := rpcClient.GetBalance(ctx, aw.PublicKey(), sgorpc.CommitmentFinalized)
		if err != nil {
			t.Fatal(err)
		}
		AliceBalance := result.Value
		amount := AliceBalance / 2 // what happens to fractions?  Joel doesn't know, but doesn't care.

		log.Printf("AliceBalance=%d", AliceBalance)
		log.Printf("Alice will send %d to Bob", amount)
		t.Fatal("kill me")
	}
}
