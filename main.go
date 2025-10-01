package main

import (
	"fmt"
	"runtime"
	"spot_demo/business"
)

func stressTest() {
	result, err := business.PutSpotLimit("hieuhovan954@gmail.com", "eriri123")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func runGotoutines() {
	defer fmt.Printf("goroutines count: %d\n", runtime.NumGoroutine())
	go stressTest()
}

func main() {
	for i := range 10 {
		runGotoutines()
		i += 1
	}
}
