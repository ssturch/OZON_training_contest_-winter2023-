package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var ticketsCount int
		fmt.Fscan(in, &ticketsCount)

		var ticket int
		var ticketsArr []int
		ticketsMap := make(map[int]bool)
        
		isBreak := false
		for i := 0; i < ticketsCount; i++ {
			fmt.Fscan(in, &ticket)
			ticketsArr = append(ticketsArr, ticket)
		}
		ticketsArr = append(ticketsArr, -1)
		for i:=0; i < ticketsCount; i++ {
		    if ticketsArr[i+1] != ticketsArr[i]{
		        _, ok := ticketsMap[ticketsArr[i]]
		        if !ok{
		            ticketsMap[ticketsArr[i]] = true
		        }else{
		            isBreak = true
		            break
		        }
		    }
		}
			

		if isBreak {
			fmt.Println("NO")
		}else{
			fmt.Println("YES")
		}
	}
}
