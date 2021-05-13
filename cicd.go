package cicd

import (
	"log"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func CloneMaster(url string, sshKey string) {

	_, err := os.Stat(sshKey)
	if err != nil {
		log.Println("Error while reading SSH key")
		return
	}
	publicKeys, err := ssh.NewPublicKeysFromFile("git", sshKey, "")
	if err != nil {
		log.Println("Error while creating public key from file")
		return
	}

	directory := "/tmp/repo"
	r, err = git.PlainClone(directory, false, &git.CloneOptions{
		Auth:     publicKeys,
		URL:      url,
		Progress: os.Stdout,
	})

}
