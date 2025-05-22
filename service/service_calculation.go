package service

import (
	"context"
	"fmt"
	"marshoot/pb"
	"strings"
)

/*
	CalculationServer:
	you can implement own solution (for finding an optimal route)
	and inject it
 */
type CalculationServer struct {
	Calculator Calculator
}

func NewCalculationServer(calc Calculator) CalculationServer {
	return CalculationServer {
		Calculator: calc,
	}
}

func (cs *CalculationServer) CalculateOptimalPath(ctx context.Context, in *pb.CalculationRequestBody) (*pb.CalculationResponseBody, error) {
	/*
		check:
		gets rid of extra city names
	*/
	cityAt := strings.ToLower(in.Origin)
	cities := []string{}

	seen := map[string]bool{}
	seen[cityAt] = true

	for _, city := range in.Cities {
		tCity := strings.ToLower(city)
		_, ok := seen[tCity]
		if ok {
			continue
		}
		seen[tCity] = true
		cities = append(cities, tCity)
	}

	optimalPath, err := cs.Calculator.GetOptimalPath(cityAt, cities)
	fmt.Println("Received: ", in.Origin, in.Cities)
	return optimalPath, err
}
