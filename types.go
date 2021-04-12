
package main

type Service struct {
	Id     int64  `json:"service_id"`
	UserId string `json:"user_id"`
	Date  string `json:"date_"`
	StateService   bool  `json:"state_service"`
	VehicleId   int64  `json:"vehicle_id"`
	Valor   float64  `json:"valor"`
	Cupos   int64  `json:"cupos"`

}
