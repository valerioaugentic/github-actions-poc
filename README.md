# GitHub Actions POC - Go API

REST API in Go with CI/CD using GitHub Actions, Docker image builds and multi-environment deployments to Fly.io.

## Project Structure

```
github-actions-poc/
├── .github/
│   └── workflows/
│       ├── deploy.yml
│       ├── deploy-qa.yml
│       ├── deploy-staging.yml
│       ├── deploy-production.yml
│       ├── develop.yml
│       └── release.yml
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   └── handlers/
│       └── environment.go
├── .env.example
├── .gitignore
├── Dockerfile
├── fly.toml
├── go.mod
└── README.md
```

## API

### GET /environment

Returns current environment information.

**Response:**
```json
{
  "environment": "development",
  "version": "1.0.0"
}
```

## URLs

| Environment | URL |
|-------------|-----|
| development | https://api-environment-development.fly.dev/environment |
| qa | https://api-environment-qa.fly.dev/environment |
| staging | https://api-environment-staging.fly.dev/environment |
| production | https://api-environment-production.fly.dev/environment |

## Local Development

### Prerequisites
- Go 1.21+
- Docker (optional)

### Run locally

```bash
cp .env.example .env
go run cmd/api/main.go
curl http://localhost:8080/environment
```

### Run with Docker

```bash
docker build -t api-environment:local .
docker run -p 8080:8080 -e ENVIRONMENT=development -e VERSION=1.0.0 api-environment:local
curl http://localhost:8080/environment
```

## CI/CD

### Workflow: develop.yml

**Trigger:** Push to `develop` branch

- Build Docker image
- Push to `ghcr.io/<owner>/api-environment:develop`

### Workflow: release.yml

**Trigger:** Tag `v*` creation on `main` branch

- Build and push image with tag version
- Automatic deploy to development

### Manual Deploys

| Workflow | Environment |
|----------|-------------|
| deploy-qa.yml | qa |
| deploy-staging.yml | staging |
| deploy-production.yml | production |

To trigger: **Actions > Select workflow > Run workflow > Enter version**

## GitHub Configuration

### 1. Environments

In **Settings > Environments**, create:

| Environment | Protection |
|-------------|------------|
| development | - |
| qa | - |
| staging | Required reviewers |
| production | Required reviewers |

### 2. Actions Permissions

In **Settings > Actions > General**:

- Workflow permissions: **Read and write permissions**

### 3. Secrets

In **Settings > Secrets and variables > Actions**:

- `FLY_API_TOKEN`: Fly.io deploy token

## Fly.io Setup

```bash
flyctl apps create api-environment-development
flyctl apps create api-environment-qa
flyctl apps create api-environment-staging
flyctl apps create api-environment-production
flyctl tokens create org
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| ENVIRONMENT | Environment name | development |
| VERSION | Application version | 1.0.0 |
| PORT | Server port | 8080 |
