package deploy

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func newSession() {

	if serverConfig.SSHPort == "" {
		serverConfig.SSHPort = "22"
	}
	hostKey, err := getHostKey(fmt.Sprintf("%s", serverConfig.SSHHost))
	if err != nil {
		log.Fatal(err)
	}
	key, err := ioutil.ReadFile(localConfig.PrivateKey)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("private key: %v", err)
	}
	config := &ssh.ClientConfig{
		User: serverConfig.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	client, _ = ssh.Dial("tcp", fmt.Sprintf("%s:%s", serverConfig.SSHHost, serverConfig.SSHPort), config)
	session, _ = client.NewSession()
}


func getHostKey(host string) (ssh.PublicKey, error) {
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				return nil, fmt.Errorf("error parsing %q: %v", fields[2], err)
			}
			break
		}
	}
	if hostKey == nil {
		return nil, fmt.Errorf("no hostkey for %s", host)
	}
	return hostKey, nil
}


func execCmd(s string) error{
	ns, _ := client.NewSession()
	defer ns.Close()
	ns.Stdout = os.Stdout
	ns.Stderr = os.Stderr
	return ns.Run(s)
}
