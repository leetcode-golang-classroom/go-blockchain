package main

import (
	"os"

	"github.com/leetcode-golang-classroom/golang-blockchain/cli"
	// "github.com/leetcode-golang-classroom/golang-blockchain/wallet"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
	// w := wallet.MakeWallet()
	// w.Address()
}
