package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)
	layout := "15:04:05"
	zeroDate, _ := time.Parse(layout, "00:00:00")

	type t []struct {
		Left  uint32
		Right uint32
	}

	for i := 0; i < testCount; i++ {

		var timeArr t
		var timeGapsCount int
		fmt.Fscan(in, &timeGapsCount)
		timeMap := make(map[int][]string)
		var isBreak bool

		for i := 0; i < timeGapsCount; i++ {
			var timeGap string
			fmt.Fscan(in, &timeGap)
			timeMap[i] = strings.Split(timeGap, "-")
			timeLeft := timeMap[i][0]
			timeRight := timeMap[i][1]
			var timeParsedLeft time.Time
			var timeParsedRight time.Time
			var err error
			//проверка правильности написания времени
			timeParsedLeft, err = time.Parse(layout, timeLeft)
			if err != nil {
				isBreak = true
				continue
			}

			if !isBreak {
				timeParsedRight, err = time.Parse(layout, timeRight)
				if err != nil {
					isBreak = true
					continue
				}
			}
			//проверка условия "левая граница не позже правой"
			if !isBreak {
				if timeParsedLeft.After(timeParsedRight) {
					isBreak = true
					continue
				}
			}
			timeLeftDur := uint32(timeParsedLeft.Sub(zeroDate).Seconds())
			timeRightDur := uint32(timeParsedRight.Sub(zeroDate).Seconds())

			timeArr = append(timeArr, struct {
				Left  uint32
				Right uint32
			}{timeLeftDur, timeRightDur})
		}

		if isBreak == false {
			sort.SliceStable(timeArr, func(i, j int) bool {
				return timeArr[i].Left < timeArr[j].Left
			})

			for i := 0; i < len(timeArr)-1; i++ {
				for j := i + 1; j < len(timeArr); j++ {
					if (timeArr[i].Left <= timeArr[j].Right) && (timeArr[i].Right >= timeArr[j].Left) {
						isBreak = true
						break
					}
				}
				if isBreak {
					break
				}
			}
		}
		if isBreak {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
