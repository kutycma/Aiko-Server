package hy

import (
	"errors"
	"fmt"
	"github.com/Github-Aiko/Aiko-Server/api/panel"
	"github.com/Github-Aiko/Aiko-Server/src/conf"
	"github.com/Github-Aiko/Aiko-Server/src/limiter"
)

func (h *Hy) AddNode(tag string, info *panel.NodeInfo, c *conf.ControllerConfig) error {
	if info.Type != "hysteria" {
		return errors.New("the core not support " + info.Type)
	}
	switch c.CertConfig.CertMode {
	case "reality", "none", "":
		return errors.New("hysteria need normal tls cert")
	}
	l, err := limiter.GetLimiter(tag)
	if err != nil {
		return fmt.Errorf("get limiter error: %s", err)
	}
	s := NewServer(tag, l)
	err = s.runServer(info, c)
	if err != nil {
		return fmt.Errorf("run hy server error: %s", err)
	}
	h.servers.Store(tag, s)
	return nil
}

func (h *Hy) DelNode(tag string) error {
	if s, e := h.servers.Load(tag); e {
		err := s.(*Server).Close()
		if err != nil {
			return err
		}
		h.servers.Delete(tag)
		return nil
	}
	return errors.New("the node is not have")
}
