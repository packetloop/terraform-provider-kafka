package kafkaadmin

import kafka "github.com/comozo/go-kafkaesque"

func clientConn(m interface{}) *kafka.Client {
	return m.(*Conn).sclient
}
