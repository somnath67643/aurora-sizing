package ssclient

import (
	"log"

	"github.com/somnath67643/aurora-sizing/clientgo/ssclient/clbstream"
	"github.com/somnath67643/aurora-sizing/clientgo/ssclient/clgrpc"
	"github.com/somnath67643/aurora-sizing/clientgo/ssclient/cltcp"
)

type ClientType int

const (
	GRPC ClientType = iota
	TCP
	BSTREAM
)

type SSClient interface {
	Close()
	Init()
	UseService(dbname string, closure func()) error
	LookupByName(prefix string, cmpType string) ([]string, error)
	LookupByType(prefix string, stype string, cmpType string) ([]string, error)
	GetObject(sname string)
	GetObjectMany(snames []string)
}

func NewSSClient(addr string, typ ClientType) SSClient {
	var cl SSClient
	switch typ {
	case GRPC:
		cl = clgrpc.NewSSClient(addr)
	case TCP:
		cl = cltcp.NewSSClient(addr)
	case BSTREAM:
		cl = clbstream.NewSSClient(addr)
	default:
		log.Fatal("Invalid ClientType")
	}
	return cl
}
