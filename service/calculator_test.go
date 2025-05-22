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
	require.NoError(t, err)
	require.NotNil(t, out) // Important check before accessing out.Cities or out.TotalDistance
	require.Equal(t, 7, len(out.Cities))
	require.True(t, out.TotalDistance > 0, "TotalDistance should be positive") // More descriptive than NotZero
}

