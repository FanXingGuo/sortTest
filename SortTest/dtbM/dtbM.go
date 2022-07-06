package dtbM

import (
	"day03/SortTest/sortMthd"
	"fmt"
	"math"
	"math/rand"
)

type Poke struct {
	num int
}
type Pokes []Poke

func ShowDist(arr []int) {
	max, min := sortMthd.MaxAndMin(arr)
	ctArr := make([]int, max-min+1)
	for _, value := range arr {
		ctArr[value-min]++
	}
	for index, value := range ctArr {
		if value != 0 {
			fmt.Printf("%v:", index+min)
			for i := 0; i < value; i++ {
				fmt.Printf("*")
			}
			fmt.Printf("\n")
		}
	}
}

func CountArr(arr []int)[]int {
	max, _ := sortMthd.MaxAndMin(arr)
	ctArr := make([]int, max+1)
	for _, value := range arr {
		ctArr[value]++
	}
	return ctArr
}

func Factorial(n int) int  {
	v:=1
	for i := 1; i <=n ; i++ {
		v=v*i
	}
	return v
}
func NormalFloat64(x int, miu int, sigma int) float64 {
	randomNormal := (1 / (float64(sigma) * math.Sqrt(2*math.Pi))) * math.Exp(-0.5*math.Pow(float64(x-miu)/float64(sigma), 2))
	return randomNormal
}
func NormalXFloat64(x float64, miu float64, sigma float64) float64 {
	randomNormal := (1 / (float64(sigma) * math.Sqrt(2*math.Pi))) * math.Exp(-0.5*math.Pow(float64(x-miu)/float64(sigma), 2))
	return randomNormal
}
func GeoDisFloat64(x int,p float64)float64{
	q:=1-p
	return math.Pow(q,float64(x)-1)*p
}
func PoisDisFloat64(x int,lambda float64)float64  {
	return (math.Exp(-lambda)*math.Pow(lambda,float64(x)))/float64(Factorial(x))
}

func NormFunction(x int) float64 {
	return NormalFloat64(x, 50, 3)
}
func normFunctionInt(x int) int{
	return int(NormalFloat64(x, 100, 1)*float64(1000))
}
func NormFunctionFloat5_1(x float64) float64{
	return NormalXFloat64(x, 5, 1)
}
func PosiFunc(x int)float64{
	return PoisDisFloat64(x,20)
}
func GeoFunc(x int)float64{
	return GeoDisFloat64(x,0.1)
}


func DistRandSelectInt(xLen int, function func(x int) int) func() int {
	// 给出指定函数分布 随机数生成器 (取值范围 [0,xLen)) 频数 为指定的 function函数返回值
	x1 := make([]int, xLen)
	freList := make([]int, 0)
	for i := 0; i < len(x1); i++ {
		x1[i] = i
	}
	// 以x为值y为频数，生成数组
	for i := 0; i < len(x1); i++ {
		times := function(x1[i])
		for j := 0; j < times; j++ {
			freList = append(freList, x1[i])
		}
	}
	return func() int {
		return freList[rand.Intn(len(freList))]
	}
}
func DistRSFloatInt(xMax int,step float64,enlar float64, function func(x float64) float64) func() int {
	// 给出指定函数分布 随机数生成器 (取值范围 [0,xLen)) 频数 为指定的 function函数返回值
	x1 := make([]float64,0)
	freList := make([]int, 0)
	xEar:=float64(1)/step
	for i := step;i < float64(xMax); i+=step {
		x1=append(x1,i)
	}
	// 以x为值y为频数，生成数组
	for i := 0; i < len(x1); i++ {
		times := int(function(x1[i])*enlar)
		//fmt.Println(times,x1[i],x1[i]*xEar,int(math.Floor(x1[i]*xEar)))
		for j := 0; j < times; j++ {
			freList = append(freList, int(math.Floor(x1[i]*xEar+0.01)))
		}
	}
	//ctArr:=CountArr(freList)
	//fmt.Println("ori:",ctArr)

	return func() int {
		return freList[rand.Intn(len(freList))]
	}
}
func DistArr(xMax int,ArrLen int, function func(x int) float64)[]int{
	x1 := make([]int, xMax)
	freList := make([]int, 0)
	for i := 1; i <= xMax; i++ {
		x1[i-1] = i
	}
	//fmt.Println(x1)
	// 以x为值y为频数，生成数组
	flag:=true
	last:=0.1
	for i := 0; i < xMax; i++ {
		times:=0
		fmt.Println(function(x1[i]))
		if flag {
			mz:=function(x1[i])
			if mz<float64(1/(ArrLen*10)){
				times =0
			}else {
				times=int(function(x1[i])*float64(ArrLen))
			}
		}
		if times==0 && function(x1[i])*float64(ArrLen)-last<0{
			flag=false
		}
		last=function(x1[i])*float64(ArrLen)
		//times=int(function(x1[i])*float64(ArrLen))
		for j := 0; j < times; j++ {
			freList = append(freList, x1[i])
		}
	}
	length:=len(freList)

	toAddLen:=ArrLen-length
	for i := 0; i < toAddLen; i++ {
		freList=append(freList,0)
	}
	length=len(freList)
	for i := 0; i < len(freList); i++ {
		idx:=rand.Intn(length)
		temp:=freList[i]
		freList[i]=freList[idx]
		freList[idx]=temp
	}
	return freList

}



