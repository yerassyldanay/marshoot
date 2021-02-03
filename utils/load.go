package utils

import (
	"encoding/json"
	"io/ioutil"
)

func LoadJsonFile(pathToFile string) (map[string]map[string]int32, error) {
	//fmt.Println(pathToFile)

	byteFileData, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		return map[string]map[string]int32{}, err
	}

	var distanceMap = map[string]map[string]int32{}
	if err = json.Unmarshal(byteFileData, &distanceMap); err != nil {
		return map[string]map[string]int32{}, err
	}

	//fmt.Printf("%#v \n", distanceMap)
	//fmt.Println(distances["алматы"]["актобе"])

	return distanceMap, nil
}


