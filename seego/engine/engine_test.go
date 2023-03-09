package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func ExampleEngine() {
	res := wrapper.NewEngineResource(*resource)
	defer res.Delete()

	opts := wrapper.NewEngineOptions(*options)
	defer opts.Delete()

	engine := wrapper.NewEngine(res, opts)
	defer engine.Delete()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		result, _ := engine.Predict(bytes)
		fmt.Println(result)
	}
}
