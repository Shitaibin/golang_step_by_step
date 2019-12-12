package sort

import (
	"math/rand"
	"reflect"
	"testing"
)

// conclusion:
// insertion sort fast ~= slow
// shell sort > insertion sort

func insertionSortFast(nums []int, a, b int) {
	var tmp, j int
	for i := a + 1; i < b; i++ {
		tmp = nums[i]
		j = i - 1
		for ; j >= a && nums[j] > tmp; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = tmp
	}
}

func insertionSortSlow(nums []int, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && nums[j-1] > nums[j]; j-- {
			nums[j-1], nums[j] = nums[j], nums[j-1]
		}
	}
}

func shellSort(nums []int, a, b int) {
	for i := a + 6; i < b; i++ {
		if nums[i] < nums[i-6] {
			nums[i], nums[i-6] = nums[i-6], nums[i]
		}
	}
	insertionSortSlow(nums, a, b)
}

// sort range [a,b]
func quickSortFast(nums []int, a, b int) {
	for a < b {
		pivot := doPivot(nums, a, b)
		if pivot-1 > a {
			quickSortFast(nums, a, pivot-1)
		}

		// 减少递归
		a = pivot + 1
	}
}

func quickSortSlow(nums []int, a, b int) {
	if a >= b {
		return
	}

	pivot := doPivot(nums, a, b)
	quickSortSlow(nums, a, pivot-1)
	quickSortSlow(nums, pivot+1, b)
}

// doPivot returns the index of pivot
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

func TestInsertionSortSlow(t *testing.T) {
	nums := []int{5, 3, 7, 1, 9, 3, 4, 6, 2, 8}
	exp := []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9}
	insertionSortSlow(nums, 0, len(nums))
	if !reflect.DeepEqual(nums, exp) {
		t.Errorf("got: %v", nums)
	}
}

func TestInsertionSortFast(t *testing.T) {
	nums := []int{5, 3, 7, 1, 9, 3, 4, 6, 2, 8}
	exp := []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9}
	insertionSortFast(nums, 0, len(nums))
	if !reflect.DeepEqual(nums, exp) {
		t.Errorf("got: %v", nums)
	}
}

func TestShellSort(t *testing.T) {
	nums := []int{5, 3, 7, 1, 9, 3, 4, 6, 2, 8}
	exp := []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9}
	shellSort(nums, 0, len(nums))
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

func BenchmarkInsertionSortSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{5, 3, 7, 1, 9, 3, 4, 6, 2, 8}
		insertionSortSlow(nums, 0, len(nums))
	}
}

func BenchmarkInsertionSortFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{5, 3, 7, 1, 9, 3, 4, 6, 2, 8}
		insertionSortFast(nums, 0, len(nums))
	}
}

func BenchmarkShellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{5, 3, 7, 1, 9, 3, 4, 6, 2, 8}
		shellSort(nums, 0, len(nums))
	}
}

const N = 100000

var numsQuickFast []int
var numsQuickSlow []int

func init() {
	nums := make([]int, N, N)
	for i := 0; i < N; i++ {
		nums[i] = rand.Int()
	}

	numsQuickFast = make([]int, N, N)
	copy(numsQuickFast, nums)
	numsQuickSlow = make([]int, N, N)
	copy(numsQuickSlow, nums)

	print("Finish prepare data, with N=", N, " numbers\n")
}

func TestQuickSort(t *testing.T) {
	t.Logf("fast:\n%v", numsQuickFast)
	quickSortFast(numsQuickFast, 0, len(numsQuickFast)-1)
	quickSortSlow(numsQuickSlow, 0, len(numsQuickSlow)-1)
	if !reflect.DeepEqual(numsQuickFast, numsQuickSlow) {
		t.Errorf("fast:\n%v,\nslow:\n%v", numsQuickFast, numsQuickSlow)
	}
}

func BenchmarkQuickSortFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		quickSortFast(numsQuickFast, 0, len(numsQuickFast)-1)
	}
}

func BenchmarkQuickSortSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		quickSortSlow(numsQuickSlow, 0, len(numsQuickSlow)-1)
	}
}
