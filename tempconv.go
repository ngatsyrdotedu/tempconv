package tempconv

import (
    "fmt"
)

type Celsius float64
type Fahrenheit float64

const AbsoluteZeroC   Celsius = -273.15
const FreezingC       Celsius = 0
const BoilingC        Celsius = 100

func (c Celsius) String() string {return fmt.Sprintf("%gºC", c)}
func (f Fahrenheit) String() string {return fmt.Sprintf("%gºF", f)}
