// cmd/blockchainconsensusoptimizerplatformultra/main.go
package main

import (
"flag"
"log"
"os"

"blockchainconsensusoptimizerplatformultra/internal/blockchainconsensusoptimizerplatformultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainconsensusoptimizerplatformultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
