package data

import (
	"<%=basePackage %>/<%=baseName %>/config"
	"net"
	"github.com/Sirupsen/logrus"
	"github.com/maleck13/stompy"
	"encoding/json"
	"time"
)

var stompClient stompy.StompClient

func InitStomp(connectionDetails *config.Stomp_config)  {
	var err error
	address := net.JoinHostPort(connectionDetails.Host,connectionDetails.Port)
	clientOpts := stompy.ClientOpts{
		Vhost:connectionDetails.Vhost,
		HostAndPort:address,
		Timeout:time.Second * 10,
		User:connectionDetails.User,
		PassCode:connectionDetails.Pass,
		Version:connectionDetails.Protocol,
	}
	stompClient =stompy.NewClient(clientOpts)
	err = stompClient.Connect()
	if nil != err{
		logrus.Fatal("failed to connect via stomp ", err)
	}

}

func DestroyStomp(){
	if nil != stompClient{
		if err := stompClient.Disconnect(); err != nil{
			logrus.Error("failed to disconnect from stomp server ", err)
		}
	}
}

//used to subscribe to messages
func Subscribe(queue, topic string ,handler stompy.SubscriptionHandler, opts stompy.StompHeaders )error{
	destination := queue + "/" + topic

	_, err := stompClient.Subscribe(destination,handler,opts,nil)
	if err != nil{
		logrus.Error("failed to subscribe to " + destination, err.Error())
		return err
	}
	return nil
}

//publish messages
func Publish(queue, topic string , jsonData interface{} , opts stompy.StompHeaders)error{
	destination := queue + "/" + topic
	data,err:=json.Marshal(jsonData)
	if err != nil{
		return err
	}
	if err := stompClient.Publish(destination,"application/json",data,opts,nil); err != nil{
		return err
	}
	return nil
}