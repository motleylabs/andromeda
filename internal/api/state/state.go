package state

import (
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/stat/market"
	"andromeda/pkg/service/web3"

	"log"
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
	globalStats, err := market.GetOverall()
	if err != nil {
		return err
	}

	State.GlobalStats = *globalStats

	solPrice, err := web3.GetSOLPrice()
	if err != nil {
		return err
	}

	State.SOLPrice = solPrice

	tps, err := web3.GetSOLTPS()
	if err != nil {
		return err
	}

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
	for {
		Update()
		time.Sleep(10 * time.Second)
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
