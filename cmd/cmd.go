package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Rsync uses rsync to pus or pull a folder from client to host
func Rsync(host, user string, port int, local, remote string, push bool) error {
	//  rsync --update -raz -e "ssh -p 28868" --progress <local> <remote>
	// -r recursize
	// -E perserve execuatable bit
	// -a archive mode
	// -e custom shell (ssh options)
	// -z use compression
	// -u update
	// --delete delete unsued files.

	args := []string{
		"-rEazu",
		"--delete",
	}

	if port != 0 {
		args = append(args, fmt.Sprintf(`-e ssh -p %d`, port))
	}

	remotePath := fmt.Sprintf("%s:%s", host, remote)

	if push {
		args = append(args, local, remotePath)
	} else {
		args = append(args, remotePath, local)
	}

	fmt.Printf("syncing ->  rsync %s\n", strings.Join(args, " "))

	cmd := exec.Command("rsync", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Cmd runs a command on a remote computer over ssh.
func Cmd(host, user string, port int, remote string, rCmd string) error {
	// ssh -t -p <port> user@host bash -c 'cd <remote> && cmd'
	args := []string{"-t"}
	if port != -1 {
		args = append(args, fmt.Sprintf("-p %d", port))
	}

	args = append(args, []string{
		fmt.Sprintf("%s@%s", user, host),
		"bash",
		"-c",
		fmt.Sprintf("'cd %s && %s'", remote, rCmd),
	}...)

	fmt.Printf("running -> ssh %s\n", strings.Join(args, " "))

	cmd := exec.Command("ssh", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
