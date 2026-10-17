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