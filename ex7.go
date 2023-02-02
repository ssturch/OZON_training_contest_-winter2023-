package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func PrintZeroz(uQ int) {
	for i := 0; i < uQ; i++ {
		fmt.Println("0")
	}
}

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var userQty int
	var pairQty int

	users := map[int]map[int]bool{}
	fmt.Fscan(in, &userQty)
	fmt.Fscan(in, &pairQty)

	if pairQty == 0 {
		PrintZeroz(userQty)
		return
	}

	for i := 0; i < pairQty; i++ {
		var tempUser int
		var tempFriend int
		fmt.Fscan(in, &tempUser)
		fmt.Fscan(in, &tempFriend)
		_, ok := users[tempUser]

		if !ok {
			users[tempUser] = map[int]bool{}
			users[tempUser][tempFriend] = false
		} else {
			users[tempUser][tempFriend] = false
		}

		_, ok = users[tempFriend]
		if !ok {
			users[tempFriend] = map[int]bool{}
			users[tempFriend][tempUser] = false
		} else {
			users[tempFriend][tempUser] = false
		}
	}

	tempMap := make(map[int]int)
	for i := 1; i < userQty+2; i++ {
		var tempArr []int
		tempMap = make(map[int]int)
		for k, _ := range users[i] {
			for m, _ := range users[k] {
				if m == i {
					continue
				}
				_, ok := users[i][m]
				if ok {
					continue
				}
				_, ok = users[i][k]
				if !ok {
					continue
				} else {
					tempMap[m]++
				}
			}
		}

		maxFriends := 0
		for _, val := range tempMap {
			if val > maxFriends {
				maxFriends = val
			}
		}

		for key, val := range tempMap {
			if val == maxFriends {
				tempArr = append(tempArr, key)
			}
		}
		sort.Ints(tempArr)
		if i == userQty+1 {
			break
		}
		//
		if len(tempArr) == 0 {
			fmt.Println("0")
			continue
		} else {
			fmt.Println(strings.Trim(fmt.Sprint(tempArr), "[]"))
		}
	}
}
