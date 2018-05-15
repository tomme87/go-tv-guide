package config

import (
	"github.com/spf13/viper"
	"github.com/tomme87/go-tv-guide/internal/pkg/storage"
	"reflect"
	"testing"
)

func TestConfig_Init(t *testing.T) {
	viper.AddConfigPath("../../../test")

	err := C.Init()
	if err != nil {
		t.Fatal(err)
	}

	expected := config{
		MongoDB: storage.MongoDB{
			DatabaseURI:             "mongodb://usr:pw@localhost/admin",
			DatabaseName:            "tv",
			ChannelCollectionName:   "test_channels",
			ProgrammeCollectionName: "test_programmes",
		},
	}

	if !reflect.DeepEqual(expected, C) {
		t.Fatal("Not the expected config struct")
	}
}
