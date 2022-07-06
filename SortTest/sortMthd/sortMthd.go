package sortMthd

import (
	"fmt"
	"github.com/psilva261/timsort/v2"
	"math"
	"sort"
	"sync"
	"time"
)

type Poke struct {
	num int
}

type Pokes []Poke

func MulitSortA(cpuNums int, arr []int, worker func(id int, jobs <-chan [2]int, results chan<- int, arr []int)) {
	length := len(arr)
	jobs := make(chan [2]int, cpuNums)
	results := make(chan int, cpuNums)
	perLen := int(length / cpuNums)
	exp := Min2Exp(cpuNums)
	mergeTimes := exp
	start := 0
	var comB []int

	// 启动CPU核数 协程
	for i := 0; i < cpuNums; i++ {
		go worker(i, jobs, results, arr)
	}
	// 布置任务
	for i := 0; i < cpuNums; i++ {
		if i == cpuNums-1 {
			ch := [2]int{start, length}
			comB = append(comB, start)
			comB = append(comB, length)
			//fmt.Println(ch)
			jobs <- ch
			continue
		}
		ch := [2]int{start, start + perLen}
		comB = append(comB, start)
		comB = append(comB, start+perLen)

		//fmt.Println(ch)
		jobs <- ch
		start += perLen
	}
	close(jobs)
	for a := 1; a <= cpuNums; a++ {
		<-results
	}
	//fmt.Println("分组合并：", arr)
	// 关闭通道，阻塞等待
	//jobs := make(chan [2]int, cpuNums)
	//归并合并
	var wg sync.WaitGroup
	jobs1 := make(chan [3]int, cpuNums)
	for i := 0; i < cpuNums; i++ {
		//go workerMerge(i, jobs1, results, arr,&wg)
	}
	//fmt.Println("perLen:", perLen)
	//fmt.Println(comB)
	//fmt.Println(mergeTimes)
	for i := 2; i <= mergeTimes+1; i++ {
		nowLen := int(math.Exp2(float64(i)))
		for j := 0; j < len(comB); j += nowLen {
			st := comB[j]
			ed := comB[j+nowLen-1] - 1
			mid := comB[nowLen/2+j] - 1
			//fmt.Println("comb:", st, mid, ed)
			ch := [3]int{st, mid, ed}
			jobs1 <- ch
			wg.Add(1)
		}
		wg.Wait()
	}
	close(jobs1)
	for a := 1; a <= cpuNums; a++ {
		<-results
	}
	//for i := 1; i <= mergeTimes; i++ {
	//	nowLen := int(math.Exp2(float64(i))) * perLen
	//	nextTimes := int(math.Exp2(float64(mergeTimes - i)))
	//	arrSt := 0
	//	for j := 0; j < nextTimes; j++ {
	//		st := Min2Int(arrSt, length)
	//		ed := Min2Int(int(arrSt+nowLen), length)
	//		mid := int((st + ed) / 2)
	//		if j == nextTimes-1 {
	//			mid = Min2Int(int(arrSt+nowLen), length)
	//			ed = length
	//		}
	//		ch := [3]int{st, mid, ed}
	//		fmt.Println(ch)
	//		jobs1 <- ch
	//		arrSt += nowLen
	//	}
}

func MulitSortB(cpuNums int, arr[]int, worker func(id int, jobs <-chan [2]int, results chan<- int, arr []int))[]int {

	fmt.Println("origin:",arr)
	zeroNums:=cpuNums-len(arr)%cpuNums
	for i := 0; i < zeroNums; i++ {
		arr=append(arr,0)
	}

	length := len(arr)
	jobs := make(chan [2]int, cpuNums)
	results := make(chan int, cpuNums)
	perLen := int(length / cpuNums)
	start := 0
	var comB []int
	//[0, 6, 11, 11, 18, 25, 28, 37, 40, 45, 47, 56, 59, 62, 66, 74, 81, 81, 87, 89, 94, 95]

	// 启动CPU核数 协程
	for i := 0; i < cpuNums; i++ {
		go worker(i, jobs, results, arr)
	}
	// 布置任务
	for i := 0; i < cpuNums; i++ {
		if i == cpuNums-1 {
			ch := [2]int{start, length}
			comB = append(comB, start)
			comB = append(comB, length)
			//fmt.Println(ch)
			jobs <- ch
			continue
		}
		ch := [2]int{start, start + perLen}
		comB = append(comB, start)
		comB = append(comB, start+perLen)

		//fmt.Println(ch)
		jobs <- ch
		start += perLen
	}
	close(jobs)
	for a := 1; a <= cpuNums; a++ {
		<-results
	}

	fmt.Println("star:",arr)

	mergeTimes:=calcMergeTimes(cpuNums)

	combTimes:=mergeTimes-2

	var wg sync.WaitGroup

	//fmt.Println(comB)
	for i := 0; i < mergeTimes; i++ {
		nowLen := int(math.Exp2(float64(i+2)))
		nextTimes:=int(math.Exp2(float64(combTimes)+1))
		comB_st:=0
		mid:=0
		//println("nowLen:",nowLen,"nextTimes:",nextTimes)
		for j := 0; j <nextTimes; j += 1 {
			comB_st=Min2Int(comB_st,len(comB)-1)
			comB_ed:=Min2Int(comB_st+nowLen,len(comB))-1

			st := comB[comB_st]
			ed := comB[comB_ed] -1
			mid=(st+ed)/2
			if i==mergeTimes-1{
				currLen:=nowLen/2
				mid=comB[currLen]-1
			}
			//println("comB_st:",comB_st,"mid:",mid,"comB_ed:",comB_ed,"comb:", st, mid, ed)
			comB_st+=nowLen

			ch := [3]int{st, mid, ed}
			//wg.Add(1)
			workerMerge2(ch,arr,&wg)
			//fmt.Println("resu:",arr)
		}
		combTimes-=1
		//wg.Wait()
	}
	arr=append(arr[zeroNums:])
	fmt.Println("end1:",arr)
	return arr
}

func MulitSortC(cpuNums int, arr[]int, worker func( index [2]int,arr []int,group *sync.WaitGroup))[]int {
	fmt.Println("Start sorting...")
	startTime := time.Now()
	zeroNums:=0
	if(len(arr)%cpuNums!=0){
		zeroNums=cpuNums-len(arr)%cpuNums
	}
	for i := 0; i < zeroNums; i++ {
		arr=append(arr,0)
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

func MultiSort(thdNums int,arr []int,method string)[]int{
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

func SingleSort(arr []int,method string){
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
		timsort.TimSort(sort.IntSlice(arr))
	}
}

func IsSorted(arr []int) int {
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
func IsSortedPoke(pokes[]Poke)int{
	for i := 1; i < len(pokes); i++ {
		if pokes[i-1].num>pokes[i].num{
			fmt.Println("有序性：无序")
			return 0
		}
	}
	fmt.Println("有序性：有序")
	return 1
}
func Si_eMo(arr []int)int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			count += arr[i]
		} else {
			count -= arr[i]
		}
	}
	fmt.Println("eMo:", count)
	return count
}

