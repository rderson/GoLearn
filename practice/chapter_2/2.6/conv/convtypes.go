package conv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Meters float64
type Feet float64
type Pounds float64
type Kilos float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (m Meters) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string { return fmt.Sprintf("%gft", f) }
func (p Pounds) String() string { return fmt.Sprintf("%glbs", p) }
func (k Kilos) String() string { return fmt.Sprintf("%gkg", k) }

