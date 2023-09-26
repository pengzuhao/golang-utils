package sshremotecmd

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

type SSHCfg struct {
	RemoteAddr string
	Username   string
	Passwd     string
	Port       int
	client     *ssh.Client
}

func GetSshSession(remoteAddr, userName, passwd string, port int) (session *ssh.Session, err error) {
	var cfg = &SSHCfg{
		RemoteAddr: remoteAddr,
		Username:   userName,
		Passwd:     passwd,
		Port:       port,
	}
	config := ssh.ClientConfig{
		User:            cfg.Username,
		Auth:            []ssh.AuthMethod{ssh.Password(cfg.Passwd)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil },
		Timeout:         time.Second,
	}
	addr := fmt.Sprintf("%s:%d", cfg.RemoteAddr, cfg.Port)
	client, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		return
	}
	cfg.client = client
	session, err = cfg.client.NewSession()
	if err != nil {
		return
	}
	return session, err
}

func CmdWithOutput(remoteAddr, userName, passwd, cmd string, port int) (bufStr string, err error) {
	session, err := GetSshSession(remoteAddr, userName, passwd, port)
	if err != nil {
		return
	}
	defer session.Close()
	// err = session.Run(cmd)
	buf, err := session.CombinedOutput(cmd)
	if err != nil {
		return
	}
	bufStr = string(buf)
	return bufStr, err
}
func CmdWithOutOutput(remoteAddr, userName, passwd, cmd string, port int) (res bool, err error) {
	session, err := GetSshSession(remoteAddr, userName, passwd, port)
	if err != nil {
		return
	}
	defer session.Close()
	err = session.Run(cmd)
	if err != nil {
		return
	}
	res = true
	return res, err
}
