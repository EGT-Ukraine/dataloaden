package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/EGT-Ukraine/dataloaden/pkg/generator"
)

func main() {
	keyType := flag.String("keys", "int", "what type should the keys be")
	slice := flag.Bool("slice", false, "this dataloader will return slices")
	pointer := flag.Bool("pointer", false, "this dataloader will return pointer")
	name := flag.String("name", "", "name of dataloader")

	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	wd, err := os.Getwd()

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	if err := generator.Generate(*name, flag.Arg(0), *keyType, *slice, *pointer, wd); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
}
