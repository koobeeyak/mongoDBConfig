package mongoDBConfig

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"gopkg.in/mgo.v2"
)

// Load MongoDB host, user, pass etc. so we can connect
func (db *DBInfo) LoadFromConfigFile(configFile string) error {
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Couldn't read config file: %v", err)
		return err
	}
	var config map[string]string
	// unmarshal json from config file into map
	if err := json.Unmarshal(content, &config); err != nil {
		log.Fatalf("Could not decode Json, %v", err)
		return err
	}
	db.Host = config["mongo_host"]
	db.Pass = config["mongo_pass"]
	db.DB = config["mongo_db"]
	db.User = config["mongo_user"]
	return nil
}

// Establish new connection to MongoDB host
// Then need to select specific collection
func NewDBConn() *mgo.Database {
	db := DBInfo{}
	err := db.LoadFromConfigFile(MONGO_CONFIG_FILE)
	if err != nil {
		panic(err)
	}
	session, err := mgo.Dial(db.Host)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetBatch(DEFAULT_MONGO_BATCH_SIZE)
	collections := session.DB(db.DB)
	if err := collections.Login(db.User, db.Pass); err != nil {
		log.Printf("Cannot login to DB.")
		return nil
	}
	return collections
}
