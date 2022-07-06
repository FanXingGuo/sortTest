package sortMthd

// 排序的多线程实现

import (
	"fmt"
	"github.com/psilva261/timsort/v2"
	"math"
	"sort"
	"sync"
	"time"
)

func MulitSortC(cpuNums int, arr[]float64, worker func( index [2]int,arr []float64,group *sync.WaitGroup))[]float64 {
	fmt.Println("Start sorting...")
	startTime := time.Now()
	zeroNums:=0
	if(len(arr)%cpuNums!=0){
		zeroNums=cpuNums-len(arr)%cpuNums
	}
	for i := 0; i < zeroNums; i++ {
		arr=append(arr,0.0)
	}

	length := len(arr)

	var wg sync.WaitGroup
	perLen := int(length / cpuNums)
	start := 0
	var comB []int
	for i := 0; i < cpuNums; i++ {
		if i == cpuNums-1 {
			ch := [2]int{start, length}
			comB = append(comB, start)
			comB = append(comB, length)
			//fmt.Println(ch)
			wg.Add(1)
			go worker(ch,arr,&wg)
			continue
		}
		ch := [2]int{start, start + perLen}
		comB = append(comB, start)
		comB = append(comB, start+perLen)
		wg.Add(1)
		go worker(ch,arr,&wg)
		start += perLen
	}
	wg.Wait()
	fmt.Println("multithreading sorting costs(s):", time.Since(startTime).Seconds())


	startTime = time.Now()

	mergeTimes:=calcMergeTimes(cpuNums)
	combTimes:=mergeTimes-2

	for i := 0; i < mergeTimes; i++ {
		nowLen := int(math.Exp2(float64(i+2)))
		nextTimes:=int(math.Exp2(float64(combTimes)+1))
		comB_st:=0
		mid:=0
		for j := 0; j <nextTimes; j += 1 {
			comB_st=Min2Int(comB_st,len(comB)-1)
			comB_ed:=Min2Int(comB_st+nowLen,len(comB))-1
			st := comB[comB_st]
			ed := comB[comB_ed] -1
			//mid=(st+ed)/2
			mid_index:=Min2Int(nowLen/2+comB_st,len(comB)-1)
			mid = comB[mid_index] - 1

			if i==mergeTimes-1{
				currLen:=nowLen/2
				mid=comB[currLen]-1
			}
			comB_st+=nowLen
			ch := [3]int{st, mid, ed}
			wg.Add(1)
			go workerMerge2(ch,arr,&wg)
		}
		combTimes-=1
		wg.Wait()
	}
	arr=append(arr[zeroNums:])
	fmt.Println("merging costs(s):", time.Since(startTime).Seconds())
	return arr
}
func GetSortNames()[]string{
	SortNames:=[]string{"BubbleSort","InsertSort","SelectSort","ShellSort","MergeSort","QuickSort","HeapSort","CountSort","BucketSort","RadixSort","TimSort"}
	return SortNames
}
func ContainsInSlice(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func MultiSort(thdNums int,arr []float64,method string)[]float64{
	switch method {
	case "BubbleSort":
		arr=MulitSortC(thdNums,arr,workerBubbleSort)
	case "InsertSort":
		arr=MulitSortC(thdNums,arr,workerInsertSort)
	case "SelectSort":
		arr=MulitSortC(thdNums,arr,workerSelectSort)
	case "ShellSort":
		arr=MulitSortC(thdNums,arr,workerShellSort)
	case "MergeSort":
		arr=MulitSortC(thdNums,arr,workerMergeSort)
	case "QuickSort":
		arr=MulitSortC(thdNums,arr,workerQuickSortU)
	case "HeapSort":
		arr=MulitSortC(thdNums,arr,workerHeapSort)
	case "CountSort":
		arr=MulitSortC(thdNums,arr,workerCountSort)
	case "BucketSort":
		arr=MulitSortC(thdNums,arr,workerBucketSort)
	case "RadixSort":
		arr=MulitSortC(thdNums,arr,workerRadixSort)
	case "TimSort":
		arr=MulitSortC(thdNums,arr,workerTimSort)
	}
	return arr
}

func SingleSort(arr []float64,method string){
	switch method {
	case "BubbleSort":
		BubbleSort(arr)
	case "InsertSort":
		InsertSort(arr)
	case "SelectSort":
		SelectSort(arr)
	case "ShellSort":
		ShellSort(arr)
	case "MergeSort":
		MergeSort(arr)
	case "QuickSort":
		QuickSortU(arr)
	case "HeapSort":
		HeapSort(arr)
	case "CountSort":
		CountSort(arr)
	case "BucketSort":
		BucketSort(arr,len(arr)/1000)
	case "RadixSort":
		RadixSort(arr)
	case "TimSort":
		timsort.TimSort(sort.Float64Slice(arr))
	}
}
// 验证数组是否有序
func IsSorted(arr []float64) int {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			//fmt.Println("无序索引","：",i-1,i,".数值：",arr[i-1],arr[i],"")
			fmt.Println("Order Check:Array Disordered")
			return 0
		}
	}
	fmt.Println("Order Check:Array Ordered")
	return 1
}

// 校验数组 元素是否被修改
func Si_eMo(arr []float64)int {
	count := 0
	for i := 0; i < len(arr); i++ {
		d:= int(arr[i]*1000)
		if d%2 == 0 {
			count +=d
		} else {
			count -= d
		}
	}
	fmt.Println("eMo:", count)
	return count
}