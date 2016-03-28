package domain_test

import (
	"testing"
	"gopkg.in/mgo.v2/bson"
	"<%=basePackage %>/<%=baseName %>/data/model"
	"<%=basePackage %>/<%=baseName %>/domain"
	"github.com/stretchr/testify/assert"
)

func getExample(id bson.ObjectId,value string)*model.Example{
	return &model.Example{
		Id:id,
		Test:value, //turn interface into string
	}
}

//set up our mock db layer
type MockExampleRepo struct {}

//implement our mocks ensuring it fulfills the right interfaces
func (mr *MockExampleRepo) FindById(id bson.ObjectId)(*model.Example,error){
	return getExample(id,"test"),nil
}

func (mr *MockExampleRepo) FindWhereField(field string, value interface{}) (*model.Example, error){
	return getExample(bson.NewObjectId(),"test"),nil
}

func (mr *MockExampleRepo) Fetch() ([]*model.Example, error){
	return []*model.Example{
		getExample(bson.NewObjectId(),"Test"),
	},nil
}

func (mr *MockExampleRepo)FetchWhereField(field string, value interface{}) ([]*model.Example, error){
	return []*model.Example{
		getExample(bson.NewObjectId(),value.(string)), //cast the interface to a string
	},nil
}

func (mr *MockExampleRepo) Save(ex *model.Example)(*model.Example,error){
	return ex,nil
}


func TestIncrementExampleOrders(t *testing.T){
	mock := &MockExampleRepo{}
	exampleService := domain.NewExampleService(mock,mock) //our mock fulfills both interfaces
	example,err := exampleService.IncrementExampleOrders(bson.NewObjectId())
	assert.NoError(t,err,"did not expect an error")
	assert.Equal(t,1,example.Orders,"expected orders to be incremented")
}
