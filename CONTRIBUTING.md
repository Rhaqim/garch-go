# GarchGo CLI Application

This is a simple CLI application written in Go that demonstrates the use of the hexagonal architecture pattern. The application provides a command-line interface for generating a Go project with a predefined directory structure and architecture.

## Getting Started

To run the application, you need to have Go installed on your machine. You can download and install Go from the official website: [https://golang.org/dl/](https://golang.org/dl/).

Once you have Go installed, you can clone this repository and run the application using the following commands:

```bash
git clone https://github.com/Rhaqim/garchgo.git
cd garchgo
```

### Running the Application

To run the application, use the following command:

```bash
go run cmd/gen/main.go

# or

go build cmd/gen/main.go
./main
```

### Building the Application

To build the application, use the following command:

```bash
go build cmd/gen/main.go

# or with make
make build
```

This will build and run the application, displaying the available commands and options. You can then use the application to generate a Go project with a predefined directory structure and architecture.

## Project Structure

```bash
GaarchGo/
│
├── cmd/
│   └── gen/
│       └── main.go               # Entry point for the 'gen' command
│
├── internal/
│   ├── adapter/                  # Adapters for CLI interaction
│   │   └── cli/                  # CLI adapter
│   │       └── cli.go            # Implementation for CLI interaction
│   │
│   ├── app/                      # Application core
│   │   ├── core/                 # Domain layer
│   │   │   └── commands.go       # commands ran on the terminal
│   │   │   └── core.go           # contains core functionalty
│   │   ├── domain/               # Domain layer
│   │   │   └── model.go          # Define domain models/interfaces
│   │   ├── service/              # Business logic layer
│   │   │   └── service.go        # Define application services
│   │   └── usecase/              # Application use cases
│   │       └── usecase.go        # Define use case interfaces
│   │
│   ├── infrastructure/           # Infrastructure layer
│   │   └── architectures/        # Stub architecture for different architectures
│   │       └── architecture.go   # Holds the root files and folders for the architectures
│   │   └── repository/           # Stub repository for future use
│   │       └── repository.go     # Interface for project repository
│   │       └── root.go           # Contains the root files and similar folders for the architectures
│   │
│   └── utils/                    # Utilities package
│       └── utils.go              # Utility functions/helpers
│
└── config/                       # Configuration files
    └── config.go                 # Configuration loading logic
```

### Explanation

- **`cmd/`**: This directory contains the main executable for the CLI application. It's responsible for parsing command-line arguments and invoking the appropriate functionality from the application core.

- **`internal/`**: This directory contains the internal packages of the application. These packages should not be imported from outside the module. This convention ensures encapsulation and prevents accidental dependencies.

  - **`app/`**: This directory contains the core logic of the application, including the domain, business logic, and use cases.

    - **`core/`**: Contains the core functionality of the application, including domain logic, business logic, and use cases.

      - **`commands.go`**: Contains the commands that can be executed from the command-line interface.

      - **`core.go`**: Contains the core functionality of the application, including domain logic, business logic, and use cases.

    - **`domain/`**: Contains the domain models and interfaces that define the core concepts of the application.

      - **`model.go`**: Contains the domain models and interfaces that define the core concepts of the application.

    - **`service/`**: Contains the application service, which encapsulates the business logic of the application.

      - **`service.go`**: Contains the application service, which encapsulates the business logic of the application.

    - **`usecase/`**: Contains the application use cases, which define the high-level interactions with the application.

      - **`usecase.go`**: Contains the application use cases, which define the high-level interactions with the application.
  
  - **`adapter/`**: This directory contains implementations of interfaces defined in the `app` package, such as CLI adapters or web API handlers.

    - **`cli/`**: Contains the CLI adapter, which is responsible for interacting with the command-line interface.

      - **`cli.go`**: Contains the implementation of the CLI adapter, including parsing command-line arguments and invoking the appropriate use cases from the application core.

  - **`infrastructure/`**: This directory contains implementations of interfaces defined in the `app` package, such as database repositories or external API clients.

    - **`architectures/`**: Contains the stub architecture for different architectures. This directory can be extended to include different architectures, such as MVC, MVP, or MVVM.

      - **`architecture.go`**: Holds the root files and folders for the architectures.

    - **`repository/`**: Contains the repository interface and its implementations. In this example, there is a stub repository that doesn't interact with a real database or external service.
  
      - **`repository.go`**: Contains the repository interface, which defines the methods for interacting with the project repository.

      - **`root.go`**: Contains the root files and similar folders for the architectures.
  
  - **`utils/`**: Contains utility functions or helpers that are used throughout the application.

    - **`utils.go`**: Contains utility functions or helpers that are used throughout the application.

- **`config/`**: Contains configuration files and logic for loading configuration settings into the application.

  - **`config.go`**: Contains the logic for loading configuration settings from environment variables, configuration files, or other sources.

This structure separates concerns clearly, making it easier to maintain and extend the application. Each layer has well-defined responsibilities, and dependencies flow inward toward the core, adhering to the principles of the hexagonal architecture.
