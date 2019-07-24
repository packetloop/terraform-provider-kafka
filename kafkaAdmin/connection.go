package kafkaadmin

import kafka "github.com/arbor/go-kafkaesque"

func clientConn(m interface{}) *kafka.Client {
	return m.(*Conn).sclient
}
