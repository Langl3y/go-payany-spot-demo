package main

import (
	"fmt"
	"spot_demo/business"
)

func stressTest() {
	result, err := business.PutSpotLimit("hieuhovan954@gmail.com", "eriri123")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func main() {
	for {
		go stressTest()
	}
}
