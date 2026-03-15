package domain

type Electronics struct {
	ID           int64 `json:"electronics_id"  db:"electronics_id"`
	ControllerID int64 `json:"controller_id"   db:"controller_id"`
	SensorID     int64 `json:"sensor_id"       db:"sensor_id"`
	WiringID     int64 `json:"wiring_id"       db:"wiring_id"`
}

type Controller struct {
	ID   int64  `json:"controller_id"   db:"controller_id"`
	Name string `json:"controller_name" db:"controller_name"`
	Info string `json:"controller_info" db:"controller_info"`
}

type Sensor struct {
	ID   int64  `json:"sensor_id"   db:"sensor_id"`
	Name string `json:"sensor_name" db:"sensor_name"`
	Info string `json:"sensor_info" db:"sensor_info"`
}

type Wiring struct {
	ID   int64  `json:"wiring_id"   db:"wiring_id"`
	Name string `json:"wiring_name" db:"wiring_name"`
	Info string `json:"wiring_info" db:"wiring_info"`
}
