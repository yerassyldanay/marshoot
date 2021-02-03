package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type City struct {
	Name string `json:"name"`
	QueryName string `json:"query_name"`
	QueryId string `json:"query_id"`
} 

var cities = []City{
	{
		Name: 			"Шымкент",
		QueryName:      "Шымкент, Туркестанская область, Казахстан",
		QueryId:       	"38909",
	},
	{
		Name:      "Алматы",
		QueryName: "Алматы, Алматинская область, Казахстан",
		QueryId:   "127269",
	},
	{
		Name:      "Тараз",
		QueryName: "Тараз, Жамбылская область, Казахстан",
		QueryId:   "28090",
	},
	{
		Name:      "Кызылорда",
		QueryName: "Кызылорда, Кызылординская область, Казахстан",
		QueryId:   "22243",
	},
	{
		Name:      "Нур-Султан",
		QueryName: "Нур-Султан, Акмолинская область, Казахстан",
		QueryId:   "101",
	},
	{
		Name:      "Кокшетау",
		QueryName: "Кокшетау, Акмолинская область, Казахстан",
		QueryId:   "14726",
	},
	{
		Name:      "Актобе",
		QueryName: "Актобе, Актюбинская область, Казахстан",
		QueryId:   "19590",
	},
	{
		Name:      "Талдыкорган",
		QueryName: "Талдыкорган, Алматинская область, Казахстан",
		QueryId:   "137093",
	},
	{
		Name:      "Атырау",
		QueryName: "Атырау, Атырауская область, Казахстан",
		QueryId:   "14773",
	},
	{
		Name:      "Уральск",
		QueryName: "Уральск, Западно-Казахстанская область, Казахстан",
		QueryId:   "21306",
	},
	{
		Name:      "Караганда",
		QueryName: "Караганда, Карагандинская область, Казахстан",
		QueryId:   "1236",
	},
	{
		Name:      "Костанай",
		QueryName: "Костанай, Костанайская область, Казахстан",
		QueryId:   "1237",
	},
	{
		Name:      "Актау",
		QueryName: "Актау, Мангистауская область, Казахстан",
		QueryId:   "21859",
	},
	{
		Name:      "Павлодар",
		QueryName: "Павлодар, Павлодарская область, Казахстан",
		QueryId:   "1238",
	},
	{
		Name:      "Петропавловск",
		QueryName: "Петропавловск, Северо-Казахстанская область, Казахстан",
		QueryId:   "25452",
	},
	{
		Name:      "Туркестан",
		QueryName: "Туркестан, Туркестанская область, Казахстан",
		QueryId:   "30167",
	},
	{
		Name:      "Усть-Каменогорск",
		QueryName: "Оскемен(Усть-Каменогорск) (Усть-Каменогорск), Восточно-Казахстанская область, Казахстан",
		QueryId:   "19603",
	},
	//{
	//	Name:      "Петропавловск",
	//	QueryName: "Петропавловск, Северо-Казахстанская область, Казахстан",
	//	QueryId:   "25452",
	//},
}

type ResponseParseCityInfoRouteInfo struct {
	Distance string `json:"distance"`
	Duration string `json:"duration"`
}

type ResponseParseCityInfo struct {
	CitiesAndAbbrs []string                    `json:"cities_and_abbrs"`
	RouteInfo []ResponseParseCityInfoRouteInfo `json:"route_info"`
}

var distances = map[string]map[string]int{}

func ParseCityInfo() {
	resourceUrl := "https://www.della.kz/ajax_request.php"

	for i := 0; i < len(cities); i++ {
		for j := i + 1; j < len(cities); j++ {

			u := url.URL{}
			q := u.Query()

			fromCity := cities[i]
			toCity := cities[j]

			q.Set("mode", "calculate_distance")
			q.Set("waypoints[0][]", "A")
			q.Add("waypoints[0][]", fromCity.QueryName)
			q.Add("waypoints[0][]", fromCity.QueryName)
			q.Set("waypoints[1][]", "B")
			q.Add("waypoints[1][]", toCity.QueryName)
			q.Add("waypoints[1][]", toCity.QueryId)

			var newUrl = resourceUrl + "?" + q.Encode()

			//fmt.Println(newUrl)
			//fmt.Println(len(cities))
			fmt.Println(fromCity.Name, toCity.Name)

			cl := http.DefaultClient

			// request
			req, err := http.NewRequest(http.MethodGet, newUrl, nil)
			if err != nil {
				fmt.Println("could not create a request, err: " + err.Error())
				fmt.Println("problem with ", fromCity.Name, toCity.Name)
				continue
			}

			// make one
			resp, err := cl.Do(req)
			if err != nil {
				fmt.Println("could not make a request, err: " + err.Error())
				fmt.Println("problem with ", fromCity.Name, toCity.Name)
				continue
			}
			defer resp.Body.Close()

			//bodyBytes, err := ioutil.ReadAll(resp.Body)
			//if err != nil {
			//	log.Fatal(err)
			//}
			//bodyString := string(bodyBytes)
			//log.Println("bodyString:", bodyString)
			//fmt.Println(resp.Body)

			// parse
			var respBody = ResponseParseCityInfo{}
			err = json.NewDecoder(resp.Body).Decode(&respBody)
			if err != nil {
				fmt.Println("could not parse body, err: " + err.Error())
				fmt.Println("problem with ", fromCity.Name, toCity.Name)
				continue
			}

			//fmt.Printf("%#v \n", respBody)

			time.Sleep(time.Millisecond * 10)

			key1 := strings.ToLower(fromCity.Name)
			key2 := strings.ToLower(toCity.Name)

			_, ok := distances[key1]
			if !ok {
				distances[key1] = map[string]int{}
			}

			_, ok = distances[key2]
			if !ok {
				distances[key2] = map[string]int{}
			}

			if len(respBody.RouteInfo) == 0 {
				fmt.Println("problem with ", fromCity.Name, toCity.Name)
				continue
			}

			routeInfo := respBody.RouteInfo[0]
			dis, err := strconv.Atoi(routeInfo.Distance)
			if err != nil {
				fmt.Println(err)
			}
			distances[key1][key2] = dis
			distances[key2][key1] = dis
		}

		//return
	}
}

func TestJson() {
	distancesString, err := json.MarshalIndent(distances, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("distances.json", distancesString, os.ModePerm)

	distancesString2, err := json.Marshal(distances)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("distances_raw.json", distancesString2, os.ModePerm)
}


/*
	{"rows":[{"elements":[{"distance":{"value":1076},"duration":{"value":774},"status":"OK"}]},{"elements":[{"distance":{"value":1503},"duration":{"value":1082},"status":"OK"}]}]}
 */

// 42.3417,69.5901 Shymkent
// 43.2220,76.8512
// 42.8984,71.3980

