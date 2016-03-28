package data

import (
	"<%=basePackage %>/<%=baseName %>/config"
	"net"
	"github.com/Sirupsen/logrus"
	"github.com/gmallard/stompngo"
	"encoding/json"
)

var StompConn *stompngo.Connection
var netConn net.Conn

func InitStomp(connectionDetails *config.Stomp_config)  {
	var err error
	address := net.JoinHostPort(connectionDetails.Host,connectionDetails.Port)
	netConn,err = net.Dial("tcp", address)
	if err != nil{
		logrus.Fatal("failed to connect via net ", err)
	}
	heads := headers(connectionDetails)
	StompConn, err = stompngo.Connect(netConn, heads)
	if nil != err{
		logrus.Fatal("failed to connect via stomp ", err)
	}

}

func DestroyStomp(){
	if nil != StompConn{
		if err := StompConn.Disconnect(stompngo.Headers{}); err != nil{
			logrus.Error("failed to disconnect from stomp server ", err)
		}
	}
	if nil != netConn{
		if err := netConn.Close(); err != nil{
			logrus.Error("failed to disconnect from network ", err)
		}
	}
}

func headers(connectionDetails *config.Stomp_config)stompngo.Headers{
	heads := stompngo.Headers{}
	heads = heads.Add("login",connectionDetails.User).Add("passcode",connectionDetails.Pass)
	heads = heads.Add("accept-version", connectionDetails.Protocol).Add("host",connectionDetails.Vhost)
	logrus.Info(heads)
	return heads
}


type SubscriptionHandler func(stompngo.MessageData)
type StompOpts map[string]string

//used to subscribe to messages
func Subscribe(queue, topic string ,handler SubscriptionHandler, opts StompOpts )error{
	destination := queue + "/" + topic
	id := stompngo.Uuid()
	h:= stompngo.Headers{}
	h = h.Add("destination",destination).Add("ack","auto").Add("id",id)
	if nil != opts{
		//override anything
		for k,v := range opts{
			h = h.Add(k,v)
		}
	}
	mesageChan, err := StompConn.Subscribe(h)
	if err != nil{
		logrus.Error("failed to subscribe to " + destination, err.Error())
		return err
	}
	go func(incoming <-chan stompngo.MessageData, mHandler SubscriptionHandler){
		for m := range incoming{
			handler(m)
		}
	}(mesageChan, handler)
	return nil
}

//publish messages
func Publish(queue, topic string , jsonData interface{} , opts StompOpts)error{
	destination := queue + "/" + topic
	s := stompngo.Headers{"destination", destination , "persistent", "true","content-type","application/json"} // send headers
	if nil != opts{
		//override anything
		for k,v := range opts{
			s = s.Add(k,v)
		}
	}
	data,err:=json.Marshal(jsonData)
	if err != nil{
		return err
	}
	StompConn.SendBytes(s,data)
	return nil
}