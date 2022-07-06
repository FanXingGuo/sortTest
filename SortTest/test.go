package main

import (
	"day03/SortTest/dtbM"
	"day03/SortTest/sortMthd"
	"fmt"
	"time"
)

//type Poke struct {
//	num int
//}
//type Pokes []Poke
//func (poke Pokes)Len()int{
//	return len(poke)
//}
//func (poke Pokes)Less(i,j int)bool{
//	return  poke[i].num<poke[j].num
//}
//func (poke Pokes)swap(i,j int)  {
//	poke[i],poke[j]=poke[j],poke[i]
//}




func main() {
	ch := make(chan struct{}, 1)
	go func() {
		//methods:=[]string{"order","disorder","reverse","normal","geo","posi","rep_20%","rep_50%","rep_80%","rep_100%"}
		fmt.Println("do something...")
		//arr:=dtbM.ODR_ARR(1000000,"disorder")

		//arr:=make([]int,0)
		//selestInt:=dtbM.DistRSFloatInt(10,0.1,float64(10000),dtbM.NormFunctionFloat5_1)
		//for i := 0; i < 1000000; i++ {
		//	arr=append(arr,selestInt())
		//}

		arr:=dtbM.ODR_ARR(1000000,"posi")
		//fmt.Println(arr)

		ctArr:=dtbM.CountArr(arr)
		fmt.Println(len(ctArr),ctArr)

		//ct:=utilTol.Arr2Str(ctArr)
		//utilTol.WriteFile("normal3.txt",ct)

		//fmt.Println("随机数已生成：","disorder","长度",100000)
		eMo:=sortMthd.Si_eMo(arr)

		start := time.Now()
		//arr=sortMthd.MultiSort(2,arr,"BubbleSort")
		//timsort.TimSort(sort.IntSlice(arr))
		//sortMthd.BucketSort(arr,10000)
		//sortMthd.BubbleSort(arr)
		//fmt.Println(dtbM.Factorial(6))
		fmt.Println("总耗时:", time.Since(start).Seconds())
		//fmt.Println(arr)
		sortMthd.IsSorted(arr)

		eMo2:=sortMthd.Si_eMo(arr)
		if eMo==eMo2{
			fmt.Println("eMo校验:相等")
		}else {
			fmt.Println("eMo校验:不相等")
			fmt.Println("eMo1:",eMo)
			fmt.Println("eMo2:",eMo2)
		}

		ch<- struct{}{}
	}()

	select {
	case <-ch:
		fmt.Println("done")
	case <-time.After(160*time.Second):
		fmt.Println("timeout")
	}
}