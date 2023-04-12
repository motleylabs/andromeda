package web3

import (
	"context"
	"os"

	"github.com/gagliardetto/solana-go/rpc"
)

type JSONRPC struct {
}

func GetSOLTPS() (uint64, error) {
	solTPC := uint64(0)
	client := rpc.New(os.Getenv("RPC_ENDPOINT"))
	limit := uint(1)

	out, err := client.GetRecentPerformanceSamples(context.TODO(), &limit)
	if err != nil {
		return solTPC, err
	}

	solTPC = out[0].NumTransactions / uint64(out[0].SamplePeriodSecs)
	return solTPC, nil
}

func GetVolume(address string) (float64, error) {
	return 0, nil
}
