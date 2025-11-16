# Hello Toby

A simple Go application that responds with "Hello Toby".

## Getting Started

This project demonstrates how to build a Go application and package it into a Docker container.

### Prerequisites

- Go 1.20 or later
- Docker

### Building the Application

To build the application locally, run the following commands:

```bash
go build -o hellotoby main.go
```

### Running the Application

You can run the application locally using:

```bash
./hellotoby
```

The application will listen on port `8080`. You can access it at `http://localhost:8080`.

### Docker

To build and run the application using Docker, follow these steps:

1. Build the Docker image:

   ```bash
   docker build -t hellotoby .
   ```

2. Run the Docker container:

   ```bash
   docker run -p 8080:8080 hellotoby
   ```

Access the application at `http://localhost:8080`.

### Continuous Integration

This project is set up with CircleCI for continuous integration. The Docker image will be automatically built and pushed to Docker Hub on each commit.

### Contributing

Feel free to submit issues or pull requests if you have suggestions or improvements.
```

### Explanation of the Sections

- **Project Title**: Clearly states the name of the project.
- **Getting Started**: Provides a brief overview of what the project does.
- **Prerequisites**: Lists the necessary tools to run the project.
- **Building the Application**: Instructions for building the application locally.
- **Running the Application**: Instructions for running the application locally.
- **Docker**: Instructions for building and running the application using Docker.
- **Continuous Integration**: Mentions the CircleCI setup for automated builds.
- **Contributing**: Encourages contributions and feedback.

This `README.md` provides a clear and concise overview of the project, making it easy for others to understand how to use and contribute to it.


