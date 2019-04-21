/*
* @Author: shishao
* @Date:   2019-03-04 19:48:41
* @Last Modified by:   shishao
* @Last Modified time: 2019-04-21 22:17:05
 */

package qsort

import "testing"

func TestQuickSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	QuickSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("Q;uickSort() failed. Got", values, "Expected 1 2 3 4 5 ")
	}
}

func TestQuickSort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	QuickSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 || values[4] != 5 {
		t.Error("Q;uickSort() failed. Got", values, "Expected 1 2 3 5 5 ")
	}
}

func TestQuickSort3(t *testing.T) {
	values := []int{5}
	QuickSort(values)
	if values[0] != 5 {
		t.Error("Q;uickSort() failed. Got", values, "Expected 5 ")
	}
}
