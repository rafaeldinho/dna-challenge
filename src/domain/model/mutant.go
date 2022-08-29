package model

import "github/meli/src/shared"

type Mutant struct {
	ID       int    `gorm:"primaryKey" json:"-"`
	ADN      string 
	IsMutant bool   
}

func (Mutant) TableName() string {
	return shared.MutantTable
}
