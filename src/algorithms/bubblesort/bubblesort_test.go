/*
* @Author: shishao
* @Date:   2019-03-04 17:28:48
* @Last Modified by:   shishao
* @Last Modified time: 2019-03-04 17:35:32
 */

package bubblesort

import "testing"

func TestBubbleSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("BubbleSort() failed. Got", values, "Expected 1 2 3 4 5")
	}
}

func TestBubbleSort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 || values[4] != 5 {
		t.Error("BubbleSort() failed. Got", values, "Expected 1 2 3 5 5")
	}
}

func TestBubbleSort3(t *testing.T) {
	values := []int{5}
	BubbleSort(values)
	if values[0] != 5 {
		t.Error("BubbleSort() failed. Got", values, "Expected 5")
	}
}
