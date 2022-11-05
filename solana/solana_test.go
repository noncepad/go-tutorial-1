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
	// RPC=remote process call;
	// rpcClient talks to a Solana Validator
	rpcClient := sgorpc.New("https://api.devnet.solana.com")

	// get free fake-money from Solana
	_, err = rpcClient.RequestAirdrop(ctx, w.PublicKey(), sgo.LAMPORTS_PER_SOL/2, sgorpc.CommitmentFinalized)
	if err != nil {
		t.Fatal(err)
	}
	// sleep to give time for Solana to process the airdrop request
	time.Sleep(10 * time.Second)
	bh, err := rpcClient.GetBalance(ctx, w.PublicKey(), sgorpc.CommitmentFinalized)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("address=%s balance=%d", w.Address(), bh.Value)
	t.Fatal("fail me")
}

func TestLoadWallet(t *testing.T) {
	clusterUrl := "https://api.devnet.solana.com"
	myPrivateKey := "1mBuFj84ek75zucKeBD2xCYexrDujzqg1KMfTL8K3s6eqUD7QZBduWgQTmWk4xzfhCThyA85P5aSXsrVmMFCWSy"
	w, err := solana.LoadWallet(clusterUrl, myPrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	// ctx so we do not wait forever
	rpcClient := sgorpc.New(clusterUrl)
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
	// type here....
	ctx := context.Background()
	clusterUrl := "https://api.devnet.solana.com"

	// Transfer from Alice to Bob

	alicePrivateKey := "1mBuFj84ek75zucKeBD2xCYexrDujzqg1KMfTL8K3s6eqUD7QZBduWgQTmWk4xzfhCThyA85P5aSXsrVmMFCWSy"
	aw, err := solana.LoadWallet(clusterUrl, alicePrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	bobPrivateKey := "3Kj5JrzKfLEjJW9xkG4eYB2vGjBbTdG5u6TcAmh7xaftVuFwGpUidDdNpAKCLecojPM4A28RhyjEEyq6zbEUYqu3"
	bw, err := solana.LoadWallet(clusterUrl, bobPrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	// fetch Alice's balance, then send half of her balance to Bob
	beforeBalance, err := aw.Balance(ctx)
	if err != nil {
		t.Fatal(err)
	}
	amount := beforeBalance / 2

	// what happens if Transfer fails
	err = aw.Transfer(ctx, amount, bw.PublicKey())
	if err != nil {
		t.Fatal(err)
	}

	afterBalance, err := aw.Balance(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if beforeBalance > afterBalance {
		t.Fatal(err)
	}
}
