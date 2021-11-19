package models

type Sensor struct {
	SensorID          int    `gorm:"primaryKey" json:"id"` //key do param que vem do frontend
	SensorName        string `json:"name"`
	SensorDescription string `json:"description"`
	CategoryID        int    `json:"category_id"`
}
