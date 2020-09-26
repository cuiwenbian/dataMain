package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Image Images struct
type Image struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FileName    string             `json:"fileName"`
	FullPath    string             `json:"fullPath"`
	OS          string             `json:"os"`
	Path        string             `json:"path"`
	Size        int64              `json:"size"`
	Cid         string             `json:"cid"`
	Updatedtime time.Time          `json:"updatedtime"`
	Createdtime time.Time          `json:"createdtime"`
}
