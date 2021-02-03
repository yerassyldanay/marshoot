package service

import "marshoot/pb"

/*
	ResponseGetOptimalPath:
	includes cities & total distance
 */
//type ResponseGetOptimalPath struct {
//	CitiesInOrder []string
//	TotalDistance int32
//}

/*
	Calculator:
	the optimal path can be calculated in various ways
	e.g. using our custom algorithm or using prepared tool such as Yandex Map or Google Map
 */
type Calculator interface {
	/*
		GetOptimalPath:
		this function accepts the city name (origin), where our journey begins
		& a list of cities, which we must visit
		returns: a list of cities ordered by the optimal path & the total distance
	 */
	GetOptimalPath(origin string, citiesToVisit []string) (resp pb.CalculationResponseBody, err error)
}
