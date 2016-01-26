package mongoDBConfig

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	DEFAULT_MONGO_BATCH_SIZE int    = 1000
	MONGO_CONFIG_FILE        string = "/etc/config.json"
)

type BrandId uint32

type DBInfo struct {
	Host string
	User string
	Pass string
	DB   string
}

type VisitorId bson.ObjectId
