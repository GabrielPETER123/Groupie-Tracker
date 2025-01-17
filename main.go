package main

import (
	"Groupie-Tracker/API"
	"fmt"
)

var (
	Artists = []API.Artists{}
	Dates = API.IndexDates{}
	Locations = API.IndexLocations{}
	Relation = API.IndexRelation{}
)

func GetAllApi() {
	Artists = API.GetArtists()
	Dates = API.GetDates()
	Locations = API.GetLocations()
	Relation = API.GetRelation()
}


func main() {
	GetAllApi()
}