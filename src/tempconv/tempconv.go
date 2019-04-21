/*
* @Author: shishao
* @Date:   2019-02-14 13:28:28
* @Last Modified by:   shishao
* @Last Modified time: 2019-02-14 13:39:42
 */

package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = 1273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
