package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/crunchyroll/rebnf"
)

func init() {
	seed := time.Now().UTC().UnixNano()
	fmt.Fprintf(os.Stderr, "Using seed: %d\n", seed)
	rand.Seed(seed)
}

func main() {
	args := os.Args
	if len(args) != 2 && len(args) != 3 {
		log.Fatalln("Usage: " + os.Args[0] + " NUM_PRODUCTIONS [PRODUCTION=statement]")
	}
	numProductions, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalln(err)
	}

	production := "statement"
	if len(args) == 3 {
		production = args[2]
	}

	grammar, err := rebnf.Parse("influxql.ebnf", nil)
	if err != nil {
		log.Fatalln(err)
	}

	if _, ok := grammar[production]; !ok {
		log.Fatalln(production + " is not a known production!")
	}

	maxRepetitions := 30
	maxRecursionDepth := 15
	padding := "???"
	isDebug := false
	ctx := rebnf.NewCtx(maxRepetitions, maxRecursionDepth, padding, isDebug)

	for i := 0; i < numProductions; i++ {
		ctx.Random(os.Stdout, grammar, production)
		fmt.Fprintln(os.Stdout)
	}
}
