package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

type kelvinFlag struct {
	Kelvin
}

func KelvinFlag(name string, value Kelvin, usage string) *Kelvin {
	f := kelvinFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Kelvin
}

func (f *kelvinFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Kelvin = CToK(Celsius(value))
		return nil
	case "F", "°F":
		f.Kelvin = FToK(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
