package main

//带文件读写功能的，排序

import (
	"day03/SortTest/sortMthd"
	"day03/SortTest/utilTol"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	var input string
	var method string
	var output string
	var cpuNums int
	var help bool


	flag.StringVar(&input,"i","","输入文件名")
	flag.IntVar(&cpuNums,"n",2,"线程数")
	flag.StringVar(&method,"m","","排序方法")
	flag.StringVar(&output,"o","","输出文件名")
	flag.BoolVar(&help,"h",false,"帮助信息")

	flag.Parse()
	if help{
		fmt.Println("=== === === === === === === === === === === ")
		fmt.Println("Sort the numbers in the file and save the result to a specific file.")
		fmt.Println("Usage：")
		fmt.Println("-i Input file name（one number per line）")
		fmt.Println("-n The Number of Thread")
		fmt.Println("-m Sorting name: BubbleSort InsertSort SelectSort ShellSort MergeSort QuickSort HeapSort CountSort BucketSort RadixSort","TimSort")
		fmt.Println("-o Output file name")
		fmt.Println("Example:")
		fmt.Println("./sortNums -i data.txt -n 2 -m BubbleSort -o resu.txt")
		fmt.Println("=== === === === === === === === === === === ")
		return
	}
	if !utilTol.CheckFileIsExist(input){
		fmt.Println("File doesn't exist.")
		return
	}
	if !sortMthd.ContainsInSlice(sortMthd.GetSortNames(),method){
		fmt.Println("Please specify the sorting methods as follows:")
		fmt.Println(sortMthd.GetSortNames())
		return
	}
	content:=utilTol.ReadFile(input)
	log.Println("Reading file finished.")

	arr:=utilTol.Str2Arr(content)
	eMo:=sortMthd.Si_eMo(arr)

	start := time.Now()
	arr=sortMthd.MultiSort(cpuNums,arr,method)
	fmt.Println("Sorting Costs(s):", time.Since(start).Seconds())

	sortMthd.IsSorted(arr)
	cont:=utilTol.Arr2Str(arr)

	eMo2:=sortMthd.Si_eMo(arr)
	if eMo==eMo2{
		fmt.Println("eMoCheck:Equal")
	}else {
		fmt.Println("eMoCheck:Not Equal")
		//fmt.Println("eMo1:",eMo)
		//fmt.Println("eMo2:",eMo2)
	}

	log.Println("Start Writing...")
	utilTol.WriteFile(output,cont)
	log.Println("Writing file finished.")




	//
	//
	//fmt.Println("input:",input)
	//
	//file,err :=os.Open(input)
	//if err!=nil{
	//	panic(err)
	//}
	//defer file.Close()
	//content,err:=ioutil.ReadAll(file)
	//fmt.Println(string(content))



}
