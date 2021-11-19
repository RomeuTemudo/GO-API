package models

type SensorData struct {
	ID   int           `json:"id;omitempty"`
	Data []SensorValue `json:"data"`
}

type SensorValue struct {
	Value     float32 `json:"value"`
	Timestamp string  `gorm:"primaryKey;type:timestamp" json:"timestamp"`
}

//`gorm:"primaryKey;type:timestamp" json:"timestamp"`
