package main

import (
	"fmt"
	"github.com/psilva261/timsort/v2"
	"sort"
)

func main(){
	l:=[]string{"c","a","b"}
	timsort.TimSort(sort.StringSlice(l))
	fmt.Println(l)

	n:=[]int{3,1,4}
	timsort.TimSort(sort.IntSlice(n))
	fmt.Println(n)
}