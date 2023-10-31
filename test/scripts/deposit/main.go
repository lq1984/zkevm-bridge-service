package main

import (
	"context"
	"math/big"

	"github.com/0xPolygonHermez/zkevm-bridge-service/utils"
	"github.com/0xPolygonHermez/zkevm-node/log"
	"github.com/ethereum/go-ethereum/common"
)

const (
	l1BridgeAddr = "0xff0EE8ea08cEf5cb4322777F5CC3E8A584B8A4A0"
	l2BridgeAddr = "0xff0EE8ea08cEf5cb4322777F5CC3E8A584B8A4A0"

	l1AccHexAddress    = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	l1AccHexPrivateKey = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	l1NetworkURL       = "http://localhost:8545"
	l2NetworkURL       = "http://localhost:8123"

	funds                = 90000000000000000 // nolint
	L1destNetwork uint32 = 1
	L2destNetwork uint32 = 0
)

var tokenAddr = common.Address{}

func main() {
	depositL1ToL2()
	//depositL2ToL1()
}

func depositL1ToL2() {
	ctx := context.Background()
	client, err := utils.NewClient(ctx, l1NetworkURL, common.HexToAddress(l1BridgeAddr))
	if err != nil {
		log.Fatal("Error: ", err)
	}
	auth, err := client.GetSigner(ctx, l1AccHexPrivateKey)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	amount := big.NewInt(funds)
	emptyAddr := common.Address{}
	if tokenAddr == emptyAddr {
		auth.Value = amount
	}
	destAddr := common.HexToAddress(l1AccHexAddress)
	log.Info("Sending bridge tx...")
	err = client.SendBridgeAsset(ctx, tokenAddr, amount, L1destNetwork, &destAddr, []byte{}, auth)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	log.Info("Success!")
}

func depositL2ToL1() {
	ctx := context.Background()
	client, err := utils.NewClient(ctx, l2NetworkURL, common.HexToAddress(l2BridgeAddr))
	if err != nil {
		log.Fatal("Error: ", err)
	}
	auth, err := client.GetSigner(ctx, l1AccHexPrivateKey)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	amount := big.NewInt(funds)
	emptyAddr := common.Address{}
	if tokenAddr == emptyAddr {
		auth.Value = amount
	}
	destAddr := common.HexToAddress(l1AccHexAddress)
	log.Info("Sending bridge tx...")
	err = client.SendBridgeAsset(ctx, tokenAddr, amount, L2destNetwork, &destAddr, []byte{}, auth)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	log.Info("Success!")
}
