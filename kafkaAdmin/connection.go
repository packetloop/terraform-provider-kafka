package kafkaadmin

import kafka "github.com/packetloop/go-kafkaesque"

func clientConn(m interface{}) *kafka.Client {
	return m.(*Conn).sclient
}
