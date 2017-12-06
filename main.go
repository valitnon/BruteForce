package main

import (
	"flag"
	"fmt"
	"github.com/ngirot/BruteForce/bruteforce"
	"os"
)

func main() {
	var bench = flag.Bool("benchmark", false, "Launch a benchmark")
	var value = flag.String("value", "", "The value to be tested")
	var alphabet = flag.String("alphabet", "", "The file containing all characters")
	var hashType = flag.String("type", "sha256", "The hash type")
	flag.Parse()

	if *bench {
		fmt.Printf("One CPU hasher : %d kh/s\n", bruteforce.BenchHasher()/1000)
		fmt.Printf("One CPU word generator : %d kw/s\n", bruteforce.BenchBruter()/1000)
		os.Exit(0)
	}

	if *value == "" {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("Start brute-forcing...\n")

	var chrono = bruteforce.NewChrono()
	chrono.Start()
	if result, error := bruteforce.Launch(*value, *alphabet, *hashType); error == nil {
		chrono.End()

		if result != "" {
			fmt.Printf("Found : %s\n", result)
			fmt.Printf("In %f s\n", chrono.DurationInSeconds())
		} else {
			fmt.Printf("Not found\n")
		}
	} else {
		fmt.Printf("Hasher %s invalid: %q\n", *hashType, error)
	}
}
