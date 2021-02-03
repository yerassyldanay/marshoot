package service

import (
	"errors"
	"fmt"
	"marshoot/pb"
	"math"
)

/*
	CalculatorCustom:
	this is my own solution for finding the optimal route
 */
type CalculatorCustom struct {
	DistanceMap map[string]map[string]int32
}

/*
	NewCalculatorCustom:
	creating a new calculator
 */
func NewCalculatorCustom(distanceMap map[string]map[string]int32) CalculatorCustom {
	return CalculatorCustom{
		DistanceMap:   distanceMap,
	}
}

/*
	GetOptimalPath:
 */
func (cc *CalculatorCustom) GetOptimalPath(origin string, citiesToVisit []string) (pb.CalculationResponseBody, error) {
	if len(citiesToVisit) == 0 {
		return pb.CalculationResponseBody{
			Error:         "destination is not specified",
		}, errors.New("destination is not specified")
	}

	for _, city := range citiesToVisit {
		_, exists := cc.DistanceMap[city]
		if !exists {
			return pb.CalculationResponseBody{
				Error: "there is no such city as " + city,
			}, errors.New("there is no such city as " + city)
		}
	}

	var newRoute = calculatorCustomCase{
		DistanceMap: &cc.DistanceMap,
	}
	newRoute.SetDefaultValues()

	newRoute.Seen[origin] = true
	newRoute.FindOptimalPath(origin, []string{}, citiesToVisit, 0, 0)

	//fmt.Println("newRoute.Cities:", newRoute.Cities)

	return pb.CalculationResponseBody{
		Cities: newRoute.Cities,
		TotalDistance: newRoute.TotalDistance,
	}, nil
}

/*
	calculatorCustomCase:
	custom calculator is established once when the server starts
	then all requests are handled by this calculatorCustomCase struct
 */
type calculatorCustomCase struct {
	DistanceMap *map[string]map[string]int32
	TotalDistance int32
	Cities []string
	Seen map[string]bool
}

func (newRoute *calculatorCustomCase) SetDefaultValues() {
	newRoute.TotalDistance = int32(math.Pow(2, 20))
	newRoute.Seen = map[string]bool{}
}

func (newRoute *calculatorCustomCase) FindOptimalPath(cityAt string, cityOrder []string, citiesToVisit []string, distance int32, numVisitedCities int) {
	if numVisitedCities == len(citiesToVisit) {
		if newRoute.TotalDistance > distance {
			newRoute.TotalDistance = distance
			newRoute.Cities = cityOrder
			fmt.Println(cityOrder, distance)
		}

		return
	}

	for _, city := range citiesToVisit {
		if was, _ := newRoute.Seen[city]; was {
			continue
		}

		distanceNext := (*newRoute.DistanceMap)[cityAt][city] + distance
		if distanceNext > newRoute.TotalDistance {
			continue
		}

		newRoute.Seen[city] = true
		newRoute.FindOptimalPath(city, append(cityOrder, city), citiesToVisit, distanceNext, numVisitedCities + 1)
		newRoute.Seen[city] = false
	}

	return
}
