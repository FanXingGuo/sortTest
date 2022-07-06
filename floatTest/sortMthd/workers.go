package sortMthd

// 常见排序的方法 线程调用

import (
	"github.com/psilva261/timsort/v2"
	"sort"
	"sync"
)

func workerQuickSortU(indexs[2]int, arr []float64,group *sync.WaitGroup) {
	defer group.Done()
	QuickSortU(arr[indexs[0]:indexs[1]])
}
func workerInsertSort(indexs[2]int, arr []float64,group *sync.WaitGroup) {
	defer group.Done()
	InsertSort(arr[indexs[0]:indexs[1]])
}
func workerSelectSort(indexs[2]int, arr []float64,group *sync.WaitGroup) {
	defer group.Done()
	SelectSort(arr[indexs[0]:indexs[1]])
}
func workerBubbleSort(indexs[2]int, arr []float64,group *sync.WaitGroup) {
	defer group.Done()
	BubbleSort(arr[indexs[0]:indexs[1]])
}
func workerShellSort(indexs[2]int, arr []float64,group *sync.WaitGroup) {
	defer group.Done()
	ShellSort(arr[indexs[0]:indexs[1]])
}
func workerMergeSort(indexs[2]int, arr []float64,group *sync.WaitGroup) {
	defer group.Done()
	MergeSort(arr[indexs[0]:indexs[1]])
}
func workerHeapSort(indexs[2]int, arr []float64,group *sync.WaitGroup) {
	defer group.Done()
	MergeSort(arr[indexs[0]:indexs[1]])
}
func workerCountSort(indexs[2]int, arr []float64,group *sync.WaitGroup) {
	defer group.Done()
	MergeSort(arr[indexs[0]:indexs[1]])
}
func workerBucketSort(indexs[2]int, arr []float64,group *sync.WaitGroup) {
	defer group.Done()
	bktLen:=len(arr)/1000
	BucketSort(arr[indexs[0]:indexs[1]],bktLen)
}
func workerRadixSort(indexs[2]int, arr []float64,group *sync.WaitGroup) {
	defer group.Done()
	RadixSort(arr[indexs[0]:indexs[1]])
}
func workerTimSort(indexs[2]int, arr []float64,group *sync.WaitGroup) {
	defer group.Done()
	//timsort.TimSort(sort.IntSlice(arr[indexs[0]:indexs[1]]))
	timsort.TimSort(sort.Float64Slice(arr[indexs[0]:indexs[1]]))
}

