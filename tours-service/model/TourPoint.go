package model

/*
TourId
Name
Description
Latitude
Longitude
ImageUrl
Secret

*/

type TourPoint struct {
	ID          int     `json:"id"`
	TourId      int     `json:"tourId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	ImageUrl    string  `json:"imageUrl"`
}
