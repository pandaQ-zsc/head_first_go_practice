package main

import (
	"errors"
	"fmt"
	"log"
)

type Liters float64
type Gallons float64
type Milliliters float64

func (l Liters) toGallons() Gallons {
	return Gallons(l * 0.264)
}
func (m Milliliters) toGallons() Gallons {
	return Gallons(m * 0.000264)
}
func (l Liters) toMillisLiters() float64 {
	return float64(l * 1000)
}

func (g Gallons) toLiters() Liters {
	return Liters(g * 3.785)
}

type Date struct {
	Year  int
	Month int
	Day   int
}

type Event struct {
	Title string
	Date
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year ")
	}
	d.Year = year
	return nil
}

func main() {
	date := Date{
		Year:  2018,
		Month: 12,
		Day:   20,
	}
	//err := date.SetYear(2022)
	//if err != nil {
	//	log.Fatal(err)
	//}
	err := date.SetYear(0)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(date)

	soda := Liters(2)
	fmt.Printf("%f liters is %f gallons\n", soda, soda.toGallons())
	water := Milliliters(500)
	fmt.Printf("%f milliliters is %f gallons\n", water, water.toGallons())

}
