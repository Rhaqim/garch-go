package domain

var Deps = ProjectDependencies{
	API:     []string{"echo", "gin", "fiber", "chi"},
	CLI:     []string{"cobra", "cli"},
	Lib:     []string{"ginkgo", "gomega", "testify"},
	Service: []string{"grpc", "nats", "rabbitmq"},
	Other:   []string{},
}

var DepsStringMap = map[string][]string{
	"API":     Deps.API,
	"CLI":     Deps.CLI,
	"Lib":     Deps.Lib,
	"Service": Deps.Service,
	"Other":   Deps.Other,
}

var DependencyMap = map[string]string{
	"echo":     "github.com/labstack/echo/v4",
	"gin":      "github.com/gin-gonic/gin",
	"fiber":    "github.com/gofiber/fiber/v2",
	"chi":      "github.com/go-chi/chi",
	"cobra":    "github.com/spf13/cobra",
	"cli":      "github.com/urfave/cli/v2",
	"ginkgo":   "github.com/onsi/ginkgo",
	"gomega":   "github.com/onsi/gomega",
	"testify":  "github.com/stretchr/testify",
	"grpc":     "google.golang.org/grpc",
	"nats":     "github.com/nats-io/nats.go",
	"rabbitmq": "github.com/streadway/amqp",
}
