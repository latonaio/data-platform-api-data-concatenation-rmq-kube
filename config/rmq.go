package config

import (
	"fmt"
	"os"
	"strings"
)

type RMQ struct {
	user  string
	pass  string
	addr  string
	port  string
	vhost string

	queueFrom string
	queueTo   []string

	queueToErrResponse string
}

func newRMQ() *RMQ {
	return &RMQ{
		user:               os.Getenv("RMQ_USER"),
		pass:               os.Getenv("RMQ_PASS"),
		addr:               os.Getenv("RMQ_ADDRESS"),
		port:               os.Getenv("RMQ_PORT"),
		vhost:              os.Getenv("RMQ_VHOST"),
		queueFrom:          os.Getenv("RMQ_QUEUE_FROM"),
		queueTo:            getEnvStrings("RMQ_QUEUE_TO"),
		queueToErrResponse: os.Getenv("NESTJS_DATA_CONNECTION_REQUEST_CONTROL_MANAGER_CONSUME"),
	}
}

func (c *RMQ) URL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/%s", c.user, c.pass, c.addr, c.port, c.vhost)
}

func (c *RMQ) QueueFrom() string {
	return c.queueFrom
}

func (c *RMQ) QueueTo() []string {
	return c.queueTo
}

func (c *RMQ) QueueToErrResponse() string {
	return c.queueToErrResponse
}

func getEnvStrings(key string) []string {
	rawVal := os.Getenv(key)
	rawVal = strings.ReplaceAll(rawVal, "\\ ", "$THIS_SECTION_IS_SPACE")
	rawVal = strings.ReplaceAll(rawVal, " ", "")
	rawVal = strings.ReplaceAll(rawVal, "$THIS_SECTION_IS_SPACE", " ")
	val := strings.Split(rawVal, ",")
	return val
}
