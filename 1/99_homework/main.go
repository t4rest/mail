package main

import (
	"strconv"
	"sort"
)

func ReturnInt() int {
	return 1
}

func ReturnFloat() float32 {
	return 1.1
}

func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

func IntSliceToString(sl []int) string {
	var str string

	for i := range sl {
		str = str + strconv.Itoa(sl[i])
	}

	return str
}

func MergeSlices(slf []float32, sli []int32) []int {
	var sl []int = make([]int, 0, len(slf) + len(sli))

	for i := range slf {
		sl = append(sl, int(slf[i]))
	}

	for i := range sli {
		sl = append(sl, int(sli[i]))
	}

	return sl
}

func GetMapValuesSortedByKey(mapi map[int]string) []string {
	var sls []string = make([]string, 0, len(mapi))
	var sli []int = make([]int, 0, len(mapi))

	for i := range mapi{
		sli = append(sli, i)
	}

	sort.Ints(sli)

	for _, i := range sli {
		sls = append(sls, mapi[i])
	}

	return sls
}
