package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var qty int
		var skill int

		fmt.Fscan(in, &qty)

		var skillArr []int
		for i := 0; i < qty; i++ {
			fmt.Fscan(in, &skill)
			skillArr = append(skillArr, skill)
		}
		
		
		var pair int
		var minDiff, minDiffTemp float64
		minDiff = math.MaxFloat64
		for i := 0; i < qty-1; i++ {
			if skillArr[i] == 0 {
				continue
			}
			for j := i+1; j < qty; j++ {
				if skillArr[j] == 0 {
					continue
				} else {
					minDiffTemp = math.Abs(float64(skillArr[i] - skillArr[j]))
					if minDiff > minDiffTemp {
						minDiff = minDiffTemp
						pair = j
					}
				}
			}
			fmt.Printf("%v %v\n", i+1, pair+1)
			skillArr[i] = 0
			skillArr[pair] = 0
			minDiff = math.MaxFloat64
		}
		fmt.Println("")
	}
}
