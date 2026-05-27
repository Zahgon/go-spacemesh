package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/genvm/cmd/gen"
)

var (
	hrp    = flag.String("hrp", "sm", "network human readable prefix")
	out    = flag.String("out", "json", "json|go")
	config = flag.String("c", "", "see example json")
)

func must(err error) { _ = "STUB: not implemented"; return }

func parseJson(f io.Reader) []gen.Input { _ = "STUB: not implemented"; return nil }

func parseCsv(f io.Reader) []gen.Input { _ = "STUB: not implemented"; return nil }

func parseInputs() []gen.Input { _ = "STUB: not implemented"; return nil }

func main() {
	flag.Parse()
	types.SetNetworkHRP(*hrp)
	if len(*config) == 0 {
		fmt.Println("please specify config with -c=<path to file>")
		os.Exit(1)
	}

	inputs := parseInputs()
	var outputs []gen.Output
	for _, input := range inputs {
		outputs = append(outputs, gen.Generate(input))
	}
	switch *out {
	case "json":
		for _, output := range outputs {
			enc := json.NewEncoder(os.Stdout)
			if err := enc.Encode(output); err != nil {
				must(err)
			}
		}
	case "go":
		fmt.Println("func MainnetAccounts() map[string]uint64 {")
		fmt.Println("    return map[string]uint64{")
		for _, output := range outputs {
			fmt.Printf("        \"%s\": %d,\n", output.Address, output.Balance)
		}
		fmt.Println("    }")
		fmt.Println("}")
	}
}
