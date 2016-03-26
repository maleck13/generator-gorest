package config

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	<% if(database == 'mongo'){ %>
         "gopkg.in/mgo.v2"
        <% } %>
	"io/ioutil"
	"os"
)

//use an interface to limit access to the config object to read only
type Configuration interface {
	GetPProfEnabled() bool
	GetExample()string
	<% if(database == 'mongo'){ %>
	 GetDatabase() *mgo.DialInfo
        <% } %>

}

type config struct {
	PProfEnabled  bool `json:"pprof_enabled"`
	Example string `json:"example"`
	<% if(database == 'mongo'){ %>
        Database *mgo.DialInfo `json:"database"`
	<% } %>
}

func (c *config) GetExample()string{
  return c.Example
}

func (c *config) GetPProfEnabled() bool {
	return c.PProfEnabled
}
<% if(database == 'mongo'){ %>
func (c *config) GetDatabase() *mgo.DialInfo {
	return c.Database
}
<% } %>

var Conf Configuration

func SetGlobalConfig(path string) {
	Conf = &config{}
	file, err := os.Open(path)
	if nil != err {
		logrus.Panic("failed to open config file ", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if nil != err {
		logrus.Panic("failed to read config file ", err)
		return
	}
	if err = json.Unmarshal(data, Conf); err != nil {
		logrus.Panic("failed to decode config file ", err)
		return
	}
}

