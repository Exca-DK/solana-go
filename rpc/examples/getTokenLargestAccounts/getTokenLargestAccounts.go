package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.EndpointRPC_MainNetBeta
	client := rpc.New(endpoint)

	pubKey := solana.MustPublicKeyFromBase58("SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt") // serum token
	out, err := client.GetTokenLargestAccounts(
		context.TODO(),
		pubKey,
		rpc.CommitmentType("finalized"),
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}