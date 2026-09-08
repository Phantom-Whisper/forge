Forge
=====

## Introduction

***forge** is a self-hosted Platform as a Service ([PaaS](https://wikipedia.org/wiki/Platform_as_a_service "More informations on what's a PaaS"))*. 

It allows developers to connect their Git repositories and deploy their applications on their own infrastructure.  
It aims to provide a simple and automated deployment experience while giving users full control over their applications, infrastructure and data.


## Tech Stack

<!-- language/framework used with version -->
- **backend**
- **frontend**
- **database**
- **containerization**
- **CI/CD**
- **API REST**
- **Infrastructure** 

## General Information

### What's forge?

***forge*** is designed to simplify the process of deploying and managing applications.

Instead of manually configuring servers, containers and deployment pipelines, developers can use ***forge*** to automate these processes from a single platform.

### Problem solved

Deploying an application on your own infrastructure often requires configuring multiple tools and services, such as:

- Git repositories
- Build environments
- Docker containers
- Servers
- Reverse proxies
- SSL certificates
- Deployment pipelines
- Application logs

***forge*** aims to provide a unified interface for these tasks while keeping the infrastructure under the user's control. *(and be a fun way for me to learn new skills!)*

### How does it work?

The user creates a project by connecting a Git repository and defining how the application should be built and deployed.

When a deployment is triggered, Forge creates a deployment job and assigns it to a worker. The worker builds and starts the application in a Docker container, performs a health check, and makes it available through a reverse proxy.

The deployment flow is:

GitHub → Forge API → Job Queue → Worker → Docker Container → Reverse Proxy → Application

## Get Started

### Prerequisites

In order to run ***forge***, the following softwares/frameworks are required:

<!-- frameworks & versions -->
<!-- - Other dependencies -->


### Installation

<!-- how to clone and install Forge. -->

```
git clone <github url> && cd forge
```

<!-- Add installation commands here. -->

### Configuration

<!-- Explain environment variables, configuration files and required settings. -->

### Running forge

<!-- Explain how to start Forge locally. -->

```
> command
```

### Development Environment

<!-- Explain how to run Forge in development mode and any additional configuration required. -->

## Testing

### Running Tests

***forge*** uses <!-- testing framework/library --> for automated tests
```
> test command
```

### Tests Structure

<!-- Explain where unit, integration and/or end-to-end tests are located. -->

## Known Bugs



## License

***forge*** is licensed under <!-- license name --> the license.  
See the [LICENSE]() file for more information.