package sortMthd

//常见排序方法 实现

import (
	"math"
	"sync"
)

func BubbleSort(arr []float64) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
}

func MaxAndMin(arr []float64) (max float64, min float64) {
	max = arr[0]
	min = arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}
func Min2Int(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func Max2Int(a float64, b float64) float64 {
	if a < b {
		return b
	} else {
		return a
	}
}

func swap(a *float64, b *float64) {
	tmp := *a
	*a = *b
	*b = tmp
}
func InsertSort(arr []float64) {
	length := len(arr)
	for i := 1; i < length; i++ {
		tmp := arr[i]
		j := i - 1
		for j >= 0 && tmp <= arr[j] {
			arr[j+1] = arr[j]
			j -= 1
		}
		arr[j+1] = tmp
	}
}
func SelectSort(arr []float64) {
	length := len(arr)
	for left, right := 0, length-1; left < right; left, right = left+1, right-1 {
		min, max := left, right
		for index := left; index <= right; index++ {
			if arr[index] < arr[min] {
				min = index
			}
			if arr[index] > arr[max] {
				max = index
			}
		}
		swap(&arr[left], &arr[min])
		if left == max {
			max = min
		}
		swap(&arr[right], &arr[max])
	}
}
func ShellSort(arr []float64) {
	length := len(arr)
	d := length / 2
	k, tmp := 0, 0.0
	for d > 0 {
		for i := d; i < length; i++ {
			if arr[i] < arr[i-d] {
				tmp = arr[i]
				k = i - d
				for k >= 0 && arr[k] > tmp {
					arr[k+d] = arr[k]
					k -= d
				}
				arr[k+d] = tmp
			}
		}
		d = d / 2
	}
}
func Merge(arr []float64, low int, mid int, high int) {
	i, j := low, mid+1
	k := 0
	b := make([]float64, high-low+1)
	for i <= mid && j <= high {
		if arr[i] < arr[j] {
			b[k] = arr[i]
			k, i = k+1, i+1
		} else {
			b[k] = arr[j]
			k, j = k+1, j+1
		}
	}
	for i <= mid {
		b[k] = arr[i]
		k, i = k+1, i+1
	}
	for j <= high {
		b[k] = arr[j]
		k, j = k+1, j+1
	}
	k = 0
	for low <= high {
		arr[low] = b[k]
		low, k = low+1, k+1
	}
}
func merge_pass(arr []float64, len int, k int) {
	i := 0
	for i < len-2*k+1 {
		Merge(arr, i, i+k-1, i+2*k-1)
		i = i + 2*k
	}
	if i < len-k+1 {
		Merge(arr, i, i+k-1, len-1)
	}
}
func MergeSort(arr []float64) {
	length := len(arr)
	k := 1
	for k <= length {
		merge_pass(arr, length, k)
		k = k * 2
	}
}
func partion(arr []float64, start int, end int) int {
	i, j := start-1, start
	for ; j < end; j++ {
		if arr[j] < arr[end] {
			i++
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}
	i++
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
	return i
}
func QuickSort(arr []float64, left int, right int) {
	if left < right {
		center := partion(arr, left, right)
		QuickSort(arr, left, center-1)
		QuickSort(arr, center+1, right)
	}
}
func QuickSortU(arr[]float64){
	length:=len(arr)
	start:=0
	end:=length-1
	stack := make([]int, end+1)
	stack_top:=0

	if start<end{
		stack[stack_top]=end
		stack_top+=1
		stack[stack_top]=start
		stack_top+=1
		for stack_top>0 {
			stack_top-=1
			left:=stack[stack_top]
			stack_top-=1
			right:=stack[stack_top]
			index:=partion(arr,left,right)
			if index-1>left{
				stack[stack_top]=index-1
				stack_top+=1
				stack[stack_top]=left
				stack_top+=1
			}
			if(right>index+1){
				stack[stack_top]=right
				stack_top+=1
				stack[stack_top]=index+1
				stack_top+=1
			}
		}
	}
}
func adjuest(arr []float64, length int, index int) {
	left := 2*index + 1
	right := 2*index + 2
	maxIdx := index
	if left < length && arr[left] > arr[maxIdx] {
		maxIdx = left
	}
	if right < length && arr[right] > arr[maxIdx] {
		maxIdx = right
	}
	if maxIdx != index {
		temp := arr[maxIdx]
		arr[maxIdx] = arr[index]
		arr[index] = temp
		adjuest(arr, length, maxIdx)
	}
}
func HeapSort(arr []float64) {
	length := len(arr)
	for i := length/2 - 1; i >= 0; i-- {
		adjuest(arr, length, i)
	}
	for i := length - 1; i >= 1; i-- {
		temp := arr[0]
		arr[0] = arr[i]
		arr[i] = temp
		adjuest(arr, i, 0)
	}
}
func CountSort(arr []float64) {
	length := len(arr)
	max, min := MaxAndMin(arr)
	arrLen := int(max - min + 1)
	countArr := make([]float64, arrLen)
	for i := 0; i < length; i++ {
		countArr[int(arr[i]-min)]++
	}
	index := 0
	for i := 0; i < arrLen; i++ {
		for j := countArr[i]; j > 0; j-- {
			arr[index] = float64(i) + min
			index++
		}
	}
}
func BucketSort(arr []float64, bktNum int) {
	length := len(arr)
	bucket_nums := bktNum + 1
	output := make([]float64, length)
	buckets := make([]int, bucket_nums)
	buckets2 := make([]int, bucket_nums)
	_, min := MaxAndMin(arr)
	d := 1.0/float64(bktNum)

	for i := 0; i < length; i++ {
		index :=int( (arr[i] - min) / d)
		if index <= bktNum {
			buckets[index]++
		} else {
			buckets[bucket_nums-1]++
		}
	}
	for i := 0; i < bucket_nums; i++ {
		buckets2[i] += buckets[i]
	}
	for i := 1; i < bucket_nums; i++ {
		buckets[i] += buckets[i-1]
	}
	for i := length - 1; i >= 0; i-- {
		index := int((arr[i] - min) / d)
		if index <= bktNum {
			output[buckets[index]-1] = arr[i]
			buckets[index]--
		} else {
			output[buckets[bucket_nums-1]-1] = arr[i]
			buckets[bucket_nums-1]--
		}
	}
	left, right := 0, 0
	for i := 0; i < bucket_nums; i++ {
		right += buckets2[i]
		InsertSort(output[left:right])
		left += buckets2[i]
	}
	for i := 0; i < length; i++ {
		arr[i] = output[i]
	}
}
func base_sort(arr []float64, exp int) {
	length := len(arr)
	output := make([]float64, length)
	buckets := make([]int, 10)
	for i := 0; i < length; i++ {
		buckets[int((arr[i]*math.Pow(10,16))/float64(exp))%10]++
	}
	for i := 1; i < 10; i++ {
		buckets[i] += buckets[i-1]
	}
	for i := length - 1; i >= 0; i-- {
		index := int((arr[i]*math.Pow(10,16))/float64(exp))%10
		output[buckets[index]-1] = arr[i]
		buckets[index]--
	}
	for i := 0; i < length; i++ {
		arr[i] = output[i]
	}
}
func RadixSort(arr []float64) {
	max, _ := MaxAndMin(arr)
	for i := 1; max*math.Pow(10,16)/float64(i) >= 1; i *= 10 {
		base_sort(arr, i)
	}
}



func Min2Exp(num int) int {
	exp := int(float64(math.Log2(float64(num))))
	dif := float64(num) - math.Exp2(float64(exp))
	if dif == 0 {
		return exp
	} else {
		return exp + 1
	}
}

func calcMergeTimes(num int)int{
	x:=math.Log2(float64(num))
	y:=math.Floor(x)
	if x-y==0{
		return int(x)
	}else {
		return int(y)+1
	}
}

func workerMerge(id int, jobs <-chan [3]int, arr []float64,wg *sync.WaitGroup) {
	defer wg.Done()
	indexs := <-jobs
	if !(indexs[0]==indexs[2]+1){
		Merge(arr, indexs[0], indexs[1], indexs[2])
	}
}
func workerMerge2(indexs [3]int, arr []float64,wg *sync.WaitGroup) {
	defer wg.Done()
	if !(indexs[0]==indexs[2]+1){
		Merge(arr, indexs[0], indexs[1], indexs[2])
	}
}




//close(jobs1)
//for a := 1; a <= cpuNums; a++ {
//	<-results
//}

