package main

import (
	"block_chain/bchain"
	"fmt"
)

func main() {
	bc := bchain.NewBlockChain()
	bc.AddTransaction(
		bchain.NewTransaction(
			"Ken",
			"Taro",
			2,
		),
	)
	bc.AddTransaction(
		bchain.NewTransaction(
			"Taro",
			"Ken",
			10,
		),
	)

	bc.Mine("Ken")
	bc.Mine("Ken")

	isValid, err := bc.IsValid()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Block Chain validation %t\n", isValid)

	fmt.Printf("%+v\n", bc)
	fmt.Printf("Ken balance: %f\n", bc.GetBalanceOfAddress("Ken"))
	fmt.Printf("Taro balance: %f\n", bc.GetBalanceOfAddress("Taro"))
}
