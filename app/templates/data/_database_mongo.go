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

//this method should only be used for initailising the mongo connection. It is a once of task. To get a session for regular queries use MongoSession
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
		logrus.Info("mongodb connected")
	}
	return mongodb
}

func DestroyMongo(){
	if nil != mongodb{
		mongodb.Close()
	}
}

//Gives you a copy of the mongo session for use by a single go routine
// wrap up the session and database for convenience
type Mongo struct {
	*mgo.Session
	Database *mgo.Database
}

func MongoSession() (*Mongo, error) {
	if nil != mongodb {
		session := mongodb.Copy()
		db := session.DB(config.Conf.GetDatabase().Database)
		return &Mongo{session,db},nil
	}
	return nil, NO_MONGO
}

