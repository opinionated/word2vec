/*
word2vec-client is a tool which uses word2vec.Client to look up similarities (using an external server).
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sajari/word2vec"
)

var addr string
var addListA, subListA string
var addListB, subListB string

func init() {
	flag.StringVar(&addr, "addr", "localhost:1234", "server address")
	flag.StringVar(&addListA, "addA", "", "comma separated list of model words to add to the target vector A")
	flag.StringVar(&subListA, "subA", "", "comma separated list of model words to subtract from the target vector A")
	flag.StringVar(&addListB, "addB", "", "comma separated list of model words to add to the target vector B")
	flag.StringVar(&subListB, "subB", "", "comma separated list of model words to subtract from the target vector B")
}

func makeExpr(addList, subList string) (word2vec.Expr, error) {
	if addList == "" && subList == "" {
		return word2vec.Expr{}, fmt.Errorf("must specify 'add' and/or 'sub' component for each target vector; see -h for more details")
	}

	result := word2vec.Expr{}
	if addList != "" {
		result.Add = strings.Split(addList, ",")
	}
	if subList != "" {
		result.Sub = strings.Split(subList, ",")
	}

	return result, nil
}

func main() {
	flag.Parse()

	if addr == "" {
		fmt.Println("must specify -addr; see -h for more details")
		os.Exit(1)
	}

	exprA, err := makeExpr(addListA, subListA)
	if err != nil {
		fmt.Printf("error creating target vector for 'A': %v\n", err)
		os.Exit(1)
	}

	exprB, err := makeExpr(addListB, subListB)
	if err != nil {
		fmt.Printf("error creating target vector for 'B': %v\n", err)
		os.Exit(1)
	}

	c := word2vec.Client{Addr: addr}

	start := time.Now()
	v, err := c.Sim(exprA, exprB)
	totalTime := time.Since(start)
	if err != nil {
		fmt.Printf("error looking up similarity: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Similarity: %v (took: %v)\n", v, totalTime)
}