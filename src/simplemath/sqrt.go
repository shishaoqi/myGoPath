/*
* @Author: shishao
* @Date:   2019-04-21 21:41:08
* @Last Modified by:   shishao
* @Last Modified time: 2019-04-22 20:37:19
 */

package simplemath

import "math"

func Sqrt(i int) int {
	v := math.Sqrt(float64(i))
	return int(v)
}
