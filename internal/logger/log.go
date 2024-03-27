package logger

import "go.uber.org/zap"

var root *zap.Logger

func NewLogger(name string) *zap.Logger{
	return root.Named(name)
}	