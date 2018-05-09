package main

import (
	"fmt"
	"log"
	"regexp"
	"runtime"
	"time"

	bbsort "./bublesort"
	gen "./generators"
	msort "./merge_sort"
	qsort "./qsort"
	selection_sort "./selection_sort"
)

func QuickSort_time(slice []int) {
	defer TimeTrack(time.Now())
	qsort.QuickSort(slice)
}
func SelectionSort_time(slice []int) {
	defer TimeTrack(time.Now())
	selection_sort.Sort(slice)
}
func MergeSort_time(slice []int) {
	defer TimeTrack(time.Now())
	msort.MergeSort(slice)
}
func BubbleSort_time(slice []int) {
	defer TimeTrack(time.Now())
	bbsort.Bubblesort(slice)
}

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)
	pc, _, _, _ := runtime.Caller(1)
	funcObj := runtime.FuncForPC(pc)
	// Regex to extract just the function name (and not the module path).
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")
	log.Println(fmt.Sprintf("%s took %s", name, elapsed))
}

func main() {
	//random slices
	log.Println(fmt.Sprintf("\n\n generate slice of size %d\n", 1000))
	rslice1 := gen.GenerateSlice(1000)
	QuickSort_time(rslice1)
	SelectionSort_time(rslice1)
	MergeSort_time(rslice1)
	BubbleSort_time(rslice1)
	log.Println(fmt.Sprintf("\n\n generate slice of size %d\n", 10000))
	rslice2 := gen.GenerateSlice(10000)
	QuickSort_time(rslice2)
	SelectionSort_time(rslice2)
	MergeSort_time(rslice2)
	BubbleSort_time(rslice2)
	log.Println(fmt.Sprintf("\n\n generate slice of size %d\n", 100000))
	rslice3 := gen.GenerateSlice(100000)
	QuickSort_time(rslice3)
	SelectionSort_time(rslice3)
	MergeSort_time(rslice3)
	BubbleSort_time(rslice3)

	log.Println(fmt.Sprintf("\n\n generate inverse slice of size %d\n", 1000))
	reslice1 := gen.GenerateInverseSlice(1000)
	QuickSort_time(reslice1)
	SelectionSort_time(reslice1)
	MergeSort_time(reslice1)
	BubbleSort_time(reslice1)
	log.Println(fmt.Sprintf("\n\n generate inverse slice of size %d\n", 10000))
	reslice2 := gen.GenerateInverseSlice(10000)
	QuickSort_time(reslice2)
	SelectionSort_time(reslice2)
	MergeSort_time(reslice2)
	BubbleSort_time(reslice2)
	log.Println(fmt.Sprintf("\n\n generate inverse slice of size %d\n", 100000))
	reslice3 := gen.GenerateInverseSlice(100000)
	QuickSort_time(reslice3)
	SelectionSort_time(reslice3)
	MergeSort_time(reslice3)
	BubbleSort_time(reslice3)
}
