yeoman generator for a simple go rest api
heavily based on the work done here https://github.com/earlonrails/generator-go-microservice

dependencies

- gorrilla mux for routing
- negroni for middleware
- logrus for logging

optional

- prometheus for metrics
- mongo
- stomp (messaging)
## usage

install yeoman

```npm install yo -g ```

``` npm install generator-gorest -g```

``` 
  
 mkdir $GOPATH/src/github.com/<user>/<appname>
 cd $GOPATH/src/github.com/<user>/<appname>
 yo gorest
 go get .
 (optional) go get github.com/tools/godep
 (optional) godep save -r ./... 
  docker build -t <appname> .
  docker-compose up -d
  
  curl http://<dockermachineip>:3000/sys/info/ping
  
 ```

