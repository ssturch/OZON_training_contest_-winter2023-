package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var qty int
		var tempPrice int
		
		fmt.Fscan(in, &qty)
		
		var priceArrInt []int
		for i:=0;  i<qty; i++{
		    fmt.Fscan(in, &tempPrice)
		    priceArrInt = append(priceArrInt, tempPrice)
		}
		

		sort.Ints(priceArrInt)

		counter := 1
		res := priceArrInt[0]

		
		for i := 1; i < len(priceArrInt); i++ {
            if priceArrInt[i-1] == priceArrInt[i] {
                counter +=1
            } else {
                counter = 1
            }
            
            if counter % 3 == 0 {
                priceArrInt[i] = 0
            }
            res += priceArrInt[i]
		}
		fmt.Println(res)
	}
}
