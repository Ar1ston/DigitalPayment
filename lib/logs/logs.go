package logs

import "fmt"

const (
	log_error = "ERROR"
	log_info  = "INFO"
	log_debug = "DEBUG"
)

type Log struct {
	serviceName string
	requestName string
}

var Logger Log

func New() *Log {
	return &Log{}
}

func (log *Log) SetService(serviceName string) *Log {
	log.serviceName = serviceName
	return log
}
func (log *Log) SetRequest(requestName string) *Log {
	log.requestName = requestName
	return log
}
func (log *Log) Printf(format string, params ...interface{}) {
	if log.requestName != "" {
		fmt.Printf("[%s][%s] %s\n", log.serviceName, log.requestName, fmt.Sprintf(format, params))
	} else {
		fmt.Printf("[%s] %s\n", log.serviceName, fmt.Sprintf(format, params))
	}
}
func (log *Log) Errorf(format string, params ...interface{}) {
	if log.requestName != "" {
		fmt.Printf("[%s][%s][%s] %s\n", log_error, log.serviceName, log.requestName, fmt.Sprintf(format, params))
	} else {
		fmt.Printf("[%s][%s] %s\n", log_error, log.serviceName, fmt.Sprintf(format, params))
	}
}
func (log *Log) Error(str string) {
	if log.requestName != "" {
		fmt.Printf("[%s][%s][%s] %s\n", log_error, log.serviceName, log.requestName, str)
	} else {
		fmt.Printf("[%s][%s] %s\n", log_error, log.serviceName, str)
	}
}
func (log *Log) Infof(format string, params ...interface{}) {
	if log.requestName != "" {
		fmt.Printf("[%s][%s][%s] %s\n", log_info, log.serviceName, log.requestName, fmt.Sprintf(format, params))
	} else {
		fmt.Printf("[%s][%s] %s\n", log_info, log.serviceName, fmt.Sprintf(format, params))
	}
}
func (log *Log) Info(str string) {
	if log.requestName != "" {
		fmt.Printf("[%s][%s][%s] %s\n", log_info, log.serviceName, log.requestName, str)
	} else {
		fmt.Printf("[%s][%s] %s\n", log_info, log.serviceName, str)
	}
}
func (log *Log) Debugf(format string, params ...interface{}) {
	if log.requestName != "" {
		fmt.Printf("[%s][%s][%s] %s\n", log_debug, log.serviceName, log.requestName, fmt.Sprintf(format, params))
	} else {
		fmt.Printf("[%s][%s] %s\n", log_debug, log.serviceName, fmt.Sprintf(format, params))
	}
}
func (log *Log) Debug(str string) {
	if log.requestName != "" {
		fmt.Printf("[%s][%s][%s] %s\n", log_debug, log.serviceName, log.requestName, str)
	} else {
		fmt.Printf("[%s][%s] %s\n", log_debug, log.serviceName, str)
	}
}
