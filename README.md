# Docker Compose and GitHub Actions Setup for PostgreSQL and Backend Service

This repository contains a Docker Compose configuration for setting up a PostgreSQL database and a backend service using Docker.
Also it contains a GitHub Actions workflow for running Docker Compose, Logs, Tests on the backend service.

## Prerequisites

Before you begin, ensure you have the following installed on your local machine:

- Docker: [Install Docker](https://docs.docker.com/get-docker/)
- Docker Compose: [Install Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/yourusername/your-repository.git
    ```
2. Navigate to the cloned repository:

   ```bash
   cd your-repository
   ```
3. Create a .env file in the root directory and specify the following environment variables:

   ```bash
   POSTGRES_USER=your_postgres_user
   POSTGRES_PASSWORD=your_postgres_password
   POSTGRES_DB=your_postgres_db
   POSTGRES_HOST=postgres
   POSTGRES_PORT=5432
   ```
4. Build and start the Docker containers using Docker Compose:

   ```bash
   docker-compose up -d
   ```
5. Access your backend service at http://localhost:8080.

## GitHub Actions Workflow

The GitHub Actions workflow defined in [.github/workflows/ci-cd.yml](.github/workflows/ci-cd.yml) performs the following steps:

1. Sets up the Docker Compose environment.
2. Builds the Docker containers.
3. Starts the Docker containers.
4. Displays the logs.
5. Runs the tests on the backend service.

To use this workflow, you need to create a GitHub repository and push the code to the repository. The workflow will run automatically.



