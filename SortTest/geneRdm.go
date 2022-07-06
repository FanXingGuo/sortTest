package main

import (
	"flag"
	"fmt"
)

//随机数生成，写入文件

import (
	"day03/SortTest/dtbM"
	"day03/SortTest/sortMthd"
	"day03/SortTest/utilTol"
)

func main2()  {
	var length int
	var method string
	var output string
	var help bool
	methods:=[]string{"order","disorder","reverse"}

	flag.IntVar(&length,"l",0,"数组长度")
	flag.StringVar(&method,"m","disorder","数组排序方式")
	flag.StringVar(&output,"o","","指定输出文件")
	flag.BoolVar(&help,"h",false,"帮助")

	flag.Parse()
	if help{
		fmt.Println("随机数生成器")
		fmt.Println("参数说明：")
		fmt.Println("-l 指定随机数组长度")
		fmt.Println("-m 随机数产生方式，顾名思义:",methods)
		fmt.Println("-o 指定输出文件名")
		fmt.Println("-h 查看帮助")
	}
	if length<0{
		fmt.Println("长度必须大于0")
		return
	}
	if ! sortMthd.ContainsInSlice(methods,method){
		fmt.Println("请指定如下方法之一")
		fmt.Println(methods)
		return
	}
	if len(output)==0{
		fmt.Println("请指定输出文件名称")
		return
	}
	arr:=dtbM.ODR_ARR(length,method)
	cont:=utilTol.Arr2Str(arr)
	utilTol.WriteFile(output,cont)

	fmt.Println("写入完成：",output)

}
