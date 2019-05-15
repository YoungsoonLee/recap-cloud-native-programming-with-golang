package dblayer

import (
	"github.com/YoungsoonLee/cloud_native/lib/mongolayer"
	"github.com/YoungsoonLee/cloud_native_recap/myEvents/events/models"
)

type DBTYPE string

const (
	MONOGODB DBTYPE = "mongodb"
	DYNAMODB DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (models.DatabaseHandler, error) {
	switch options {
	case MONOGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
}
