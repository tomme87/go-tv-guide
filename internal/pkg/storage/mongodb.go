package storage

import (
	"github.com/tomme87/go-tv-guide/pkg/xmltv"
	"gopkg.in/mgo.v2"
	"log"
)

// MongoDB data
type MongoDB struct {
	DatabaseURI             string `mapstructure:"uri"`
	DatabaseName            string `mapstructure:"name"`
	ChannelCollectionName   string `mapstructure:"channel_collection"`
	ProgrammeCollectionName string `mapstructure:"programme_collection"`
}

var db *mgo.Database

// Connect to the MongoDB database.
func (m *MongoDB) Connect() {
	session, err := mgo.Dial(m.DatabaseURI)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(m.DatabaseName)
}

// Insert data to db
func (m *MongoDB) Insert(data []interface{}, collection string) error {
	err := db.C(collection).Insert(data...)
	return err
}

// InsertChannels channels to db
func (m *MongoDB) InsertChannels(channels []xmltv.Channel) error {
	var ui []interface{}
	for _, t := range channels {
		ui = append(ui, t)
	}
	return m.Insert(ui, m.ChannelCollectionName)
}

// InsertProgrammes programmes to db
func (m *MongoDB) InsertProgrammes(programmes []xmltv.Programme) error {
	var ui []interface{}
	for _, t := range programmes {
		ui = append(ui, t)
	}
	err := db.C(m.ProgrammeCollectionName).Insert(ui...)
	return err
}
