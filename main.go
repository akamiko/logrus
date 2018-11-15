package main

import (
	"context"

	"github.com/akamiko/logging-sample4/lib/logger/logrus"
)

func main() {
	l := logrus.NewLogger()

	l.Info(context.Background(), "message= %s ", "info message!!")
}
