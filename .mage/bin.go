package zmage

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Bin mg.Namespace

func (Bin) Miam() error {
	return goBuild("go.zenithar.org/miam/cli/miam", "miam")
}

func goBuild(packageName, out string) error {
	fmt.Printf(" > Building %s [%s]\n", out, packageName)

	varsSetByLinker := map[string]string{
		"go.zenithar.org/miam/internal/version.Version":   tag(),
		"go.zenithar.org/miam/internal/version.Revision":  hash(),
		"go.zenithar.org/miam/internal/version.Branch":    branch(),
		"go.zenithar.org/miam/internal/version.BuildUser": os.Getenv("USER"),
		"go.zenithar.org/miam/internal/version.BuildDate": time.Now().Format(time.RFC3339),
		"go.zenithar.org/miam/internal/version.GoVersion": runtime.Version(),
	}
	var linkerArgs []string
	for name, value := range varsSetByLinker {
		linkerArgs = append(linkerArgs, "-X", fmt.Sprintf("%s=%s", name, value))
	}
	linkerArgs = append(linkerArgs, "-s", "-w")

	return sh.RunWith(map[string]string{
		"CGO_ENABLED": "0",
	}, "go", "build", "-ldflags", strings.Join(linkerArgs, " "), "-mod=vendor", "-o", fmt.Sprintf("bin/%s", out), packageName)
}

// -----------------------------------------------------------------------------
