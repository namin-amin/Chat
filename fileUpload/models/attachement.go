package models

import (
	"Chat/base"
)

// Attachment
//
//	represents the diff media type
//
// Todo types of media etc.
type Attachment struct {
	base.BaseModel
	UserId string `json:"userId"`
	Name   string `json:"name"`
	Path   string `json:"path"`
}

// GetSavedPath
//
// Gets the full path of the saved attachment
func (a *Attachment) GetSavedPath() string {
	return a.Path + a.Name
}
