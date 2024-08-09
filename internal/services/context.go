package services

import (
	"context"
	"covid-journal/internal/logging"
	"reflect"
)

type ServiceKey string

const serviceKey ServiceKey = "services"

type serviceMap map[reflect.Type]reflect.Value

func NewServiceContext(c context.Context) context.Context {
	if c.Value(serviceKey) == nil {
		return context.WithValue(c, serviceKey, make(serviceMap))
	} else {
		return c
	}
}

func LoggerFactory() logging.Logger {
	// TODO: create struct that implements Logger interface
	return nil
}
