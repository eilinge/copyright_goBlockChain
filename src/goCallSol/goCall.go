package main

import (
	"fmt"
	"log"

	// "net/rpc"
	"github.com/ethereum/go-ethereum/rpc"
)

func newAcc(pass string) {
	cli, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal("failed to connect to geth", err)
	}
	defer cli.Close()

	var account string
	err = cli.Call(&account, "personal_newAccount", pass) // 调用ether的personal.newAccount方法
	if err != nil {
		log.Fatal("failed to personal_newAccount", err)
	}
	fmt.Println("new account: ", account)
}

func main() {
	newAcc("eilinge")
}
