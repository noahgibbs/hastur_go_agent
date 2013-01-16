package main

import (
       "fmt"
       "net"
       "encoding/json"
       zmq "github.com/alecthomas/gozmq"
)

type HasturZMQ struct {
     context zmq.Context
     socket zmq.Socket
     sequence int64
}

type Hastur interface {
     Send(message map[string] interface{}) error
     Close()
}

func HasturSender() (h Hastur) {
     h := new(HasturZMQ)

     h.context, _ = zmq.NewContext()
     h.socket, _ = context.NewSocket(zmq.DEALER)

     h.socket.Connect("tcp://hastur.ooyala.com:8126")
}

func (h HasturZMQ*) SendToHastur(messageMap map[string] interface{}) error {
    msg, _ := json.Marshal(messageMap)

    h.socket.Send([]byte(msg), 0)
}

func (h HasturZMQ*) Close() {
    h.socket.Close()
    h.context.Close()
}
