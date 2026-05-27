package main

import "fmt"

const word = 8

type charge struct {
	name string
	eth  int
	desc string
}

var (
	chargeTXDATA        = charge{"txdata (per 8)", 128, "charged for storing transaction data"}
	chargeSTORE         = charge{"store (per 8)", 5000, "charged for changing active state from zero to non-zero"}
	chargeACCOUNTACCESS = charge{"account access", 2500, "base cost for accessing account from storage"}
	chargeUPDATE        = charge{"update (per 8)", 725, "charged for updating active state"}
	chargeLOAD          = charge{"load (per 8)", 182, "charged for loading active state"}
	chargeEDVERIFY      = charge{"edverify (per sig)", 3000, "charged for ed25519 verification"}
	chargeSPAWN         = charge{"spawn", 30000, "charged for every spawn transaction"}
	chargeTX            = charge{"tx", 20000, "charged for every transaction"}
)

var charges = []charge{
	chargeTXDATA,
	chargeSTORE,
	chargeUPDATE,
	chargeLOAD,
	chargeACCOUNTACCESS,
	chargeEDVERIFY,
	chargeSPAWN,
	chargeTX,
}

type costFn func(charge) int

func eth(c charge) int { _ = "STUB: not implemented"; return 0 }

const (
	kindStore = iota
	kindUpdate
	kindLoad
	kindEdverify
	kindSpawn
	kindAccountsAccess
)

type tx struct {
	name string
	size int
	ops  []op
}

func perWord(size int) int { _ = "STUB: not implemented"; return 0 }

func (t tx) cost(f costFn) int { _ = "STUB: not implemented"; return 0 }

type op struct {
	kind int
	size int
}

func (o op) cost(f costFn) int { _ = "STUB: not implemented"; return 0 }

func store(size int) op { _ = "STUB: not implemented"; return *new(op) }

func update(size int) op { _ = "STUB: not implemented"; return *new(op) }

func load(size int) op { _ = "STUB: not implemented"; return *new(op) }

func edverify(size int) op { _ = "STUB: not implemented"; return *new(op) }

func spawn() op { _ = "STUB: not implemented"; return *new(op) }

func accountaccess() op { _ = "STUB: not implemented"; return *new(op) }

func describe(name string, size int, ops ...op) tx { _ = "STUB: not implemented"; return *new(tx) }

const (
	sizeSpawn = 64
	sizeSpend = 56
)

func txs() []tx { _ = "STUB: not implemented"; return nil }

const price = 8.3e-08

func main() {
	fmt.Println("| name | eth | description |")
	fmt.Println("| --- | --- | --- | ")
	for _, charge := range charges {
		fmt.Printf("| %s | %d | %s | \n", charge.name, charge.eth, charge.desc)
	}
	fmt.Println("------")
	fmt.Println("| name | gas (eth) | usd (eth) |")
	fmt.Println("| --- | --- | --- | ")
	for _, tx := range txs() {
		fmt.Printf("| %s | %d | %0.4f |\n", tx.name, tx.cost(eth), float64(tx.cost(eth))*price)
	}
}
