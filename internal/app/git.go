package app

import (
	"fmt"
	"net"
	"os"

	"github.com/go-git/go-git/v5"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func GitClone(dir string, repo string) error {
	fmt.Printf(" - %v \n", repo)

	_ = handleSshKey()

	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL: repo,
	})

	if err != nil {
		return err
	}

	return nil
}

func handleSshKey() ssh.ClientConfig, error {

	socket := os.Getenv("SSH_AUTH_SOCK")
	user := os.Getenv("USER")

	conn, err := net.Dial("unix", socket)
	if err != nil {
		return _, fmt.Errorf("failed to open SSH_AUTH_SOCK: %v", err)
	}

	agentClient := agent.NewClient(conn)
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(agentClient.Signers),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return config
}
