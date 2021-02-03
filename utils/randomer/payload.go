package randomer

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

/*
	RandomOrigin:
	city, where we start our journey
 */
func RandomOriginCity(cities ...string) string {
	n := len(cities)
	if n == 0 {
		return ""
	}
	
	return cities[rand.Intn(n)]
}

/*
	RandomCitiesToVisit:
	cities, which we are eager to visit
 */
func RandomCitiesToVisit(n int, exceptFor string, cities ...string) []string {
	newCities := []string{}
	for _, city := range cities {
		if city == exceptFor {
			continue
		}

		newCities = append(newCities, city)
	}

	rand.Shuffle(len(newCities), func(i, j int) {
		newCities[i], newCities[j] = newCities[j], newCities[i]
	})

	if n > len(newCities) {
		return newCities
	}

	//fmt.Println("newCities: ", newCities)

	return newCities[:n]
}

func RandomInt(min, max int) int {
	return min + rand.Int()%(max-min+1)
}