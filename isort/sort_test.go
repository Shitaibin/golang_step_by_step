package isort

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

// conclusion:
// insertion sort fast ~= slow
// shell sort > insertion sort

func insertionSortFast(nums []int, a, b int) {
	var tmp, j int
	for i := a + 1; i <= b; i++ {
		tmp = nums[i]
		j = i - 1
		for ; j >= a && nums[j] > tmp; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = tmp
	}
}

func insertionSort(nums []int, a, b int) {
	for i := a + 1; i <= b; i++ {
		for j := i; j > a && nums[j-1] > nums[j]; j-- {
			nums[j-1], nums[j] = nums[j], nums[j-1]
		}
	}
}

func shellSort(nums []int, a, b int) {
	for i := a + 6; i <= b; i++ {
		if nums[i] < nums[i-6] {
			nums[i], nums[i-6] = nums[i-6], nums[i]
		}
	}
	insertionSort(nums, a, b)
}

// sort range [a,b]
// func quickSortFast(nums []int, a, b int) {
// 	for a < b {
// 		pivot := doPivot(nums, a, b)
// 		if pivot-1 > a {
// 			quickSortFast(nums, a, pivot-1)
// 		}

// 		// 减少递归
// 		a = pivot + 1
// 	}
// }

// sort range [a,b]
func quickSortFast(nums []int, a, b int) {
	if b-a > 12 {
		if a >= b {
			return
		}

		pivot := doPivot(nums, a, b)
		quickSortSlow(nums, a, pivot-1)
		quickSortSlow(nums, pivot+1, b)
	}

	for i := a + 6; i < b; i++ {
		if nums[i] < nums[i-6] {
			nums[i], nums[i-6] = nums[i-6], nums[i]
		}
	}
	insertionSort(nums, a, b)
}

func quickSortSlow(nums []int, a, b int) {
	if a >= b {
		return
	}

	pivot := doPivot(nums, a, b)
	quickSortSlow(nums, a, pivot-1)
	quickSortSlow(nums, pivot+1, b)
}

// doPivot returns the index of pivot`
func doPivot(nums []int, a, b int) int {
	pivot := a
	for a < b {
		for a < b && nums[b] >= nums[pivot] {
			b--
		}
		for a < b && nums[a] <= nums[pivot] {
			a++
		}
		nums[a], nums[b] = nums[b], nums[a]
	}

	// a is new pivot index
	nums[a], nums[pivot] = nums[pivot], nums[a]
	return a
}

type IntSlice []int

// func (i *IntSlice) Len() int {
// 	return len(*i)
// }

// func (i *IntSlice) Less(x, y int) bool {
// 	return (*i)[x] <= (*i)[y]
// }

// func (i *IntSlice) Swap(x, y int) {
// 	(*i)[x], (*i)[y] = (*i)[y], (*i)[x]
// }

func (i IntSlice) Len() int {
	return len(i)
}

func (i IntSlice) Less(x, y int) bool {
	return i[x] <= i[y]
}

func (i IntSlice) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func TestInsertionSort(t *testing.T) {
	nums := []int{5, 3, 7, 1, 9, 3, 4, 6, 2, 8}
	exp := []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9}
	insertionSort(nums, 0, len(nums)-1)
	if !reflect.DeepEqual(nums, exp) {
		t.Errorf("got: %v", nums)
	}
}

func TestInsertionSortFast(t *testing.T) {
	nums := []int{5, 3, 7, 1, 9, 3, 4, 6, 2, 8}
	exp := []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9}
	insertionSortFast(nums, 0, len(nums)-1)
	if !reflect.DeepEqual(nums, exp) {
		t.Errorf("got: %v", nums)
	}
}

func TestShellSort(t *testing.T) {
	nums := []int{5, 3, 7, 1, 9, 3, 4, 6, 2, 8}
	exp := []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9}
	shellSort(nums, 0, len(nums)-1)
	if !reflect.DeepEqual(nums, exp) {
		t.Errorf("got: %v", nums)
	}
}

func TestQuickSortSlow(t *testing.T) {
	nums := []int{5, 3, 7, 1, 9, 3, 4, 6, 2, 8}
	exp := []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9}
	quickSortSlow(nums, 0, len(nums)-1)
	if !reflect.DeepEqual(nums, exp) {
		t.Errorf("got: %v", nums)
	}
}

func TestQuickSortFast(t *testing.T) {
	nums := []int{5, 3, 7, 1, 9, 3, 4, 6, 2, 8}
	exp := []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9}
	quickSortFast(nums, 0, len(nums)-1)
	if !reflect.DeepEqual(nums, exp) {
		t.Errorf("got: %v", nums)
	}
}

