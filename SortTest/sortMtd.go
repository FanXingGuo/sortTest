package main

import (
	"day03/SortTest/dtbM"
	"day03/SortTest/sortMthd"
	"fmt"
	"math"
	"sync"
	"time"
)

func worker1(id int, jobs <-chan [2]int, results chan<- int, arr []int) {
	indexs := <-jobs
	sortMthd.BubbleSort(arr[indexs[0]:indexs[1]])
	results <- 1
}
func worker2(indexs[2]int, arr []int,group *sync.WaitGroup) {
	defer group.Done()
	sortMthd.QuickSortU(arr[indexs[0]:indexs[1]])
}
func normFunction(x int) float64 {
	return dtbM.NormalFloat64(x, 50, 3)
}
func normFunctionInt(x int) int{
	return int(dtbM.NormalFloat64(x, 100, 1)*float64(1000))
}
func PosiFunc(x int)float64{
	return dtbM.PoisDisFloat64(x,8)
}
func GeoFunc(x int)float64{
	return dtbM.GeoDisFloat64(x,0.1)
}
func lineFunc(x int) int {
	return 50
}
func la_lambdax(x int) int {
	y := 0.1 * math.Pow(math.E, -0.1*float64(x))
	return int(y * 1000)
}

// slice-sort-in-go/sort_int_slice.go
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

//func main() {
//	sl := IntSlice([]int{89, 14, 8, 9, 17, 56, 95, 3})
//	fmt.Println(sl) // [89 14 8 9 17 56 95 3]
//	sort.Sort(sl)
//	fmt.Println(sl) // [3 8 9 14 17 56 89 95]
//}


func main() {
	//normInt := dtbM.DistRandSelectInt(200, normFunctionInt)
	//selectedLists := make([]int, 0)
	//for i := 0; i < 10000; i++ {
	//	selectedLists = append(selectedLists, normInt())
	//
	//}
	//cont:=utilTol.Arr2Str(selectedLists)
	//utilTol.WriteFile("data.txt",cont)
	//selectedLists:=dtbM.DistArr(100,600,normFunction)
	//selectedLists:=dtbM.ODR_ARR(10000000,"normal")
	//selectedLists:=dtbM.RepeatDis(1000,1)
	selectedLists:=dtbM.ODR_ARR(50000,"rep_100%")

	//fmt.Println(selectedLists)
	//dtbM.ShowDist(selectedLists)
	fmt.Println(dtbM.CountArr(selectedLists))


	sortMthd.IsSorted(selectedLists)
	sortMthd.Si_eMo(selectedLists)

	start := time.Now()
//[0, 6, 11, 11, 18, 25, 28, 37, 40, 45, 47, 56, 59, 62, 66, 74, 81, 81, 87, 89, 94, 95]
	//sortMthd.InsertSort(selectedLists)
	//sortMthd.SelectSort(selectedLists)
	//sortMthd.BubbleSort(selectedLists)
	//sortMthd.ShellSort(selectedLists)
	//sortMthd.QuickSort(selectedLists, 0, len(selectedLists)-1)
	//sort.Sort(selectedLists)
	//sortMthd.MergeSort(selectedLists)
	//sortMthd.HeapSort(selectedLists)
	//sortMthd.CountSort(selectedLists)
	//sortMthd.ShellSort(selectedLists)
	sortMthd.BucketSort(selectedLists,100)



	//selectedLists=sortMthd.MulitSortC(8, selectedLists, worker2)

	fmt.Println("该函数执行完成耗时：", time.Since(start))
	sortMthd.Si_eMo(selectedLists)
	sortMthd.IsSorted(selectedLists)

	// 指定函数分布
	//x1 := make([]int, 100)
	//freList := make([]int, 0)
	//selectLists := make([]int, 0)
	//for i := 0; i < len(x1); i++ {
	//	x1[i] = i
	//}
	//// 以x为值y为频数，生成数组
	//for i := 0; i < len(x1); i++ {
	//	times := int(dtbM.NormalFloat64(x1[i], 50, 10) * 1000)
	//
	//	for j := 0; j < times; j++ {
	//		freList = append(freList, x1[i])
	//	}
	//}
	////dtbM.ShowDist(freList)
	//rand.Seed(1)
	//for i := 0; i < 3000; i++ {
	//	selectLists = append(selectLists, freList[rand.Intn(len(freList))])
	//}
	//dtbM.ShowDist(selectLists)

	// 排序
	//const length = 1000 * 1
	////cpuNums := runtime.NumCPU()  CPU核数 为偶数
	//
	//var arr []int
	//rand.Seed(1)
	//for i := 0; i < length; i++ {
	//	//arr = append(arr, rand.Intn(10))
	//	arr = append(arr, int(math.Abs(rand.NormFloat64()*10000))%100)
	//}
	//fmt.Println()

	//sortMthd.IsSorted(arr)
	//sortMthd.Si_eMo(arr)
	//
	//start := time.Now()

	//sortMthd.MulitSort(1, arr, worker1)
	//sortMthd.ShellSort(arr)

	//elapsed := time.Since(start)

	//fmt.Println("该函数执行完成耗时：", elapsed)

	//sortMthd.IsSorted(arr)
	//sortMthd.Si_eMo(arr)

	//====

	//fmt.Println(arr)
	//sortMthd.BubbleSort(arr[:10])
	//sortMthd.SelectSort(arr)
	//sortMthd.InsertSort(arr)
	//sortMthd.ShellSort(arr)
	//sortMthd.MergeSort(arr)
	//sortMthd.QuickSort(arr, 0, len(arr)-1) //全闭区间
	//sortMthd.HeapSort(arr)
	//sortMthd.CountSort(arr)
	//sortMthd.BucketSort(arr, 10)
	//sortMthd.RadixSort(arr)
	//fmt.Println(arr)
	//exp := math.Floor(math.Log2(9)) + 1
	//fmt.Println(math.Exp2(exp))
	//
	//fmt.Println(runtime.NumCPU())

}
