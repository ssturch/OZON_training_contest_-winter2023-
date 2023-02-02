package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var lines, columns, val, clickCount int
		fmt.Fscan(in, &lines)
		fmt.Fscan(in, &columns)

		table := make([][]int, lines)
		for i := range table {
			table[i] = make([]int, columns)
		}

		for i := 0; i < lines; i++ {
			for j := 0; j < columns; j++ {
				fmt.Fscan(in, &val)
				table[i][j] = val
			}
		}
       
        var clicks []int
        var click int
        fmt.Fscan(in,&clickCount)
        for i:=0; i < clickCount; i++ {
            fmt.Fscan(in,&click)
            clicks = append(clicks, click-1)
        }
        
        for _,cl := range clicks {
            sort.SliceStable(table, func(i, j int) bool {
                return table[i][cl] < table[j][cl]
            })
        }
        
        
        
        for i := 0; i < lines; i++ {
            tempStrSlice := make([]string, len(table[i]))
            for i, v := range table[i]{
                tempStrSlice[i] = strconv.Itoa(v)
            }
			fmt.Println(strings.Join(tempStrSlice," "))
		}
        fmt.Print("\n")
	}
}

