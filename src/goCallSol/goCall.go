package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func newAcc(pass string) {
	cli, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal("failed to connect to geth", err)
	}
	defer cli.Close()

	var account string
	err = cli.Call(&account, "personal_newAccount", pass)
	if err != nil {
		log.Fatal("failed to personal_newAccount", err)
	}
	fmt.Println("new account: ", account)
}

func main() {
	newAcc("eilinge")
}
