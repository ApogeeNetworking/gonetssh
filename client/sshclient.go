package client

import (
	"fmt"
	"io"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

var (
	// Additional Ciphers that may be needed for old out-of-date Devices
	sshCiphers = []string{
		"aes256-ctr",
		"aes128-ctr",
		"aes128-cbc",
		"3des-cbc",
		"aes192-ctr",
		"aes192-cbc",
		"aes256-cbc",
		"aes128-gcm@openssh.com",
	}
	// Additional KeyExchanges that may be needed for old out-of-date Devices
	sshKeyExchanges = []string{
		"diffie-hellman-group-exchange-sha1",
		"diffie-hellman-group1-sha1",
		"diffie-hellman-group14-sha1", // Aruba SSH
		"ecdh-sha2-nistp256",
		"ecdh-sha2-nistp384",
		"ecdh-sha2-nistp521",
	}
)

// SSH ...
type SSH struct {
	host    string
	user    string
	pass    string
	client  *ssh.Client
	session *ssh.Session
	input   chan *string
	stop    chan struct{}
	prompt  string
	reader  io.Reader
	writer  io.WriteCloser
	Enable  string
}

// NewSSH ...
func NewSSH(host, user, pass string) (*SSH, error) {
	return &SSH{
		host: host,
		user: user,
		pass: pass,
	}, nil
}

// NewClient connects ssh client
func (s *SSH) NewClient(cfg *ssh.ClientConfig) (*ssh.Client, error) {
	return ssh.Dial("tcp", s.host+":22", cfg)
}

func (s *SSH) keyInter(u, in string, q []string, e []bool) ([]string, error) {
	// Just send the password back for all questions
	answers := make([]string, len(q))
	for i := range answers {
		answers[i] = s.pass
	}

	return answers, nil
}

func (s *SSH) getPassword() (string, error) {
	return s.pass, nil
}

// NewClientConfig setups an ssh client config struct
func (s *SSH) NewClientConfig() *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: s.user,
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(s.keyInter),
			ssh.Password(s.pass),
			ssh.PasswordCallback(s.getPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}
}

// Connect establishes socket with Device with retries
func (s *SSH) Connect(retries int) error {
	cfg := s.NewClientConfig()
	// Attach Additional Ciphers and KeyExchanges
	cfg.Ciphers = append(cfg.Ciphers, sshCiphers...)
	cfg.KeyExchanges = append(cfg.KeyExchanges, sshKeyExchanges...)
	sshClient, err := s.NewClient(cfg)
	if err != nil {
		// Before we give up on a failed handshake
		if strings.Contains(err.Error(), "handshake") {
			count := retries - 1
			if count == 0 {
				return fmt.Errorf("all ssh conn retries exhausted: %v", err)
			}
			return s.Connect(count)
		}
		return fmt.Errorf("error establishing ssh connection: %v", err)
	}
	s.client = sshClient
	// Create new SSH Session
	sshSession, err := s.client.NewSession()
	if err != nil {
		// Close the SSH Client Connection
		s.client.Conn.Close()
		return fmt.Errorf("error setting up ssh session: %v", err)
	}
	s.session = sshSession
	// Setup Terminal Modes on the SSH Session (to remember if we are exec|config|et al)
	modes := ssh.TerminalModes{
		ssh.ECHO:          0, // disable echoing
		ssh.OCRNL:         0,
		ssh.TTY_OP_ISPEED: 38400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 38400, // output speed = 14.4kbaud
	}
	if err := s.session.RequestPty("xterm", 0, 500, modes); err != nil {
		// Close the Session
		s.session.Close()
		// Close the Connection
		s.client.Conn.Close()
		return fmt.Errorf("req for pseudo terminal failed: %v", err)
	}
	s.reader, _ = s.session.StdoutPipe()
	s.writer, _ = s.session.StdinPipe()
	// Start SSH Session Login Shell
	if err := s.session.Shell(); err != nil {
		s.Disconnect()
		return fmt.Errorf("unable to start ssh session shell: %v", err)
	}
	return nil
}

// Disconnect ...
func (s *SSH) Disconnect() {
	if s.client != nil {
		s.client.Conn.Close()
	}
	if s.session != nil {
		s.session.Close()
	}
}

// Read ...
func (s *SSH) Read() (string, error) {
	buf := make([]byte, 204800000)
	n, err := s.reader.Read(buf)
	return string(buf[:n]), err
}

// Write ...
func (s *SSH) Write(cmd string) int {
	n, _ := s.writer.Write([]byte(cmd))
	time.Sleep(100 * time.Millisecond)
	return n
}

// ExecEnable ...
func (s *SSH) ExecEnable(pass string) {
	s.writer.Write([]byte("enable\n"))
	time.Sleep(100 * time.Millisecond)
	s.writer.Write([]byte(pass + "\n"))
}
