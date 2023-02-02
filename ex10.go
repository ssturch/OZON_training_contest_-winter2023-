package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	rymMap := make(map[string][]string)
	var anyWord [2]string
	var libSize int
	fmt.Fscan(in, &libSize)
	for i := 0; i < libSize; i++ {
		var word string
		fmt.Fscan(in, &word)
		if i < 2 {
			anyWord[i] = word
		}
		origWord := word
		for j := len(word) - 1; j >= 0; j-- {
			_, ok := rymMap[word[j:]]
			if !ok {
				var libArr []string
				libArr = append(libArr, origWord)
				rymMap[word[j:]] = libArr
			} else {
				tempLib := rymMap[word[j:]]
				tempLib = append(tempLib, origWord)
				rymMap[word[j:]] = tempLib
			}
		}
	}
	var checkSize int
	fmt.Fscan(in, &checkSize)
	var out string
	for k := 0; k < checkSize; k++ {
		var word string
		fmt.Fscan(in, &word)
		out = anyWord[1]
		for l := len(word) - 1; l >= 0; l-- {
			val, ok := rymMap[word[l:]]
			if ok {
				for m := 0; m < len(val); m++ {
					if word != val[m] {
						out = val[m]
						break
					}
				}
			}
		}


		fmt.Println(out)
	}
}
