package main

import "fmt"

func findPeak(A []int, c chan int, start, end int) {
	l := start
	r := end
	if (A[l] == A[r]) {
		c<-l
		return
	}
	
	retLeft := true
	for l <= r {
		mid := l + (r - l) / 2
		if mid == start {
			l++
		} else if mid == end {
			r--
		} else {
			if A[mid] > A[mid - 1] && A[mid] > A[mid + 1] {
				c<-mid
				return
			} else if A[mid] > A[mid - 1] || A[mid + 1] > A[mid] {
				l = mid + 1
				retLeft = false
			} else if A[mid] > A[mid + 1] || A[mid - 1] > A[mid] {
				r = mid - 1
			} else {
				r--
			}
		}
	}
	if retLeft {
		c<-start
	} else {
		c<-end
	}
}

func peakIndexInMountainArray(A []int, c chan int, lenPerWkr int) {
	start := 0
	n := len(A)
	for start < n {
		end := start + lenPerWkr - 1
		if end >= n {
			end = n - 1
		}
		go findPeak(A, c, start, end)
		start = end + 1
	}
}

func getPeakFromCh(A []int, ch chan int) int {
	maxNum := 0
	for idx := range ch {
		if A[idx] > maxNum {
			maxNum = A[idx]
		}
	}
	return maxNum
}

func main() {
	var A = []int{1,2,3,4,3,2,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	n := len(A)
	lenPerWkr := 4
	if n % lenPerWkr > 0 {
		n = n / lenPerWkr + 1
	} else {
		n = n / lenPerWkr
	}
	ch := make(chan int, n)
	peakIndexInMountainArray(A, ch, lenPerWkr)
	for len(ch) < n {}
	close(ch)
	
	peak := getPeakFromCh(A, ch)
	fmt.Println(peak)
}

