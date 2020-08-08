package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	_len := len(nums1) + len(nums2)
	_sign := _len % 2
	_a := _len / 2
	_b := _a
	_addon := 0
	_cur := 0

	var a int
	var b int

	if _sign == 0 {
		_b -= 1
	}

	// find _a && _b
	for i := 0 ; i < len(nums2) && _cur + _addon <= _a; i ++ {
		t := nums2[i]
		_i := findIndex(nums1, t, _cur)

		if _addon + _i == _b {
			b = t
		}

		if _addon + _i == _a {
			a = t
			break
		}

		_cur = _i
		_addon += 1
	}



	return float64(a + b) / 2
}

func findIndex(arr []int, target int, start int) int {
	_len := len(arr)
	_low := start
	_high := _len - 1
	var _mid int

	for _low <= _high {
		_mid = (_low + _high) / 2
		if arr[_mid] > target {
			_high = _mid - 1
		} else if arr[_mid] < target {
			_low = _mid + 1
			if _low > _high {
				return _mid + 1
			}
		} else {
			return _mid + 1
		}
	}
	return _mid
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{4, 5}))
}
