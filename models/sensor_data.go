package models

type SensorData struct {
	SensorID  int     `json:"id"` //key do param que vem do frontend
	Value     float64 `json:"value"`
	Timestamp int64   ` gorm: autoCreateTime time ;json:"timestamp"`
}
