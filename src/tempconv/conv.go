/*
* @Author: shishao
* @Date:   2019-02-14 13:39:13
* @Last Modified by:   shishao
* @Last Modified time: 2019-02-14 13:42:13
 */

package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
