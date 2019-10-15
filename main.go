package main

import (
	"github.com/j178/sshd/sshlog"
	"github.com/j178/sshd/sshserver"
	"github.com/sirupsen/logrus"
)

func main() {
	logManager := sshlog.NewEmptyManager()
	logger := logrus.New()
	shutdownC := make(chan struct{})
	sshd, err := sshserver.New(logManager, logger, "1.1", "127.0.0.1:2224", shutdownC, false)
	if err != nil {
		msg := "Cannot create new SSH Server"
		logger.WithError(err).Error(msg)
		return
	}
	logger.Fatal(sshd.Start())
}
