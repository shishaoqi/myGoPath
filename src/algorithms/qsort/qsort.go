/*
* @Author: shishao
* @Date:   2019-03-04 17:44:23
* @Last Modified by:   shishao
* @Last Modified time: 2019-03-04 21:24:23
 */

package qsort

func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		if values[i] <= temp && i <= p {
			i++
		}

		if i <= p {
			values[p] = values[i]
			p = i
		}
	}

	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
