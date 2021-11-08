package models

type Sensor struct {
	SensorID          int    `gorm:"primaryKey" json:"id"` //key do param que vem do frontend
	SensorName        string `json:"name"`
	SensorDescription string `json:"description"`
	CategoryID        int    `json:"category_id"`
}

//nome da tabela da BD, caso nao use o gorm para gerar tabelas(em vez de usar o sensor como nome da tabela posso meter o que quiser)
/*func (Sensor) TableName() string {
	return "sensor_list"
}*/

/*func (s *Sensor) Store() {
 */
