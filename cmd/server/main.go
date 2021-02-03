package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"marshoot/pb"
	"marshoot/service"
	"marshoot/utils"
	"net"
	"path/filepath"
)

func main() {
	// provided args
	host := flag.String("host", "0.0.0.0", "the client host")
	port := flag.String("port", "9000", "the client port")
	flag.Parse()

	fmt.Println("Client is running on: ", *host + ":" + *port)

	// loading a json file (there is info on distance between cities)
	pathToFile := "data/distances.json"
	pathToFile, err := filepath.Abs(pathToFile)
	if err != nil {
		log.Fatal("problem with a file path. err: ", err)
	}

	distanceMap, err := utils.LoadJsonFile(pathToFile)
	if err != nil {
		log.Fatal("could not load json file. err: ", err)
	}

	calcCustom := service.NewCalculatorCustom(distanceMap)
	calcServer := service.NewCalculationServer(&calcCustom)

	grpcServer := grpc.NewServer()
	pb.RegisterCalculationServiceServer(grpcServer, &calcServer)

	address := fmt.Sprintf("%s:%s", *host, *port)
	fmt.Println("address:", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("could not create a listener. err: ", err)
	}

	// listing the requests
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("could not start a server. err: ", err)
	}
}

