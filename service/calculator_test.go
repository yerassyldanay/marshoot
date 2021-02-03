package service

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"

	//"github.com/stretchr/testify/require"
	"marshoot/pb"
	"testing"
)

func TestCalculatorCustom(t *testing.T) {
	in := pb.CalculationRequestBody{
		Origin: "алматы",
		Cities: []string{
			"кызылорда",
			"тараз",
			"шымкент",
			"актау",
			"уральск",
			//"актобе",
			"атырау",
			//"костанай",
			"караганда",
		},
	}
	out, err := testCalculationServer.CalculateOptimalPath(context.Background(), &in)
	//require.NoError(t, err)
	require.NotZero(t, len(out.Cities))
	require.NotZero(t, out.TotalDistance)
	fmt.Println(err)

	fmt.Println(out)
	//fmt.Printf("%#v \n", out.Cities)
	//fmt.Printf("%#v \n", out.TotalDistance)
}

