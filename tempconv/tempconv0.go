//указываем наименование нового пакета
package tempconv

import(
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Calvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func (c Celsius) String() string { 
	return fmt.Sprintf("%g°C", c) 
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°C", f)
} 
func (k Calvin) String() string {
	return fmt.Sprintf("%g°C", k)
}