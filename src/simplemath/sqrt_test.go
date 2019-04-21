/*
* @Author: shishao
* @Date:   2019-04-21 21:42:58
* @Last Modified by:   shishao
* @Last Modified time: 2019-04-21 21:44:30
 */

package simplemath

import "testing"

func TestSqrt1(t *testing.T) {
	v := Sqrt(16)
	if v != 4 {
		t.Errorf("Sqrt(16) failed. Got %v, expected 4.", v)
	}
}
