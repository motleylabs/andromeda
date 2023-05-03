package state

import (
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/stat/market"
	"andromeda/pkg/service/web3"

	"log"
	"sync"
	"time"
)

type State struct {
	Version      int64
	SOLPrice     float64
	TPS          uint64
	MarketVolume float64
	GlobalStats  types.StatRes
}

func (State *State) Initialize() error {
	State.Version = time.Now().Unix()

	wg := sync.WaitGroup{}
	wg.Add(3)

	// Get overall stat
	var globalStats *types.StatRes
	var statResError error

	go func() {
		defer wg.Done()
		globalStats, statResError = market.GetOverall()
	}()

	// Get SOL Price
	var solPrice float64
	var solPriceError error

	go func() {
		defer wg.Done()
		solPrice, solPriceError = web3.GetSOLPrice()
	}()

	// Get SOL TPS
	var tps uint64
	var solTpsError error

	go func() {
		defer wg.Done()
		tps, solTpsError = web3.GetSOLTPS()
	}()

	wg.Wait()

	// error handler
	if statResError != nil {
		log.Printf("State Intialize >> Market GetOverall; %s", statResError.Error())
		return statResError
	}
	if solPriceError != nil {
		log.Printf("State Initialize >> Web3 GetSOLPrice; %s", solPriceError.Error())
		return solPriceError
	}
	if solTpsError != nil {
		log.Printf("State Initialize >> Web3 GetSOLTPS; %s", solTpsError.Error())
		return solTpsError
	}

	State.GlobalStats = *globalStats
	State.SOLPrice = solPrice
	State.TPS = tps
	State.MarketVolume = 0

	return nil
}

var CurrentState State

func Update() {
	state := State{}
	err := state.Initialize()
	if err == nil {
		log.Printf("State#Update: from version %v to %v\n", CurrentState.Version, state.Version)
		log.Printf("State#Update: new state %+v\n", state)
		CurrentState = state
	}
}

func Runloop() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		Update()
	}
}

const MaxRetries = 10

func Ensure() bool {
	for retries := 0; CurrentState.Version == 0 && retries < MaxRetries; retries++ {
		Update()
	}

	return CurrentState.Version != 0
}

func GetReport() types.ReportRes {
	return types.ReportRes{
		Volume:   CurrentState.MarketVolume,
		TPS:      CurrentState.TPS,
		SOLPrice: CurrentState.SOLPrice,
	}
}

func GetSOLPrice() float64 {
	return CurrentState.SOLPrice
}