func TestQuickSort(t *testing.T) {
	quickSortFast(numsQuickFast, 0, len(numsQuickFast)-1)
	quickSortSlow(numsQuickSlow, 0, len(numsQuickSlow)-1)
	if !reflect.DeepEqual(numsQuickFast, numsQuickSlow) {
		t.Errorf("fast:\n%v,\nslow:\n%v", numsQuickFast, numsQuickSlow)
	}
}

func TestQuickSortWithNNumbers(t *testing.T) {
	is := IntSlice(numsQuickFast)
	sort.Sort(&is)
	quickSortSlow(numsQuickSlow, 0, len(numsQuickSlow)-1)
	if !reflect.DeepEqual(numsQuickFast, numsQuickSlow) {
		t.Errorf("fast:\n%v,\nslow:\n%v", numsQuickFast, numsQuickSlow)
	}
}

func TestSortSort(t *testing.T) {
	nums := []int{5, 3, 7, 1, 9, 3, 4, 6, 2, 8}
	exp := []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9}
	is := IntSlice(nums)
	sort.Sort(&is)
	if !reflect.DeepEqual(nums, exp) {
		t.Errorf("got: %v", nums)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{5, 3, 7, 1, 9, 3, 4, 6, 2, 8}
		insertionSort(nums, 0, len(nums)-1)
	}
}

func BenchmarkInsertionSortFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{5, 3, 7, 1, 9, 3, 4, 6, 2, 8}
		insertionSortFast(nums, 0, len(nums)-1)
	}
}

func BenchmarkShellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{5, 3, 7, 1, 9, 3, 4, 6, 2, 8}
		shellSort(nums, 0, len(nums)-1)
	}
}

const N = 10000

var numsQuickFast []int
var numsQuickSlow []int

func init() {
	genData(1000)
	genData(10000)
	genData(100000)
	genData(1000000)
}

func genData(n int) {
	filename := genFileName(n)
	if _, err := os.Stat(filename); err == nil {
		println("File:", filename, "exists")
		return
	}

	nums := make([]int, n, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.Int()
	}

	numsQuickFast = make([]int, n, n)
	copy(numsQuickFast, nums)
	numsQuickSlow = make([]int, n, n)
	copy(numsQuickSlow, nums)

	print("Finish prepare data, with n=", n, " numbers\n")

	data, err := json.Marshal(nums)
	if err != nil {
		print("Marshal data error: ", err.Error(), "\n")
		return
	}
	f, err := os.Create(filename)
	if err != nil {
		print("Open file error: ", err.Error(), "\n")
		return
	}
	_, err = f.Write(data)
	if err != nil {
		print("Write data error: ", err.Error(), "\n")
		return
	}
	print("Finish wirte data\n")
}

func readData(n int) []int {
	filename := genFileName(n)
	var nums []int
	file, err := os.Open(filename)
	if err != nil {
		println("Open file error:", err.Error())
		return nil
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		println("Read file error:", err.Error())
		return nil
	}

	// var nums IntSlice
	if err = json.Unmarshal(data, &nums); err != nil {
		println("Unmarshal slice error:", err.Error())
	}
	return nums
}

func genFileName(n int) string {
	return "data_" + strconv.Itoa(n) + ".json"
}

func BenchmarkQuickSortFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := readData(N)
		b.StartTimer()
		quickSortFast(nums, 0, len(nums)-1)
	}
}

func BenchmarkQuickSortSlow(b *testing.B) {
	b.StopTimer()
	rawNums := readData(N)
	nums := make([]int, len(rawNums), len(rawNums))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(nums, rawNums)
		quickSortSlow(nums, 0, len(nums)-1)
	}
}

func BenchmarkSortSort(b *testing.B) {
	b.StopTimer()
	rawNums := readData(N)
	nums := make([]int, len(rawNums), len(rawNums))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		// b.StopTimer() timer will waster more time
		copy(nums, rawNums)
		is := IntSlice(nums)
		// b.StartTimer()
		sort.Sort(is)
	}
}

// 标准库优化的，虐死自己实现的，才1000个数
// BenchmarkQuickSortFast-8       	    3490	    333624 ns/op
// BenchmarkQuickSortSlow-8       	    3529	    331864 ns/op
// BenchmarkSortSort-8            	   30051	     40065 ns/op
// PASS
// ok  	_/Users/shitaibin/Workspace/golang_step_by_step/sort	8.196s
