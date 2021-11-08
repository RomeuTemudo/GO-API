package models

type SensorsCategories struct {
	CategoryID          int    ` json:"id"` //key do param que vem do frontend
	CategoryDescription string `json:"description"`
}
