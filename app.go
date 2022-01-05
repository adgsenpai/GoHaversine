package main

import (
	"fmt"
	"math"
	"bufio"
	"strconv"
	"os"
)

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

const radius = 6371 // Earth's mean radius in kilometers

func degrees2radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func (origin Coordinates) Distance(destination Coordinates) float64 {
	degreesLat := degrees2radians(destination.Latitude - origin.Latitude)
	degreesLong := degrees2radians(destination.Longitude - origin.Longitude)
	a := (math.Sin(degreesLat/2)*math.Sin(degreesLat/2) +
		math.Cos(degrees2radians(origin.Latitude))*
			math.Cos(degrees2radians(destination.Latitude))*math.Sin(degreesLong/2)*
			math.Sin(degreesLong/2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := radius * c

	return d
}

func main() {
	const bitSize = 64

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your latitude pointA: ")
	latitudeA, _ := reader.ReadString('\n')
	fmt.Print("Enter your longitude pointA: ")
	longitudeA, _ := reader.ReadString('\n')

	fmt.Print("Enter your latitude pointB: ")
	latitudeB, _ := reader.ReadString('\n')
	fmt.Print("Enter your longitude pointB: ")
	longitudeB, _ := reader.ReadString('\n')


	latA,err := strconv.ParseFloat(latitudeA[:len(latitudeA)-2],bitSize)
	latB,err := strconv.ParseFloat(latitudeB[:len(latitudeB)-2],bitSize)
	longA,err := strconv.ParseFloat(longitudeA[:len(longitudeA)-2],bitSize)
	longB,err := strconv.ParseFloat(longitudeB[:len(longitudeB)-2],bitSize)
	
	if err != nil { }


	pointA := Coordinates{latA, longA}
	pointB := Coordinates{latB, longB}

	fmt.Println("Point A : ", pointA)
	fmt.Println("Point B : ", pointB)

	distance := pointA.Distance(pointB)
	fmt.Printf("The distance from point A to point B is %.2f kilometers.\n", distance)
	reader.ReadString('\n')
}