package storage

import "github.com/tomme87/go-tv-guide/pkg/xmltv"

// Storage interface for storing channels and programmes
type Storage interface {
	Connect()
	Insert(data []interface{}, collection string) error
	InsertChannels(channels []xmltv.Channel) error
	InsertProgrammes(programmes []xmltv.Programme) error
}
