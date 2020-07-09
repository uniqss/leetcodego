package main

import (
	"fmt"
	"math"
)

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fret := findMedianSortedArrays(nums1, nums2)
	fmt.Println(fret)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	i := 0
	j := 0
	begin := 0
	end := m
	ivpre := 0
	iv := 0
	jvpre := 0
	jv := 0
	var leftmax float64
	var rightmin float64
	for begin <= end {
		i = (begin + end) / 2
		j = (m+n+1)/2 - i
		if i > 0 {
			ivpre = nums1[i-1]
		} else {
			ivpre = math.MinInt32
		}
		if j > 0 {
			jvpre = nums2[j-1]
		} else {
			jvpre = math.MinInt32
		}
		if i < m {
			iv = nums1[i]
		} else {
			iv = math.MaxInt32
		}
		if j < n {
			jv = nums2[j]
		} else {
			jv = math.MaxInt32
		}
		if ivpre < jv {
			leftmax = math.Max(float64(ivpre), float64(jvpre))
			rightmin = math.Min(float64(iv), float64(jv))
			begin = i + 1
		} else {
			end = i - 1
		}
	}
	if (m + n) % 2 == 0 {
		return (leftmax + rightmin) / 2
	} else {
		return leftmax
	}
}
