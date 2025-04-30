package main

type City struct {
	code      string
	cityName  string
	longitude int
	latitude  int
}

var cities = []City{
	{"LAX", "Los Angeles", 118, 34},
	{"JFK", "New York", -74, 40},
}
