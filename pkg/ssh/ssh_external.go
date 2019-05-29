/*
Copyright 2018 The Doctl Authors All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ssh

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func runExternalSSH(r *Runner) error {
	args := []string{}
	if r.KeyPath != "" {
		args = append(args, "-i", r.KeyPath)
	}

	if r.KnownHostsPath != "" {
		args = append(args, "-o", fmt.Sprintf("UserKnownHostsFile=%s", r.KnownHostsPath))
	}

	sshHost := r.Host
	if r.User != "" {
		sshHost = r.User + "@" + sshHost
	}

	if r.Port > 0 {
		args = append(args, "-p", strconv.Itoa(r.Port))
	}

	if r.AgentForwarding {
		args = append(args, "-A")
	}

	args = append(args, sshHost)
	if r.Command != "" {
		args = append(args, r.Command)
	}

	cmd := exec.Command("ssh", args...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
