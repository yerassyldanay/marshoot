package randomer

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPayload(t *testing.T) {
	cities := []string{
		"костанай", "туркестан", "шымкент", "актобе", "караганда", "кызылорда", "петропавловск",
		"тараз", "уральск", "атырау", "кокшетау", "нур-султан", "талдыкорган", "актау",
		"алматы", "павлодар", "усть-каменогорск",
	}

	city := RandomOriginCity(cities...)
	require.NotZero(t, city)

	newCities := RandomCitiesToVisit(10, city, cities...)
	require.Equal(t, 10, len(newCities))
	fmt.Println(city)
	fmt.Println(newCities)
}
