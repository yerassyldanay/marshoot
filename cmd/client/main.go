package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"marshoot/pb"
	"marshoot/utils"
	"marshoot/utils/randomer"
	"path/filepath"
)

func main() {
	// provided args
	serverHost := flag.String("host", "0.0.0.0", "the server host (for client)")
	serverPort := flag.String("port", "9000", "the server port (for client)")
	flag.Parse()

	address := fmt.Sprintf("%s:%s", *serverHost, *serverPort)

	// establishing a connection with a server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not dial with a server. err: ", err)
	}

	calcClient := pb.NewCalculationServiceClient(conn)

	/*
		get cities
	 */
	pathToFile := "data/distances.json"
	pathToFile, err = filepath.Abs(pathToFile)
	if err != nil {
		log.Fatal("problem with a file path. err: ", err)
	}

	distanceMap, err := utils.LoadJsonFile(pathToFile)
	if err != nil {
		log.Fatal("could not load json file. err: ", err)
	}
	cities := []string{}
	for city, _ := range distanceMap {
		cities = append(cities, city)
	}

	// generating a random city (to start) and a list of random cities (to visit)
	length := len(cities)
	cityAt := randomer.RandomOriginCity(cities...)
	newCities := randomer.RandomCitiesToVisit(randomer.RandomInt(1, length), cityAt, cities...)

	// request
	resp, err := calcClient.CalculateOptimalPath(context.Background(), &pb.CalculationRequestBody{
		Origin: cityAt,
		Cities: newCities,
	})
	if err != nil {
		log.Fatal("could not send a request to the server. err: ", err)
	}

	// response
	fmt.Println("Please, visit in the next order: ", resp.Cities)
	fmt.Println("The total optimal distance will be: ", resp.TotalDistance)
}

