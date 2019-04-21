/*
* @Author: shishao
* @Date:   2019-04-21 21:39:22
* @Last Modified by:   shishao
* @Last Modified time: 2019-04-21 21:40:53
 */

package simplemath

import "testing"

func TestAdd1(t *testing.T) {
	r := Add(1, 2)
	if r != 3 {
		t.Errorf("Add(1, 2) failed. Got %d, expected 3.", r)
	}
}
