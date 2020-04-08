package main

import (
	"fmt"
	"math"
)

type Converter struct{}

// Centimeter - Feet conversion
type Centimeter float64
type Feet float64

func (cvr Converter) CentimeterToFeet(c Centimeter) Feet {
	// conversion to feet
	var res = Feet(c / 30.48)
	return res
}
func (cvr Converter) FeetToCentimeter(f Feet) Centimeter {
	// conversion to centimeter
	var res = Centimeter(f * 30.48)
	return res
}

// Minute - Second conversion
type Minute float64
type Second float64

func (cvr Converter) SecondToMinute(s Second) Minute {
	// conversion to minute
	var res = Minute(s / 60)
	return res
}
func (cvr Converter) MinuteToSecond(m Minute) Second {
	// conversion to seconds
	var res = Second(m * 60)
	return res
}

// Seconds - Milliseconds
type Millisecond float64
type Seconds float64

func (cvr Converter) SecondToMilli(sec Seconds) Millisecond {
	// conversion to millisecond
	var res = Millisecond(sec * 1000)
	return res
}
func (cvr Converter) MilliToSecond(mil Millisecond) Seconds {
	// conversion to seconds
	var res = Seconds(mil / 1000)
	return res
}

// Celsius - Fahrenheit
type Celsius float64
type Fahrenheit float64

func (cvr Converter) CelsiusToFahrenheit(cel Celsius) Fahrenheit {
	// conversion to Fahrenheit
	var res = Fahrenheit((cel * 1.8) + 32)
	return res
}
func (cvr Converter) FahrenheitToCelsius(fa Fahrenheit) Celsius {
	// conversion to Celsius
	var res = Celsius((fa - 32) / 1.8)
	return res
}

// Radians - Degree
type Radian float64
type Degree float64

func (cvr Converter) RadianToDegree(rd Radian) Degree {
	// conversion to Degree
	var res = Degree(rd * (180 / math.Pi))
	return res
}
func (cvr Converter) DegreeToRadian(dg Degree) Radian {
	// conversion to Radian
	var res = Radian(dg * (math.Pi / 180))
	return res
}

// Kilogram - Pound
type Kilogram float64
type Pound float64

func (cvr Converter) KilogramToPound(kg Kilogram) Pound {
	// conversion to Pound
	var res = Pound(kg * 2.20462)
	return res
}
func (cvr Converter) PoundToKilogram(pd Pound) Kilogram {
	// conversion to Kilogram
	var res = Kilogram(pd / 2.20462)
	return res
}

// Kilobyte - Megabyte
type Kilobyte float64
type Megabyte float64

func (cvr Converter) KilobyteToMegabyte(kb Kilobyte) Megabyte {
	// conversion to Megabyte
	var res = Megabyte(kb / 1024)
	return res
}
func (cvr Converter) MegabyteToKilobyte(mb Megabyte) Kilobyte {
	// conversion to Kilobyte
	var res = Kilobyte(mb * 1024)
	return res
}

func main() {
	cvr := Converter{}
	// call the method and printing the result at the same time
	fmt.Println(cvr.CentimeterToFeet(65))
	fmt.Println(cvr.FeetToCentimeter(65))
	fmt.Println(cvr.SecondToMinute(300))
	fmt.Println(cvr.MinuteToSecond(7))
	fmt.Println(cvr.SecondToMilli(8))
	fmt.Println(cvr.MilliToSecond(45000))
	fmt.Println(cvr.CelsiusToFahrenheit(7))
	fmt.Println(cvr.FahrenheitToCelsius(100))
	fmt.Println(cvr.RadianToDegree(2))
	fmt.Println(cvr.DegreeToRadian(270))
	fmt.Println(cvr.KilogramToPound(5))
	fmt.Println(cvr.PoundToKilogram(16))
	fmt.Println(cvr.KilobyteToMegabyte(51200))
	fmt.Println(cvr.MegabyteToKilobyte(512))

}
