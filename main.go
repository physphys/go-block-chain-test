package main

import (
	"block_chain/bchain"
	"fmt"
)

func main() {
	c := bchain.NewChain()

	t, err := bchain.NewTransaction("Ken", "Taro", 2)
	if err != nil {
		fmt.Println(err)

		return
	}
	c.AddTransaction(t)

	t, err = bchain.NewTransaction("Taro", "Ken", 10)
	if err != nil {
		fmt.Println(err)

		return
	}
	c.AddTransaction(t)

	if err := c.Mine("Ken"); err != nil {
		return
	}
	if err := c.Mine("Ken"); err != nil {
		return
	}

	if err := c.Validate(); err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf("%+v\n", c)
	fmt.Printf("Ken balance: %f\n", c.GetBalanceOfAddress("Ken"))
	fmt.Printf("Taro balance: %f\n", c.GetBalanceOfAddress("Taro"))
}
