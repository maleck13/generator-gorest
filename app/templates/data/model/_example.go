/**
This file has several definitions. It has the model and the repository. It also has several interface definitions.
The interface definitions are to improve how testable the code is. Using interfaces allows us to override the behaviour in our test cases
 */

package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"errors"
	"<%=basePackage %>/<%=baseName %>/data"
)

var(
	exampleTestEmptyError = errors.New("test cannot be empty")
)

//using interfaces like this allows for simpler non dependant unit testing as we can replace the implementation with a mocked out one

type ExampleFinder interface {
	FindById(bson.ObjectId) (*Example, error)
	FindWhereField(string, interface{}) (*Example, error)
}

type ExampleFetcher interface {
	Fetch() ([]*Example, error)
	FetchWhereField(string, interface{}) ([]*Example, error)
}

type ExampleSaver interface {
	Save(*Example) (*Example, error)
}

type ExampleFinderFetcher interface {
	ExampleFetcher
	ExampleFinder
}

//this is the model we will be persisting
type Example struct {
	Test string
	Orders int
	Id   bson.ObjectId `bson:"_id,omitempty"`
}

//adding simple validation
func (e *Example)Validate()error{
	if e.Test == ""{
		return exampleTestEmptyError
	}
	return nil
}

func NewExample() *Example {
	return &Example{Id:bson.NewObjectId()}
}

const EXAMPLE_COL = "example"

type ExampleRepo struct {
	collection *mgo.Collection
}


func NewExampleRepo(mongo *data.Mongo) *ExampleRepo {
	collection := mongo.Database.C(EXAMPLE_COL)
	return &ExampleRepo{collection:collection}
}

//these two make it fulfill the ExampleFinder interface
func (er *ExampleRepo)FindById(id bson.ObjectId) (*Example, error) {
	m := &Example{}
	er.collection.FindId(id).One(m)
	return m, nil
}

func (er *ExampleRepo)FindWhereField(field string,value interface{}) (*Example, error)  {
	result := &Example{}
	if err := er.collection.Find(bson.M{field:value}).One(result); err != nil{
		return nil,err
	}
	return result,nil
}

//these make it implement the ExampleFetcher interface
func (er *ExampleRepo) Fetch() ([]*Example, error){
	var result []*Example
	if err := er.collection.Find(bson.M{}).All(&result); err != nil{
		return nil,err
	}
	return result,nil
}

func (er *ExampleRepo)FetchWhereField(field string, value interface{}) ([]*Example, error){
	var result []*Example
	if err := er.collection.Find(bson.M{field:value}).All(&result); err != nil{
		return nil,err
	}
	return result,nil
}



//this makes it an ExampleSaver
func (er *ExampleRepo)Save(example *Example) (*Example, error) {
	if err := example.Validate(); err != nil{
		return nil,err
	}
	if "" == example.Id {
		example.Id = bson.NewObjectId()
	}
	if _, err := er.collection.UpsertId(example.Id, example); err != nil {
		return nil, err
	}
	return example, nil
}



