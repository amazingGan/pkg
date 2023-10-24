package geo

import (
	g "github.com/golang/geo/s2"
)

func GetTwoPointDistance(lat1, lng1, lat2, lng2 float64) float64 {
	point1 := g.LatLngFromDegrees(lat1, lng1)
	point2 := g.LatLngFromDegrees(lat2, lng2)
	// 角度距离转换成实际距离（单位为米）地球半径
	return point1.Distance(point2).Radians() * 6371000.0
}

func IsValidLatitude(latitude float64) bool {
	return latitude >= -90 && latitude <= 90
}

func IsValidLongitude(longitude float64) bool {
	return longitude >= -180 && longitude <= 180
}
