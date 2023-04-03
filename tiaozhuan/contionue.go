package tiaozhuan

import "fmt"

func ContinuDemo() {
forloop:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}

}
