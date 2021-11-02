package models

type Sensor struct {
	SensorID   int    ` json:"id"`
	SensorName string `json:"name"`
}

//nome da tabela da BD, caso nao use o gorm para gerar tabelas(em vez de usar o sensor como nome da tabela posso meter o que quiser)
/*func (Sensor) TableName() string {
	return "sensor_list"
}*/

/*func (s *Sensor) Store() {
 */
