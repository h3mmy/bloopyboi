package models

import (
	"fmt"
	"reflect"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ServiceRegistry interface {
	AddService(bloopyService BloopyService) error
	GetService(svcName string) (*BloopyService, bool)
	Size() int
}

type BloopyServiceBroker struct {
	services map[string]*BloopyService
	logger   *zap.Logger
}

func NewBloopyServiceBroker() *BloopyServiceBroker {
	zlogConfig := zap.NewDevelopmentConfig()
	zlogConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	zlogger, _ := zlogConfig.Build()

	return &BloopyServiceBroker{
		services: make(map[string]*BloopyService),
		logger:   zlogger,
	}
}

func (bsb *BloopyServiceBroker) AddService(bloopyService BloopyService) error {
	svcName := reflect.TypeOf(bloopyService).Name()
	bsb.services[svcName] = &bloopyService
	bsb.logger.Info(fmt.Sprintf("Adding service with name %s: %v", svcName, bloopyService))
	return nil
}

func (bsb *BloopyServiceBroker) GetService(svcName string) (*BloopyService, bool) {
	svc, ok := bsb.services[svcName]
	if ok {
		return svc, true
	}
	bsb.logger.Warn(fmt.Sprintf("No Registry Entry for Service: %s", svcName))
	return nil, false
}

func (bsb *BloopyServiceBroker) Size() int {
	return len(bsb.services)
}
