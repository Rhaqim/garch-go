package config

type ProjectDependencies struct {
	API     []string
	Web     []string
	CLI     []string
	Lib     []string
	Service []string
	Other   []string
}

var Deps = ProjectDependencies{
	API:     []string{"echo", "gin", "fiber", "chi"},
	Web:     []string{"react", "vue", "angular", "svelte"},
	CLI:     []string{"cobra", "cli"},
	Lib:     []string{"ginkgo", "gomega", "testify"},
	Service: []string{"grpc", "nats", "rabbitmq"},
	Other:   []string{},
}
