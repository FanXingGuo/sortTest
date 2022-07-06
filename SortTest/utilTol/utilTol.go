package utilTol

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func CheckFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
func ReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	return string(content)
}

func WriteFile(path string,content string){
	cont := []byte(content)
	err := ioutil.WriteFile(path, cont, 0644)
	if err != nil {
		panic(err)
	}
}
func Arr2Str(arr []int)string{
	//content:=""
	//
	//for _,v :=range arr{
	//	str:=strconv.Itoa(v)
	//	content+=str+"\n"
	//}
	//return content
	var temp = make([]string, len(arr))
	for k, v := range arr {
		temp[k] = fmt.Sprintf("%d", v)
	}
	var result = strings.Join(temp, "\n")
	return result
}

func WriteArr(path string,arr []int){

	var temp = make([]string, len(arr))
	for k, v := range arr {
		temp[k] = fmt.Sprintf("%d", v)
	}
	var result = strings.Join(temp, "\n")
	fmt.Println(result)
}
func Str2Arr(s string)[]int{
	a:=strings.Split(s,"\n")
	arr:=make([]int,0)
	for _,v :=range a{
		if(len(v)==0){
			continue
		}
		int_s, err := strconv.Atoi(v)
		if err!=nil{
			panic(err)
		}
		arr=append(arr,int_s)
	}
	return arr
}
