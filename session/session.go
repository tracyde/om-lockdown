// Package session provides functions to support operating on OpsManager
package session

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

// Session stores information about the opsman we are configuring
type Session struct {
	Addr     string
	Username string
	Password string
	config   *ssh.ClientConfig
}

// New creates a new ssh session to OpsManager using provided ssh config
func New(hostname, username, password string, config *ssh.ClientConfig) *Session {
	return &Session{Addr: fmt.Sprintf("%s:%s", hostname, "22"), config: config}
}

// New creates a new ssh session to OpsManager using generic ssh config
func NewGeneric(hostname, username, password string) *Session {
	config := &ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return &Session{
		Addr:     fmt.Sprintf("%s:%s", hostname, "22"),
		Username: username,
		Password: password,
		config:   config}
}

func (s *Session) newSession() (*ssh.Session, error) {
	conn, err := ssh.Dial("tcp", s.Addr, s.config)
	if err != nil {
		return nil, err
	}

	session, err := conn.NewSession()
	if err != nil {
		return nil, err
	}

	// Globally set HISTIGNORE, seems like a good idea to
	// ignore sudo commands that contain passwords
	session.Setenv("HISTIGNORE", "‘*sudo -S*’")
	return session, nil
}

// ExecuteCmd runs a command on remote opsman and returns results
func (s *Session) ExecuteCmd(command string) (string, error) {
	session, err := s.newSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)

	return string(output), err
}
