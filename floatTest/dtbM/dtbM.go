package dtbM

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)
// 返回指定类型随机数
func DisOrderArr(length int)[]float64{
	arr:=make([]float64,0)
	for i := 0; i < length; i++ {
		arr=append(arr,rand.Float64())
	}
	return arr
}
func DisOrderArr2(length int)[]float64{
	arr:=make([]float64,0)
	for i := 0; i < length; i++ {
		a:=rand.Float64()
		b:=fmt.Sprintf("%.16f", a)
		c,_ := strconv.ParseFloat(b,64)
		arr=append(arr,c)
	}
	return arr
}
func DisOrderArr3(length int)[]float64{
	arr:=make([]float64,0)
	for i := 0; i < length; i++ {
		arr=append(arr,float64(rand.Int31())*math.Pow(0.1,13))
	}
	return arr
}

func ODR_ARR(length int,method string)[]float64{
	switch method {
	case "disorder":
		return DisOrderArr(length)
	case "disorder16B":
		return DisOrderArr2(length)
	case "disorder10E":
		return DisOrderArr3(length)

	}
	return nil
}

func GetDistribute()[]string{
	methods:=[]string{"disorder","disorder16B","disorder10E"}
	return methods
}

