package web3

import (
	"context"
	"fmt"
	"os"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

func GetATA(wallet string, mint string) (solana.PublicKey, uint8, error) {
	return solana.FindAssociatedTokenAddress(solana.MustPublicKeyFromBase58(wallet), solana.MustPublicKeyFromBase58(mint))
}

func GetMintOwner(address string) (*string, error) {
	client := rpc.New(os.Getenv("RPC_ENDPOINT"))

	mint := solana.MustPublicKeyFromBase58(address)
	out, err := client.GetTokenLargestAccounts(
		context.TODO(),
		mint,
		rpc.CommitmentConfirmed,
	)

	if err != nil {
		return nil, err
	}

	for _, account := range out.Value {
		if account.Amount == "1" {
			tokenAcct := account.Address
			tokenOut, err := client.GetAccountInfo(
				context.TODO(),
				tokenAcct,
			)

			if err != nil {
				return nil, err
			}

			dec := bin.NewBinDecoder(tokenOut.Value.Data.GetBinary())
			acct := &token.Account{}
			err = dec.Decode(acct)

			if err != nil {
				return nil, err
			}

			o := acct.Owner.String()
			return &o, nil
		}
	}

	return nil, fmt.Errorf("no owner found")
}
