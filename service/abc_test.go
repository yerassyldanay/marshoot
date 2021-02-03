package service

import (
	"log"
	"marshoot/utils"
	"testing"
)

var testCalculationServer *CalculationServer

func TestMain(m *testing.M) {
	pathToFile := "../data/distances.json"
	distanceMap, err := utils.LoadJsonFile(pathToFile)
	if err != nil {
		log.Fatal("could not load distance map. err:", err)
		return
	}

	calculatorCustom := NewCalculatorCustom(distanceMap)
	t := NewCalculationServer(&calculatorCustom)
	testCalculationServer = &t

	m.Run()
}

