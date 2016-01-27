package mongoDBConfig

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	DEFAULT_MONGO_BATCH_SIZE int    = 1000
	MONGO_CONFIG_FILE        string = "/etc/config.json"
)

// our db is comprised of multiple brands' visitor info
// we make the distinction by assigning BrandIds
type BrandId uint32

// keep track of unique visitors using Mongo's automatically assigned ids
type VisitorId bson.ObjectId

type DBInfo struct {
	Host string
	User string
	Pass string
	DB   string
}