func OrderedArr(length int)[]int{
	arr:=make([]int,0)
	for i := 0; i < length; i++ {
		arr=append(arr,i+1)
	}
	return arr
}
func DisOrderArr(length int)[]int{
	arr:=OrderedArr(length)
	for i := 0; i < length; i++ {
		idx:=rand.Intn(length)
		temp:=arr[i]
		arr[i]=arr[idx]
		arr[idx]=temp
	}
	return arr
}
func ReversedArr(length int)[]int  {
	arr:=make([]int,0)
	for i := 0; i < length; i++ {
		arr=append(arr,length-i)
	}
	return arr
}
func RepeatDis(length int,rate float32)[]int{
	ReapLen:=int(float32(length)*rate)-1
	origLen:=length-ReapLen
	origArr:=DisOrderArr(origLen)
	repNum:=origArr[len(origArr)-1]
	for i := 0; i < ReapLen; i++ {
		origArr=append(origArr,repNum)
	}
	lenOfArr:=len(origArr)
	for i := origLen; i < lenOfArr; i++ {
		rnd:=rand.Intn(origLen)

		temp:=origArr[rnd]
		origArr[rnd]=origArr[i]
		origArr[i]=temp
	}
	return origArr
}
//[1,2,3,4,5,6,7,8,9,10]
//[1,2,3,4,5,6,7,8,9,9]

func ODR_ARR(length int,method string)[]int{
	switch method {
	case "order":
		return OrderedArr(length)
	case "disorder":
		return DisOrderArr(length)
	case "reverse":
		return ReversedArr(length)
	case "normal":
		normInt := DistRandSelectInt(200, normFunctionInt)
		selectedLists := make([]int, 0)
		for i := 0; i < length; i++ {
			selectedLists = append(selectedLists, normInt())
		}
		return selectedLists
	case "normal3":
		arr:=make([]int,0)
		selestInt:=DistRSFloatInt(10,0.1,float64(10000),NormFunctionFloat5_1)
		for i := 0; i < length; i++ {
			arr=append(arr,selestInt())
		}
		return arr
	case "geo":
		selectedLists:=DistArr(100,length,GeoFunc)
		return selectedLists
	case "posi":
		selectedLists:=DistArr(100,length,PosiFunc)
		return selectedLists
	case "rep_20%":
		return RepeatDis(length,0.2)
	case "rep_50%":
		return RepeatDis(length,0.5)
	case "rep_80%":
		return RepeatDis(length,0.8)
	case "rep_100%":
		return RepeatDis(length,1)
	}
	return nil
}

func GetDistribute()[]string{
	methods:=[]string{"order","disorder","reverse","normal","normal3","geo","posi","rep_20%","rep_50%","rep_80%","rep_100%"}
	return methods
}

