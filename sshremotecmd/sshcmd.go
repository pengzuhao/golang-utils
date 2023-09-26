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

func GetSshSession(remoteAddr, userName, passwd string, port int) (session *ssh.Session) {
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
		panic(err)
	}
	cfg.client = client
	session, err = cfg.client.NewSession()
	if err != nil {
		panic(err)
	}
	return session
}

func CmdWithOutput(remoteAddr, userName, passwd, cmd string, port int) string {
	session := GetSshSession(remoteAddr, userName, passwd, port)
	defer session.Close()
	// err = session.Run(cmd)
	buf, err := session.CombinedOutput(cmd)
	if err != nil {
		panic(err)
	}
	return string(buf)
}
func CmdWithOutOutput(remoteAddr, userName, passwd, cmd string, port int) bool {
	session := GetSshSession(remoteAddr, userName, passwd, port)
	defer session.Close()
	err := session.Run(cmd)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
