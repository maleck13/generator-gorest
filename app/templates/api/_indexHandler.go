package api

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/context"
	"<%=basePackage %>/<%=baseName %>/config"
	<% if("mongo" === database || "yes" == messaging) { %>
	"<%=basePackage %>/<%=baseName %>/data"
	<% } %>
	<% if("mongo" === database) { %>
	"<%=basePackage %>/<%=baseName %>/data/model"
	"<%=basePackage %>/<%=baseName %>/domain"
	<% } %>
	<% if( "yes" == messaging) { %>
	"github.com/maleck13/stompy"
	"github.com/Sirupsen/logrus"
	<% } %>
)

//Example route handler
func IndexHandler(rw http.ResponseWriter, req *http.Request) HttpError {
	encoder := json.NewEncoder(rw)
	resData := make(map[string]string)
	resData["example"] = config.Conf.GetExample()

	val,has := context.GetOk(req,"test")
	if has{
		resData["context"] = val.(string)
	}

	if err := encoder.Encode(resData); err != nil {
		return NewHttpError(err, http.StatusInternalServerError)
	}
	return nil
}

<% if("mongo" === database ) { %>
//example of using mongo with a domain service. Doing two things here creating an document and then using using a service to do domain specific logic
// its a forced example. Normally you wouldn't be creating the document like this
func IndexMongo(rw http.ResponseWriter, req *http.Request)HttpError{
	var encoder = json.NewEncoder(rw)
	//get a new mongo session
	mongodb ,err := data.MongoSession()
	if err != nil{
		return NewHttpError(err, http.StatusInternalServerError)
	}
	//ensure we close after the method is complete
	defer mongodb.Close()
	//get a new model and set some values
	exampleModel := model.NewExample()
	exampleModel.Test = "test"
	//create a new repo and save to mongo
	exampleRepp := model.NewExampleRepo(mongodb)
	exampleModel, err = exampleRepp.Save(exampleModel)
	if err != nil{
		return NewHttpErrorWithContext(err,http.StatusInternalServerError,"failed to save to mongo")
	}
	//get a service
	//passing in example rep twice as it is a saver and a finder
	exampleService := domain.NewExampleService(exampleRepp,exampleRepp)
	//use our domain logic
	exampleModel, err = exampleService.IncrementExampleOrders(exampleModel.Id)
	if err != nil{
		return NewHttpError(err,http.StatusInternalServerError)
	}

	if err := encoder.Encode(exampleModel); err != nil{
		return NewHttpError(err, http.StatusInternalServerError)
	}
	return nil
}

<% } %>

<% if( "yes" == messaging) { %>
//example of publishing to stomp
func IndexStomp(rw http.ResponseWriter, req *http.Request)HttpError{
	resData := make(map[string]string)
	resData["example"] = config.Conf.GetExample()

	data.Subscribe("test","test",func(msg stompy.Frame){
		jsonData := make(map[string]string)
		if err := json.Unmarshal(msg.Body,&jsonData); err != nil{
			logrus.Error("failed to unmarshal msg ", err.Error())
		}
		logrus.Info("handling msg 1: ", jsonData)

	},nil)

	data.Subscribe("test","test",func(msg stompy.Frame){
		jsonData := make(map[string]string)
		if err := json.Unmarshal(msg.Body,&jsonData); err != nil{
			logrus.Error("failed to unmarshal msg ", err.Error())
		}
		logrus.Info("handling msg 2: ", jsonData)
	},nil)

	for i :=0; i < 10; i++ {
		data.Publish("test", "test",resData, nil)
	}
	return nil
}

<% } %>