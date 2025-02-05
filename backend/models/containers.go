package models

import "time"

type ContainerStatus struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	IPAddress   string    `gorm:"unique;not null" json:"ip_address"`
	Status      string    `json:"status"`
	LastChecked time.Time `json:"last_checked"`
}
