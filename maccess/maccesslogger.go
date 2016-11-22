package mlogaccess

import (
	"io"
	"log"
	"os"
)

const (
	LOGFILE_NAME   = "access.log"
	LOG_MSG_PREFIX = "[New Access Record] "
)

type AccessLogger struct {
	entry *log.Logger
}

func (al *AccessLogger) LogRecord(record AccessRecord) {
	al.entry.Printf(
		"[%s accessed from %s]\n[%s : %s]\n[%s : %s]\n[%d : %d]\n[%f : %f]",
		record.RequestedResource,
		record.RequestOrigin.IP,
		record.RequestOrigin.CountryName,
		record.RequestOrigin.CountryCode,
		record.RequestOrigin.City,
		record.RequestOrigin.ZipCode,
		record.RequestOrigin.MetroCode,
		record.RequestOrigin.AreaCode,
		record.RequestOrigin.Latitude,
		record.RequestOrigin.Longitude,
	)
}

func (al *AccessLogger) LogStatus(m string) {
	al.entry.Printf("[status : %s]", m)
}

func NewAccessLogger() *AccessLogger {
	logfile, err := os.OpenFile(
		LOGFILE_NAME,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)

	if err != nil {
		log.Printf(
			"error opening %s, log file: %s",
			LOGFILE_NAME,
			err,
		)
	}

	//defer logfile.Close() <- ????

	return &AccessLogger{
		entry: log.New(
			io.MultiWriter(logfile, os.Stdout),
			LOG_MSG_PREFIX,
			log.Ldate|log.Ltime|log.Lshortfile,
		),
	}
}
