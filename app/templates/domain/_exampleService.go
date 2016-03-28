package domain

import (
	"<%=basePackage %>/<%=baseName %>/data/model"
	"gopkg.in/mgo.v2/bson"
)

//this is a forced example. Just really to show how we use the mongo db stuff
type ExampleService struct {
	finderFetcher model.ExampleFinderFetcher
	saver model.ExampleSaver
}

func NewExampleService(finder model.ExampleFinderFetcher, saver model.ExampleSaver) ExampleService{
	return ExampleService{finderFetcher:finder,saver:saver}
}


func (es ExampleService) IncrementExampleOrders(exampleId bson.ObjectId)(*model.Example,error){
	example,err := es.finderFetcher.FindById(exampleId)
	if err != nil{
		return nil,err
	}
	example.Orders +=1
	example,err = es.saver.Save(example)
	if err != nil{
		return nil,err
	}
	return example,err
}
