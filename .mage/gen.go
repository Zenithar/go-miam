package zmage

import (
	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Gen mg.Namespace

// Generate initializers
func (Gen) Wire() {
	color.Blue("### Wiring dispatchers")
	// mustGoGenerate("gRPC", "go.zenithar.org/spotigraph/cli/spotigraph/internal/dispatchers/grpc")
}

// Generate mocks for tests
func (Gen) Mocks() {
	color.Blue("### Mocks")

	// mustGoGenerate("Repositories", "go.zenithar.org/spotigraph/internal/repositories")
	// mustGoGenerate("Services", "go.zenithar.org/spotigraph/internal/services")
}

// Generate mocks for tests
func (Gen) Decorators() {
	color.Blue("### Decorators")

	// mustGoGenerate("Repositories", "go.zenithar.org/spotigraph/internal/repositories/pkg/...")
	// mustGoGenerate("Services", "go.zenithar.org/spotigraph/internal/services/pkg/...")
}

// Generate initializers
func (Gen) Migrations() {
	color.Blue("### Database migrations")

	// mustGoGenerate("PostgreSQL", "go.zenithar.org/spotigraph/internal/repositories/pkg/postgresql")
}

// Generate protobuf
func (Gen) Protobuf() error {
	color.Blue("### Protobuf")

	return sh.RunV("prototool", "all", "--fix", "pkg/protocol")
}
