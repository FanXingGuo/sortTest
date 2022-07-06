package main

//常见排序方法 单排序实现

import (
	"day03/SortTest/sortMthd"
	"day03/SortTest/utilTol"
	"flag"
	"fmt"
	"time"
)

func main() {
	var input string
	var output string
	var method string
	var help bool



	flag.StringVar(&input,"i","","输入文件名")
	flag.StringVar(&output,"o","","输出文件名")
	flag.StringVar(&method,"m","","排序方法")
	flag.BoolVar(&help,"h",false,"帮助信息")

	flag.Parse()
	if help{
		fmt.Println("=== === === === === === === === === === === ")
		fmt.Println("对文件排序,然后保存结果")
		fmt.Println("参数要求：")
		fmt.Println("-i 输入文件名称 （文本要求每行一个数）")
		fmt.Println("-m 指定排序方法,方法如下:BubbleSort InsertSort SelectSort ShellSort MergeSort QuickSort HeapSort CountSort BucketSort RadixSort TimSort")
		fmt.Println("-o 指定输出文件名")
		fmt.Println("例如：")
		fmt.Println("./")
		fmt.Println("=== === === === === === === === === === === ")
		return
	}
	if !utilTol.CheckFileIsExist(input){
		fmt.Println("文件不存在")
		return
	}
	if !sortMthd.ContainsInSlice(sortMthd.GetSortNames(),method){
		fmt.Println("请指定如下名称之一：")
		fmt.Println(sortMthd.GetSortNames())
		return
	}
	start := time.Now()
	content:=utilTol.ReadFile(input)
	fmt.Println("读耗时:", time.Since(start).Seconds())

	start = time.Now()
	arr:=utilTol.Str2Arr(content)
	fmt.Println("转数组耗时:", time.Since(start).Seconds())

	eMo:=sortMthd.Si_eMo(arr)

	start = time.Now()
	sortMthd.SingleSort(arr,method)
	fmt.Println("排序耗时:", time.Since(start).Seconds())
	sortMthd.IsSorted(arr)

	//cont:=utilTol.Arr2Str(arr)
	//utilTol.WriteFile(output,cont)
	eMo2:=sortMthd.Si_eMo(arr)
	if eMo==eMo2{
		fmt.Println("eMo校验:相等")
	}else {
		fmt.Println("eMo校验:不相等")
		fmt.Println("eMo1:",eMo)
		fmt.Println("eMo2:",eMo2)
	}
	start = time.Now()
	cont:=utilTol.Arr2Str(arr)
	fmt.Println("转文本耗时:", time.Since(start).Seconds())

	start = time.Now()
	utilTol.WriteFile(output,cont)
	//utilTol.WriteArr(output,arr)
	fmt.Println("写耗时:", time.Since(start).Seconds())


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
