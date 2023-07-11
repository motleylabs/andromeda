package web3

import (
	"context"
	"log"
	"os"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

func GetTokensByOwner(address string) ([]string, error) {
	client := rpc.New(os.Getenv("RPC_ENDPOINT"))

	wallet := solana.MustPublicKeyFromBase58(address)
	tokenProgramID := solana.MustPublicKeyFromBase58(TOKEN_PROGRAM_ID)
	out, err := client.GetTokenAccountsByOwner(
		context.TODO(),
		wallet,
		&rpc.GetTokenAccountsConfig{
			ProgramId: &tokenProgramID,
		},
		&rpc.GetTokenAccountsOpts{
			Encoding: solana.EncodingBase64Zstd,
		},
	)
	if err != nil {
		return nil, err
	}

	mints := []string{}
	for _, rawAccount := range out.Value {
		var tokAcc token.Account
		data := rawAccount.Account.Data.GetBinary()
		dec := bin.NewBinDecoder(data)
		err := dec.Decode(&tokAcc)
		if err != nil {
			log.Printf("Wallet GetTokensByOwner >> %s", err.Error())
			continue
		}
		if tokAcc.Amount > 0 {
			mints = append(mints, tokAcc.Mint.ToPointer().String())
		}
	}

	return mints, nil
}
