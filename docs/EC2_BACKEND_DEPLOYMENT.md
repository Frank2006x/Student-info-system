# Deploying the Backend to EC2

This guide deploys the Go API in `backend/` to an EC2 instance. The recommended production setup is:

```text
Frontend (Vercel) -> HTTPS/reverse proxy -> EC2 Go API -> PostgreSQL (RDS)
```

Do not expose PostgreSQL directly to the internet. Allow RDS access only from the EC2 security group.

## Required backend changes

### 1. Make the HTTP address configurable

`backend/cmd/api/main.go` currently listens on a hard-coded `:8080`. Change it to read an environment value:

```go
httpAddr := os.Getenv("HTTP_ADDR")
if httpAddr == "" {
	httpAddr = ":8080"
}

log.Printf("student information API listening on %s", httpAddr)
if err := app.Listen(httpAddr); err != nil {
	log.Fatal(err)
}
```

Add `"os"` to the imports. `HTTP_ADDR=:8080` is suitable when Nginx or an AWS load balancer is in front of the API.

### 2. Keep production secrets out of Git

The existing configuration already reads `DATABASE_URL` from `.env` or environment variables. Keep `.env` local only. On EC2, set the value through Docker/systemd or AWS Secrets Manager.

Example production values:

```text
APP_ENV=production
HTTP_ADDR=:8080
DATABASE_URL=postgres://USER:PASSWORD@RDS_ENDPOINT:5432/students_db?sslmode=require
```

Use a URL-encoded password if it contains characters such as `@`, `:`, `/`, or `#`.

### 3. Configure CORS when the frontend is deployed separately

The current frontend does not call the API yet. When Vercel calls the EC2 API from the browser, add Fiber CORS middleware and allow only the real frontend origin, for example:

```go
app.Use(cors.New(cors.Config{
	AllowOrigins: []string{"https://app.example.com"},
}))
```

Do not use `*` in production if the API will use cookies or authenticated browser requests.

## Should this be Dockerized?

Docker is optional, but recommended for this project. It gives EC2 the same Go runtime every time, avoids installing Go and project tools on the server, and makes restarts and upgrades predictable.

Docker does not replace RDS. Run the API container on EC2 and use RDS for production PostgreSQL. The existing `docker-compose.yml` is for local PostgreSQL development and should not be used as the production database setup.

Add `backend/Dockerfile`:

```dockerfile
FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/student-api ./cmd/api

FROM gcr.io/distroless/static-debian13:nonroot
COPY --from=build /out/student-api /student-api
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/student-api"]
```

Also add `backend/.dockerignore`:

```text
.env
.git
postman
**/*_test.go
```

## EC2 setup

1. Create an EC2 instance, preferably Ubuntu LTS, in the same VPC as RDS.
2. Configure the EC2 security group to allow SSH only from your IP and HTTP/HTTPS as required.
3. Configure the RDS security group to allow TCP `5432` only from the EC2 security group.
4. Install Docker on EC2 and log in to the image registry you use.
5. Build the image locally or in CI and push it to Docker Hub, Amazon ECR, or another private registry.
6. Pull the image on EC2 and run it with the production environment values.

Example Docker commands on EC2:

```bash
docker pull YOUR_REGISTRY/student-api:latest

docker run -d \
  --name student-api \
  --restart unless-stopped \
  -p 8080:8080 \
  -e APP_ENV=production \
  -e HTTP_ADDR=:8080 \
  -e DATABASE_URL='postgres://USER:PASSWORD@RDS_ENDPOINT:5432/students_db?sslmode=require' \
  YOUR_REGISTRY/student-api:latest
```

Do not put the real command with the database password in shell history. Prefer an environment file with permissions restricted to root, or AWS Secrets Manager/SSM Parameter Store.

## Database migration

Run the migration against the production database before starting a release:

```bash
migrate -path backend/db/migration \
  -database "$DATABASE_URL" \
  -verbose up
```

Do not run `make migrateup` in production because the current Makefile contains local development credentials and `localhost`.

## Reverse proxy and HTTPS

Do not make the Go process directly responsible for public HTTPS. Put Nginx, an Application Load Balancer, or another managed reverse proxy in front of it.

The public API should be available at a stable hostname such as:

```text
https://api.example.com
```

The proxy forwards to `http://127.0.0.1:8080` or the EC2 private address. Configure the frontend with:

```text
NEXT_PUBLIC_API_BASE_URL=https://api.example.com
```

## Verification

After deployment, check the API locally on EC2 and through the public hostname:

```bash
curl http://127.0.0.1:8080/health
curl https://api.example.com/health
docker logs --tail 100 student-api
```

Expected health response:

```json
{"status":"ok"}
```

## Native alternative: systemd

Docker is not mandatory. A native deployment can build the binary with `go build`, copy it to EC2, and run it under systemd with `DATABASE_URL` and `HTTP_ADDR` in a protected environment file. This has fewer moving parts, but the server must then manage the Go version, binary replacement, rollback, and service dependencies itself.

For this repository, Docker plus RDS is the simpler repeatable production path.