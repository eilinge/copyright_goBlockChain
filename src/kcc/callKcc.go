package main

import (
	"fmt"
	"kcctoken"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func _transfer(_addr string, _num int64) {
	cli, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal("failed to ethclient.Dial")
	}
	instance, err := kcctoken.NewToken(common.HexToAddress("0xf3a68ae8ca860a1803aad6bf6aa6b1e6ee80f6b1"), cli)
	if err != nil {
		log.Fatal("failed to kcctoken.NewToken")
	}
	file, err := os.Open("../../data/keystore/UTC--2019-06-19T22-59-49.902830200Z--6cfef69b6f13ade124d5f670ce7e00cdb54e8d3e")
	if err != nil {
		log.Fatal("failed to os.Open")
	}
	auth, err := bind.NewTransactor(file, "eilinge")
	if err != nil {
		log.Fatal("failed to bind.NewTransactor")
	}
	_, err = instance.Transfer(auth, common.HexToAddress(_addr), big.NewInt(_num))
	if err != nil {
		log.Fatal("failed to instance.Transfer")
	}
	fmt.Println("Transfer token success")
}

func main() {
	_transfer("0xc3c480c212bc5485062655f3848691974d784b69", 10000)
}
