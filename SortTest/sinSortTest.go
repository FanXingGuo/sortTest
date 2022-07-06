package main

import (
	"day03/SortTest/dtbM"
	"day03/SortTest/sortMthd"
	"flag"
	"fmt"
	"time"
)

func main() {
	var length int
	var distri string
	var method string
	var help bool
	methods:=[]string{"order","disorder","reverse"}



	flag.IntVar(&length,"l",0,"数组长度")
	flag.StringVar(&distri,"d","","数组顺序")
	flag.StringVar(&method,"m","","排序方法")
	flag.BoolVar(&help,"h",false,"帮助信息")

	flag.Parse()
	if help{
		fmt.Println("=== === === === === === === === === === === ")
		fmt.Println("对文件排序,然后保存结果")
		fmt.Println("参数要求：")
		fmt.Println("-l 数组长度")
		fmt.Println("-d 数组顺序:有序、无序、逆序。分别对应",method)
		fmt.Println("-m 指定排序方法,方法如下:BubbleSort InsertSort SelectSort ShellSort MergeSort QuickSort HeapSort CountSort BucketSort RadixSort")
		fmt.Println("例如：")
		fmt.Println("./sortNums -l 1000 -d order -n 2 -m BubbleSort1")
		fmt.Println("=== === === === === === === === === === === ")
		return
	}

	if length<0{
		fmt.Println("长度必须大于0")
		return
	}
	if ! sortMthd.ContainsInSlice(methods,distri){
		fmt.Println("请指定如下方法之一")
		fmt.Println(methods)
		return
	}
	if !sortMthd.ContainsInSlice(sortMthd.GetSortNames(),method){
		fmt.Println("请指定如下名称之一：")
		fmt.Println(sortMthd.GetSortNames())
		return
	}

	arr:=dtbM.ODR_ARR(length,distri)
	fmt.Println("随机数已生成：",distri,"长度",length)
	eMo:=sortMthd.Si_eMo(arr)


	start := time.Now()
	sortMthd.SingleSort(arr,method)
	fmt.Println("总耗时:", time.Since(start).Seconds())

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
