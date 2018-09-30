package tars

import (
	"fmt"
	logger "github.com/TarsCloud/TarsGo/tars/util/rogger"
	"strings"
)

type Admin struct {
}

func (a *Admin) Shutdown() error {
	for obj, s := range goSvrs {
		TLOG.Debug("shutdown", obj)
		//TODO
		go s.Shutdown()
	}
	shutdown <- true
	return nil
}

func (a *Admin) Notify(command string) (string, error) {
	cmd := strings.Split(command, " ")
	switch cmd[0] {
	case "tars.viewversion":
		return GetServerConfig().Version, nil
	case "tars.setloglevel":
		switch cmd[1] {
		case "INFO":
			logger.SetLevel(logger.INFO)
		case "WARN":
			logger.SetLevel(logger.WARN)
		case "ERROR":
			logger.SetLevel(logger.ERROR)
		case "DEBUG":
			logger.SetLevel(logger.DEBUG)
		case "NONE":
			logger.SetLevel(logger.OFF)
		}
		return fmt.Sprintf("%s succ", command), nil
	case "tars.loadconfig":
		cfg := GetServerConfig()
		remoteConf := NewRConf(cfg.App, cfg.Server, cfg.BasePath)
		_, err := remoteConf.GetConfig(cmd[1])
		if err != nil {
			return fmt.Sprintf("Getconfig Error!: %s", cmd[1]), err
		}
		return fmt.Sprintf("Getconfig Success!: %s", cmd[1]), nil

	case "tars.connection":
		return fmt.Sprintf("%s not support now!", command), nil
	default:
		if fn, ok := adminMethods[cmd[0]]; ok {
			return fn(command)
		}
		return fmt.Sprintf("%s not support now!", command), nil
	}
}

func RegisterAdmin(name string, fn adminFn) {
	adminMethods[name] = fn
}
