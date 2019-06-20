# go语言编程

## 问题描述

一、 问题描述:
(共4道题，都是必做题)

1) 使用go（或者其他语言）调用ERC20智能合约时，调用查看余额和转账有什么区别?

	转账时: `instance.Transfer(auth, common.HexToAddress(_addr), big.NewInt(_num))`, auth是需要传入的私钥+密码`bind.NewTransactor(file, "eilinge")`, 进行签名交易.
	而查看余额不需要传入私钥+密码
	
2) 编码实现：启动三个协程，第一个协程将0，1，2，3……依次传递给第二个协程，第二个协程将得到的数字平方后传递给第三个协程,第三个协程负责打印得到的数字。

```package main

import (
	"fmt"
	"time"
)

// 编码实现：启动三个协程,第一个协程将0,1,2,3……依次传递给第二个协程,
// 第二个协程将得到的数字平方后传递给第三个协程,第三个协程负责打印得到的数字

// 三个协程(相互通信)
// 一个channal(全局, 不需要作为参数传入到函数)
// chan<- uint8 必须在协程中执行

var chanTrans chan uint8

func transfer() {
	var n1 uint8
	for {
		n := <-chanTrans
		n1 = n * n
		go receive(n1)
	}
}

func receive(n uint8) {
	fmt.Println(n)
}

func main() {
	chanTrans = make(chan uint8)
	defer close(chanTrans)
	var i uint8
	go transfer()
	go func() {
		for i = 0; i < 10; i++ {
			chanTrans <- i
		}
	}()
	time.Sleep(time.Millisecond * 10)
}
```

3) 简述如果使用go(其他语言也可) 调用ERC20智能合约转账函数都有哪些步骤?

	1. 编写实现erc20的智能合约(kccToken.sol)
	2. 使用`ethereum/go-ethereum/cmd/abigen` 编译的abigen.exe 可执行程序
	3. 获取kccToken的abi
	4. 使用`abigen -type token -pkg kcctoken -abi kcc.abi -out kcc.go`, 生成kcc.go文件
	5. 使用geth开启私链`geth --datadir ./data --networkid 15 --port 30303 --rpc --rpcaddr 0.0.0.0 --rpcport 8545 --rpcvhosts "*" --rpcapi 'db,net,eth,web3,personal' --rpccorsdomain "*" --ws --wsaddr "localhost" --wsport "8546" --wsorigins "*" --nat "any" --nodiscover --dev --dev.period 1 console 2> 1.log`
	6. 解锁账户`personal.unlockAccount(eth.coinbase)`, 并执行挖矿`miner.start()`
	7. 使用remix连接私网`http://ocalhost:8545`, 进行部署合约
	8. 编写go语言程序, 连接私网`ethclient.Dial("http://localhost:8545")`,获取合约实例`kcctoken.NewToken(common.HexToAddress("0xf3a68ae8ca860a1803aad6bf6aa6b1e6ee80f6b1"), cli)`
	9. 获取私钥+密钥`bind.NewTransactor(file, "eilinge")`,对转账进行签名`instance.Transfer(auth, common.HexToAddress(_addr), big.NewInt(_num))`

4) 对第3问，用代码加以实现，不限于go语言。

解题提示
二、 解题提示

1) 从gas消耗角度考虑;
2) 通道+协程的方式;
3) 列出操作步骤，注意从sol代码开始的步骤;
4) 实现后可用图片加文字进行介绍。
