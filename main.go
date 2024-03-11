package main

import (
	"fmt"
	"math"
)

type courier struct {
	Longitude float64
	Latitude  float64
	Order     *order
}

type order struct {
	Longitude float64
	Latitude  float64
	Cost      int64
}

func main() {
	couriers := []courier{}
	orders := []order{}

	feature(couriers, orders)
}

func feature(couriers []courier, orders []order) {
	for _, o := range orders {
		var closestCourier *courier
		minDistance := math.MaxFloat64

		for i, c := range couriers {
			if c.Order != nil {
				continue
			}

			dist := distance(o, c)
			if dist < minDistance {
				minDistance = dist
				closestCourier = &couriers[i]
			}
		}

		if closestCourier != nil {
			closestCourier.Order = &o
			fmt.Printf("Order with cost %d assigned to courier at location (%.4f, %.4f)\n", o.Cost, closestCourier.Latitude, closestCourier.Longitude)
		}
	}
}

func distance(o order, c courier) float64 {
	return math.Sqrt(math.Pow(o.Latitude-c.Latitude, 2) + math.Pow(o.Longitude-c.Longitude, 2))
}
