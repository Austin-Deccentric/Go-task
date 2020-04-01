package main
import (
	"fmt"
	"math"
)
// Converter converts between different SI units
type Converter struct{}

//Feet is a unit of Length  
type Feet float64
//Centimeter is a unit of length
type Centimeter float64

//FeetToCentimeter converts From Feets to Centimeters
func (cvr Converter) FeetToCentimeter(a Feet) Centimeter{
	ans:= a * 30.48
	return Centimeter(ans)
}
//CentimeterToFeet converts From Centimeters Feets to
func (cvr Converter) CentimeterToFeet(a Centimeter) Feet{
	return Feet(a/30.48)
}

//Minute is a unit of Time
type Minute float64
//Second is a unit of Time
type Second float64

//MinuteToSecond converts from minutes to seconds.
func (cvr Converter) MinuteToSecond(t Minute) Second{
	return Second(t*60)
}

//SecondToMinute converts from Seconds to minutes
func (cvr Converter) SecondToMinute(t Second) Minute{
	return Minute(t/60)
}

//Millisecond is a unit of time
type Millisecond float64

//SecondtoMillisecond converts from second to millisecond
func (cvr Converter) SecondtoMillisecond(t Second) Millisecond{
	return Millisecond(t*1000)
}

//MillisecondtoSecond converts from millisecond to second
func (cvr Converter) MillisecondtoSecond(t Millisecond) Second{
	return Second(t/1000)
}

//Celsius is a unit of Temperature
type Celsius float64
//Fahrenheit is a unit of Tempreture
type Fahrenheit float64

//CelsiusToFahrenheit converts from Celsius to Fahrenheit
func (cvr Converter) CelsiusToFahrenheit(T Celsius) Fahrenheit{
	return Fahrenheit((T *9/5) + 32)
}

//FahrenheitToCelsius converts from Celsius to Fahrenheit
func (cvr Converter) FahrenheitToCelsius(T Fahrenheit) Celsius{
	return Celsius((T -32) * 5/9)
}

//Radian is a unit of Phase plane angle
type Radian float64
//Degree is a unit of Phase plane angle
type Degree float64

//RadianToDegree converts from radians to degree
func (cvr Converter) RadianToDegree(z Radian) Degree{
	return Degree((z*(180/math.Pi)))
}

//DegreeToRadian converts from degree to radian
func (cvr Converter) DegreeToRadian(z Degree) Radian{
	return Radian((z*(math.Pi/180)))
}

//Kilogramm is a unit of mass
type Kilogramm float64
//Pound is a unit of mass
type Pound float64

//KilogrammToPound converts from kilogramm to Pound
func (cvr Converter) KilogrammToPound(m Kilogramm) Pound{
	return Pound(m*2.20462)
} 

//PoundToKilogramm converts from Pounds to Kiogramms
func (cvr Converter) PoundToKilogramm(m Pound) Kilogramm{
	return Kilogramm(m/2.20462)
}

func main() {
	cvr := Converter{}
	fmt.Printf("The result is %.3v Feet\n",cvr.CentimeterToFeet(185))
	fmt.Printf("The result is %.3v Centimeter\n",cvr.FeetToCentimeter(5.11))
	fmt.Printf("The result is %.3v Seconds\n",cvr.MinuteToSecond(1))
	fmt.Printf("The result is %.3v Minutes\n",cvr.SecondToMinute(720))
	fmt.Printf("The result is %v Milliseconds\n",cvr.SecondtoMillisecond(1))
	fmt.Printf("The result is %.3v Seconds\n",cvr.MillisecondtoSecond(10000))
	fmt.Printf("The result is %.3v Fahrenheit\n",cvr.CelsiusToFahrenheit(38))
	fmt.Printf("The result is %.3v Celsius\n",cvr.FahrenheitToCelsius(0))
	fmt.Printf("The result is %.3v Degree\n",cvr.RadianToDegree(1))
	fmt.Printf("The result is %.3v rad\n",cvr.DegreeToRadian(180))
	fmt.Printf("The result is %.3v Pounds\n",cvr.KilogrammToPound(75))
	fmt.Printf("The result is %.3v Kilogramms\n",cvr.PoundToKilogramm(250))	
}