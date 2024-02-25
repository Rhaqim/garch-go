# GarchGo CLI Application

Garch is a command-line interface (CLI) application built with Go. It is used to generate Go boilerplate code for different projects and depending on the architecture, different folder structures are generated.

## Table of Contents

- [GarchGo CLI Application](#garchgo-cli-application)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Architectures](#architectures)
    - [Hexagonal Architecture](#hexagonal-architecture)
      - [Explanation of each directory (Hexagonal Architecture)](#explanation-of-each-directory-hexagonal-architecture)
    - [Clean Architecture](#clean-architecture)
      - [Explanation of each directory (Clean Architecture)](#explanation-of-each-directory-clean-architecture)
    - [Onion Architecture](#onion-architecture)
      - [Explanation of each directory (Onion Architecture)](#explanation-of-each-directory-onion-architecture)
    - [Domain Driven Design](#domain-driven-design)
      - [Explanation of each directory (Domain Driven Design)](#explanation-of-each-directory-domain-driven-design)
    - [Comparison of Architectures](#comparison-of-architectures)
  - [License](#license)
  - [Acknowledgements](#acknowledgements)
  - [Contributing](#contributing)

## Installation

To install Garch, you need to have Rust and Cargo installed on your system. Once you have them set up, you can install Garch by running the following command:

- via Go:

```bash
go get -u github.com/Rhaqim/garch-go
```

- via Git:

```bash
git clone https://github.com/Rhaqim/garch-go.git
cd garch-go
make install
```

- executeable -  the executeable will be located in the `root` directory of the project.

```bash
./garch-go --help
```

## Usage

To use Garch, you need to run the following command:

```bash
garch --help
```

This will display the help message, which will show you how to use Garch.

## Architectures

Garch supports the following architectures:

- [Hexagonal Architecture](#hexagonal-architecture)
- [Clean Architecture](#clean-architecture)
- [Onion Architecture](#onion-architecture)
- [Domain Driven Design](#domain-driven-design)

### [Hexagonal Architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software))

In a hexagonal architecture, the focus is on defining your application's core domain logic and isolating it from external concerns such as frameworks or databases. Here's a file structure for a project following the hexagonal architecture:

```bash
project/
├── cmd/
│   └── yourapp/
│       └── main.go             # Entry point of your application
│
├── pkg/                        # Shared packages
│
├── internal/
│   ├── app/                    # Application core
│   │   ├── entity/             # Domain entities
│   │   ├── repository/         # Interface and implementation of repositories
│   │   ├── usecase/            # Use cases (business logic)
│   │   └── service/            # Application services
│   │
│   ├── adapter/                # Adapters to connect ports to the external world
│   │   ├── delivery/           # Delivery adapters (e.g., HTTP, gRPC)
│   │   ├── persistence/        # Persistence adapters (e.g., database)
│   │   └── external/           # External service adapters (e.g., email, payment gateway)
│   │
│   └── config/                 # Configuration files
│
└── scripts/                    # Utility scripts (optional)
```

#### Explanation of each directory (Hexagonal Architecture)

- `cmd/`: Contains the entry point of your application. You may have multiple directories here if your project has multiple executables.
- `pkg/`: Contains shared packages that can be used across different parts of your project.
- `internal/`: Contains the main implementation of your application.
  - `app/`: Contains the core application logic.
    - `entity/`: Contains domain entities representing your business objects.
    - `repository/`: Contains interfaces defining repository contracts and their implementations.
    - `usecase/`: Contains use cases or interactors, which encapsulate business rules.
    - `service/`: Contains application services.
  - `adapter/`: Contains adapters to connect ports to the external world.
    - `delivery/`: Contains delivery adapters (e.g., HTTP, gRPC).
    - `persistence/`: Contains persistence adapters (e.g., database).
    - `external/`: Contains adapters for interacting with external services.
  - `config/`: Contains configuration files for your application.
- `scripts/`: Contains utility scripts that can be used during development or deployment (optional).

This structure supports the principles of hexagonal architecture by clearly separating the core domain logic from external concerns. The core logic resides in the `app` directory, while adapters in the `adapter` directory bridge the core logic with the external world. This allows for easier testing, maintenance, and evolution of your application.

### Clean Architecture

In a clean architecture, you typically organize your code into layers based on their level of abstraction and dependency direction. Here's a common file structure for a project following clean architecture principles:

```bash
project/
├── cmd/
│   └── yourapp/
│       └── main.go             # Entry point of your application
│
├── internal/
│   ├── app/                    # Application core
│   │   ├── entity/             # Domain entities
│   │   ├── repository/         # Interface and implementation of repositories
│   │   ├── usecase/            # Use cases (business logic)
│   │   └── service/            # Interface and implementation of services
│   │
│   ├── delivery/               # Delivery mechanisms (UI, API)
│   │   ├── handler/            # HTTP handlers/controllers
│   │   └── presenter/          # Response formatting/presentation logic
│   │
│   └── infrastructure/         # External interfaces implementations
│       ├── persistence/        # Database access, ORM models
│       └── external/           # External services (e.g., email, payment gateway)
│
├── pkg/                        # Shared packages (optional)
│
└── configs/                    # Configuration files (optional)
```

#### Explanation of each directory (Clean Architecture)

- `cmd/`: Contains the entry point of your application. You may have multiple directories here if your project has multiple executables.
- `internal/`: Contains the main implementation of your application.
  - `app/`: Contains the core application logic.
    - `entity/`: Contains domain entities representing your business objects.
    - `repository/`: Contains interfaces defining repository contracts and their implementations.
    - `usecase/`: Contains use cases or interactors, which encapsulate business rules.
    - `service/`: Contains interfaces defining service contracts and their implementations.
  - `delivery/`: Contains the delivery mechanisms of your application (e.g., web controllers, API handlers).
    - `handler/`: Contains HTTP handlers or controllers.
    - `presenter/`: Contains logic for formatting and presenting responses.
  - `infrastructure/`: Contains implementations of external interfaces (e.g., database access, external services).
    - `persistence/`: Contains database access logic and ORM models.
    - `external/`: Contains implementations for interacting with external services.
- `pkg/`: Contains shared packages that can be used across different parts of your project (optional).
- `configs/`: Contains configuration files for your application (optional).

This structure helps in separating concerns and keeping the codebase organized. It also supports flexibility and maintainability by allowing components to be replaced or modified without affecting other parts of the system.

### Onion Architecture

In Onion Architecture, the focus is on organizing your application around its core domain logic, with layers representing increasing levels of abstraction and dependency direction. Here's a file structure for a project following the Onion Architecture:

```bash
project/
├── cmd/
│   └── yourapp/
│       └── main.go             # Entry point of your application
│
├── pkg/                        # Shared packages (optional)
│
├── internal/
│   ├── app/                    # Application core
│   │   ├── entity/             # Domain entities
│   │   ├── repository/         # Interface and implementation of repositories
│   │   ├── usecase/            # Use cases (business logic)
│   │   └── service/            # Application services
│   │
│   ├── infrastructure/         # Infrastructure layer
│   │   ├── persistence/        # Database access, ORM models
│   │   └── external/           # External services (e.g., email, payment gateway)
│   │
│   └── interfaces/             # Interface adapters
│       ├── delivery/           # Delivery mechanisms (e.g., HTTP, gRPC)
│       └── persistence/        # Persistence adapters (e.g., database)
│
└── configs/                    # Configuration files (optional)
```

#### Explanation of each directory (Onion Architecture)

- `cmd/`: Contains the entry point of your application. You may have multiple directories here if your project has multiple executables.
- `pkg/`: Contains shared packages that can be used across different parts of your project (optional).
- `internal/`: Contains the main implementation of your application.
  - `app/`: Contains the core application logic.
    - `entity/`: Contains domain entities representing your business objects.
    - `repository/`: Contains interfaces defining repository contracts and their implementations.
    - `usecase/`: Contains use cases or interactors, which encapsulate business rules.
    - `service/`: Contains application services.
  - `infrastructure/`: Contains implementations of external interfaces, such as database access and external services.
    - `persistence/`: Contains database access logic and ORM models.
    - `external/`: Contains implementations for interacting with external services.
  - `interfaces/`: Contains interface adapters that bridge the application core with external concerns.
    - `delivery/`: Contains delivery mechanisms such as HTTP, gRPC, etc.
    - `persistence/`: Contains persistence adapters for interacting with the database.
- `configs/`: Contains configuration files for your application (optional).

This structure supports the principles of Onion Architecture by organizing the codebase around the core domain logic and ensuring that dependencies flow inward, with the inner layers no

### Domain Driven Design

In Domain-Driven Design (DDD), the focus is on modeling your software around the core domain, using a layered architecture to separate concerns and ensure a clear separation of responsibilities. Here's a file structure for a project following the DDD architecture:

```bash
project/
├── cmd/
│   └── yourapp/
│       └── main.go             # Entry point of your application
│
├── pkg/                        # Shared packages (optional)
│
├── internal/
│   ├── domain/                 # Domain layer
│   │   ├── entity/             # Domain entities
│   │   ├── repository/         # Interface and implementation of repositories
│   │   └── service/            # Domain services
│   │
│   ├── application/            # Application layer
│   │   ├── dto/                # Data transfer objects
│   │   └── service/            # Application services
│   │
│   ├── infrastructure/         # Infrastructure layer
│   │   ├── persistence/        # Database access, ORM models
│   │   └── external/           # External services (e.g., email, payment gateway)
│   │
│   └── interfaces/             # Interface layer
│       ├── delivery/           # Delivery mechanisms (e.g., HTTP, gRPC)
│       └── persistence/        # Persistence adapters (e.g., database)
│
└── configs/                    # Configuration files (optional)
```

#### Explanation of each directory (Domain Driven Design)

- `cmd/`: Contains the entry point of your application. You may have multiple directories here if your project has multiple executables.
- `pkg/`: Contains shared packages that can be used across different parts of your project (optional).
- `internal/`: Contains the main implementation of your application.
  - `domain/`: Contains the domain layer, including domain entities, repositories, and domain services.
    - `entity/`: Contains domain entities representing your business objects.
    - `repository/`: Contains interfaces defining repository contracts and their implementations.
    - `service/`: Contains domain services that encapsulate domain logic.
  - `application/`: Contains the application layer, which coordinates domain entities and services to fulfill application use cases.
    - `dto/`: Contains data transfer objects used to transfer data between layers.
    - `service/`: Contains application services that orchestrate interactions between domain entities and services.
  - `infrastructure/`: Contains the infrastructure layer, including implementations of external interfaces such as database access and external services.
    - `persistence/`: Contains database access logic and ORM models.
    - `external/`: Contains implementations for interacting with external services.
  - `interfaces/`: Contains the interface layer, which adapts the application to the external world.
    - `delivery/`: Contains delivery mechanisms such as HTTP, gRPC, etc.
    - `persistence/`: Contains persistence adapters for interacting with the database.
- `configs/`: Contains configuration files for your application (optional).

This structure supports the principles of Domain-Driven Design by organizing the codebase around the domain model and ensuring clear separation of concerns. It allows for better maintainability, testability, and scalability of your application.

### Comparison of Architectures

1. **Clean Architecture**:
   - **Key Principle**: Separation of concerns and dependency rule.
   - **Layers**: Divides the application into concentric circles, with the innermost circle representing the core domain logic and each subsequent circle representing outer layers such as application, infrastructure, and interfaces.
   - **Dependency Direction**: Dependencies point inward toward the core domain logic.
   - **Main Focus**: Emphasizes organizing codebase around business logic, with clear separation of concerns and independence from external concerns.

2. **Hexagonal Architecture**:
   - **Key Principle**: Ports and adapters.
   - **Layers**: Focuses on ports (interfaces) and adapters (implementations), separating application logic from external concerns.
   - **Dependency Direction**: Application logic is at the center, surrounded by ports representing interactions with external systems, and adapters translating those interactions.
   - **Main Focus**: Emphasizes defining clear boundaries between internal application logic and external dependencies, enabling easier testing and adaptability.

3. **Onion Architecture**:
   - **Key Principle**: Separation of concerns with layers of abstraction.
   - **Layers**: Organizes the codebase into layers, with the innermost layer representing the core domain logic, followed by layers representing application, infrastructure, and interfaces.
   - **Dependency Direction**: Dependencies flow inward toward the core domain logic, with each layer being unaware of outer layers.
   - **Main Focus**: Focuses on organizing codebase around the core domain logic and ensuring that dependencies point inward, promoting modularity and maintainability.

4. **Domain-Driven Design (DDD)**:
   - **Key Principle**: Focus on the domain model and ubiquitous language.
   - **Layers**: Typically includes domain, application, infrastructure, and interface layers, but emphasizes the importance of the domain layer and modeling the core domain logic.
   - **Dependency Direction**: Emphasizes modeling the domain logic using entities, value objects, and domain services, with other layers serving to support and interact with the domain layer.
   - **Main Focus**: Centers around understanding and modeling the core domain logic using a shared ubiquitous language, with other layers supporting and facilitating interactions with the domain model.

| Architecture          | Focus                                                                                     | Key Principles                                                                                   | Pros                                                                                                 | Cons                                                                                                 |
|-----------------------|-------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------|
| Hexagonal Architecture | Defining core domain logic and isolating it from external concerns                       | Core domain logic, adapters to connect ports to the external world                                | Clear separation of concerns, easier testing, maintenance, and evolution of your application        | May introduce additional complexity, especially for smaller projects                                    |
| Clean Architecture     | Organizing code into layers based on their level of abstraction and dependency direction | Separation of concerns, flexibility, maintainability, allowing components to be replaced or modified | Clear separation of concerns, flexibility, maintainability, allowing components to be replaced or modified | May introduce additional complexity, especially for smaller projects                                    |
| Onion Architecture     | Organizing application around its core domain logic, with layers representing increasing levels of abstraction and dependency direction | Organizing code around the core domain logic, ensuring that dependencies flow inward | Organizing code around the core domain logic, ensuring that dependencies flow inward | May introduce additional complexity, especially for smaller projects                                    |
| Domain-Driven Design   | Modeling software around the core domain, using a layered architecture to separate concerns | Modeling software around the core domain, using a layered architecture to separate concerns | Clear separation of concerns, better maintainability, testability, and scalability of your application | May introduce additional complexity, especially for smaller projects                                    |

In summary, while all four architectures emphasize separation of concerns and organizing codebase around the core domain logic, they differ in their specific approaches, focus, and principles. Clean Architecture focuses on dependency rule, Hexagonal Architecture on ports and adapters, Onion Architecture on separation of concerns with layers of abstraction, and Domain-Driven Design on modeling the domain logic and ubiquitous language. Each architecture has its strengths and can be chosen based on the specific requirements and goals of the project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Rust Language](https://www.rust-lang.org/)
- [Rustup](https://rustup.rs/)
- [Cargo](https://doc.rust-lang.org/cargo/)
- [Clap](https://clap.rs/)
- [Rust Community](https://www.rust-lang.org/community)
<!-- - [Askama](
- [Handlebars](
- [Rustfmt](
- [Rustdoc]( -->

## Contributing

If you would like to contribute to this project, please read the [CONTRIBUTING.md](CONTRIBUTING.md) file for details on our code of conduct, and the process for submitting pull requests to us.
