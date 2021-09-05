package main

import (
	"app/internal/ssh_tunnel"
	"fmt"
	"github.com/kardianos/service"
	"log"
	"os"
	"time"
)

type program struct {
}

func (p *program) Start(s service.Service) error {

	fmt.Println(time.Now())

	// Start should not block. Do the actual work async.
	go p.run()

	return nil
}
func (p *program) run() {
	h := &ssh_tunnel.TunnelHandler{}
	_ = h.Do()
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.

	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "SshTunnel",
		DisplayName: "SshTunnel Service",
		Description: "SshTunnel Service.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	// logger 用于记录系统日志
	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) == 2 { //如果有命令则执行
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else { //否则说明是方法启动了
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
	}
	if err != nil {
		logger.Error(err)
	}
}
