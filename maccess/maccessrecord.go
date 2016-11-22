package mlogaccess

import (
	"errors"
	"github.com/msawangwan/mgeoloc"
)

var NilRecordError = errors.New("tried to read a nil record")

type AccessRecord struct {
	RequestedResource string
	RequestOrigin     mgeoloc.GeographicLocation
}

func NewAccessRecord(ip, resource string) *AccessRecord {
	return &AccessRecord{
		RequestedResource: resource,
		RequestOrigin:     mgeoloc.FromAddr(ip),
	}
}
