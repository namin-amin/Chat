package base

import "time"

// BaseModel
//
// Model contains base common columns that's needed in all models
type BaseModel struct {
	Id      string    `gorm:"<-:create;type:string" json:"id"`     //Id of the Model represents unique primary key
	Created time.Time `gorm:"autoCreateTime:milli" json:"created"` //created time of the row
	Updated time.Time `gorm:"autoUpdateTime:milli" json:"updated"` //updating time of the row
}
