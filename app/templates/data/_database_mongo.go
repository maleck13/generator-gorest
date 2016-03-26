package data

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"<%=basePackage %>/<%=baseName %>/config"
	"time"
)

var (
	mongodb *mgo.Session
	NO_MONGO = errors.New("no mongo session")
)

func InitMongo() *mgo.Session {
	if nil == mongodb {
		dbInfo := config.Conf.GetDatabase()
		if nil == dbInfo {
			logrus.Fatal("no db info ")
		}
		dbInfo.FailFast = true
		dbInfo.Timeout = 20 * time.Second
		session, err := mgo.DialWithInfo(dbInfo)
		if nil != err {
			logrus.Fatal("failed to connect to db ", err)
		}
		session.SetMode(mgo.Monotonic, true)
		mongodb = session
	}
	return mongodb
}

func NewMongoSession() (*mgo.Session, error) {
	if nil != mongodb {
		return mongodb.Copy(), nil
	}
	return nil, NO_MONGO
}

