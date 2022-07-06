package dtbM

import (
	"math/rand"
)
// 返回指定类型随机数
func DisOrderArr(length int)[]float64{
	arr:=make([]float64,0)
	for i := 0; i < length; i++ {
		arr=append(arr,rand.Float64())
	}
	return arr
}

func ODR_ARR(length int,method string)[]float64{
	switch method {
	case "disorder":
		return DisOrderArr(length)
	}
	return nil
}

func GetDistribute()[]string{
	methods:=[]string{"disorder"}
	return methods
}

