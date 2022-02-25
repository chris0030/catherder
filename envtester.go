package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func difference(a, b []string) []string {
    mb := make(map[string]struct{}, len(b))
    for _, x := range b {
        mb[x] = struct{}{}
    }
    var diff []string
    for _, x := range a {
        if _, found := mb[x]; !found {
            diff = append(diff, x)
        }
    }
    return diff
}

func Keys(m map[string]string) (keys []string) {
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}

func main() {
	expectedEnv, err := godotenv.Read(os.Args[1])
	testedEnv, err := godotenv.Read(os.Args[2])

	if err != nil {
		log.Fatal(err)
	}
	missing := difference(Keys(expectedEnv), Keys(testedEnv))
	additional := difference(Keys(testedEnv), Keys(expectedEnv))
	var exitCode int = 0
	if len(missing) > 0 {
		fmt.Printf("The following keys are missing: %s\n", missing)
		exitCode = 1
	}
	if len(additional) > 0 {
		fmt.Printf("The following keys are not expected: %s\n", additional)
		exitCode = 1
	}
	os.Exit(exitCode)
}