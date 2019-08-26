package zmage

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Docker mg.Namespace

// Build docker image.
func (Docker) Build() error {
	color.Red("# Docker -------------------------------------------------------------------")
	fmt.Printf("BUILD_DATE : %s\n", time.Now().Format(time.RFC3339))
	fmt.Printf("VERSION : %s\n", tag())
	fmt.Printf("VCS_REF : %s\n", hash())

	fmt.Printf(" > Production image\n")
	return sh.RunV("docker", "build",
		"-f", "deployment/docker/Dockerfile",
		"--build-arg", fmt.Sprintf("BUILD_DATE=%s", time.Now().Format(time.RFC3339)),
		"--build-arg", fmt.Sprintf("VERSION=%s", tag()),
		"--build-arg", fmt.Sprintf("VCS_REF=%s", hash()),
		"-t", "miam:latest",
		".")
}
