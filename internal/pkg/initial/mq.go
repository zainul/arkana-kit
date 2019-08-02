package initial

import (
	"fmt"

	"github.com/streadway/amqp"
)

var mqConnection *amqp.Connection

func init() {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", conf.MQ.User, conf.MQ.Password, conf.MQ.Server, conf.MQ.Port))

	if err != nil {
		fmt.Println("failed to run mq")
		return
	}

	mqConnection = conn

	fmt.Println("success to connect mq")
}

// GetMQ ...
func GetMQ() *amqp.Connection {
	return mqConnection
}
