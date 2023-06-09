package accounts

import (
	accountbalances "hi-supergirl/blockchain-with-go-exercise/accounts/AccountBalances"
	accounttokenbalances "hi-supergirl/blockchain-with-go-exercise/accounts/AccountTokenBalances"
	addresscheck "hi-supergirl/blockchain-with-go-exercise/accounts/AddressCheck"
	generatingnewwallets "hi-supergirl/blockchain-with-go-exercise/accounts/GeneratingNewWallets"
	keystores "hi-supergirl/blockchain-with-go-exercise/accounts/Keystores"
)

func ReadBalanceInfo() {
	accountbalances.ReadBalanceInfo()
}

func ERC20Testcases() {
	accounttokenbalances.ERC20Testcases()
}

func CreateWallet() {
	generatingnewwallets.CreateWallet()
}

func CreateKeyStore() {
	keystores.CreateKeyStore()
}

func ImportKeyStore() {
	keystores.ImportKeyStore()
}

func CheckPublicAddress() {
	addresscheck.CheckPublicAddress()
}

func CheckAddressIsSmartContract() {
	addresscheck.CheckAddressIsSmartContract()
}
