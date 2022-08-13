package config

import (
	"github.com/apache/thrift/lib/go/thrift"
)

type Configure interface {
	Configure(method string, json string) error
	GetProcessor() thrift.TProcessor
}
