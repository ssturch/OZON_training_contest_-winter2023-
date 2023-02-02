package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type StructHeap []*IdleProcStruct
type IdleProcStruct struct {
	idle int
	proc int
}

func (h StructHeap) Len() int           { return len(h) }
func (h StructHeap) Less(i, j int) bool { return h[i].idle < h[j].idle }
func (h StructHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *StructHeap) Push(x any) {
	//item := x.(*StructHeap)
	item := x.(*IdleProcStruct)
	*h = append(*h, item)
}

func (h *StructHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func main() {

	in := bufio.NewReader(os.Stdin)

	var procQty, taskQty, procComsumpt, timeStart, timeSleep, res int
	fmt.Fscan(in, &procQty)
	fmt.Fscan(in, &taskQty)

	timeNdurMap := make(map[int]int)
	timeHeap := &IntHeap{}
	effectHeap := &IntHeap{}
	idleNprocHeap := &StructHeap{}
	heap.Init(timeHeap)
	heap.Init(effectHeap)
	heap.Init(idleNprocHeap)

	for i := 0; i < procQty; i++ {
		fmt.Fscan(in, &procComsumpt)
		heap.Push(effectHeap, procComsumpt)
	}
	for i := 0; i < taskQty; i++ {
		fmt.Fscan(in, &timeStart)
		fmt.Fscan(in, &timeSleep)
		heap.Push(timeHeap, timeStart)
		timeNdurMap[timeStart] = timeSleep
	}

	for timeHeap.Len() > 0 {
		var tempProc interface{}
		tempTime := heap.Pop(timeHeap)
		if len(*idleNprocHeap) != 0 {
			for {
				tempIdle := heap.Pop(idleNprocHeap)
				if tempIdle.(*IdleProcStruct).idle <= tempTime.(int) {
					heap.Push(effectHeap, tempIdle.(*IdleProcStruct).proc)
				} else {
					heap.Push(idleNprocHeap, tempIdle)
				}
				if idleNprocHeap.Len() == 0 {
					break
				} else if tempIdle.(*IdleProcStruct).idle > tempTime.(int) {
					break
				}
			}
		}

		if len(*effectHeap) != 0 {
			tempProc = heap.Pop(effectHeap)
			itemIdleProc := &IdleProcStruct{
				idle: tempTime.(int) + timeNdurMap[tempTime.(int)],
				proc: tempProc.(int),
			}
			heap.Push(idleNprocHeap, itemIdleProc)
			res += tempProc.(int) * timeNdurMap[tempTime.(int)]
		}
	}
	fmt.Println(res)
}
