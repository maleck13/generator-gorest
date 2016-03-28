yeoman generator for a simple go rest api
heavily based on the work done here https://github.com/earlonrails/generator-go-microservice

dependencies

- gorrilla mux for routing
- negroni for middleware
- logrus for logging

optional

- prometheus for metrics
- mongo

there are options for database support

mongo vai mgo

mysql 


## usage

install yeoman

```npm install yo -g ```

``` npm install generator-gorest -g```

``` 
  
 mkdir $GOPATH/src/github.com/<user>/<appname>
 cd $GOPATH/src/github.com/<user>/<appname>
 yo gorest
 go get .
 go build .
 ./appname serve
 
 ```


## TODO 

- add a dockerfile
- add a docker-compose file