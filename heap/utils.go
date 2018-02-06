package heap

import (
	"reflect"
	"fmt"
	"testing"
	"math/rand"
	"sort"
)

func basicTestHeap(t *testing.T, h HeapIf) {
	arr := []int{1,3,2,2,4,5}
	sortedArr := make([]int, 0, 0)
	for _, v := range arr {
		h.Append(v)
	}
	for h.Len() > 0 {
		sortedArr = append(sortedArr, h.Pop().(int))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	if !reflect.DeepEqual(sortedArr, arr) {
		t.Log(fmt.Sprintf("expect:%v", arr) + fmt.Sprintf("but get:%v", sortedArr))
		t.Fail()
	}
}

func testHeap(t *testing.T, h HeapIf) {
	arrSize := rand.Intn(100) + 50
	arr := make([]int, arrSize, arrSize)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	sortedArr := make([]int, 0, 0)
	for _, v := range arr {
		h.Append(v)
	}
	for h.Len() > 0 {
		sortedArr = append(sortedArr, h.Pop().(int))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	if !reflect.DeepEqual(sortedArr, arr) {
		t.Log(fmt.Sprintf("expect:%v", arr) + fmt.Sprintf("but get:%v", sortedArr))
		t.Fail()
	}
}

func testHeapUnion(t *testing.T, h1 HeapIf, h2 HeapIf) {
	arrSize := rand.Intn(100) + 50
	arrSize1 := rand.Intn(arrSize)
	arr := make([]int, arrSize, arrSize)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	sortedArr := make([]int, 0, 0)
	for i, v := range arr {
		if i < arrSize1 {
			h1.Append(v)
		} else {
			h2.Append(v)
		}
	}
	h1 = h1.Union(h2).(HeapIf)
	for h1.Len() > 0 {
		sortedArr = append(sortedArr, h1.Pop().(int))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	if !reflect.DeepEqual(sortedArr, arr) {
		t.Log(fmt.Sprintf("expect:%v", arr) + fmt.Sprintf("but get:%v", sortedArr))
		t.Fail()
	}
}

func benchmarkHeap(b *testing.B, h HeapIf) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrSize := 10000
		arr := make([]int, arrSize, arrSize)
		for i := range arr {
			arr[i] = rand.Intn(arrSize)
		}
		b.StartTimer()

		sortedArr := make([]int, 0, 0)
		for _, v := range arr {
			h.Append(v)
		}
		for h.Len() > 0 {
			sortedArr = append(sortedArr, h.Pop().(int))
		}
	}
}
