FROM golang:1.26-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/student-api ./cmd/api

FROM alpine:3.22

RUN apk add --no-cache ca-certificates \
	&& adduser -D -H -s /sbin/nologin studentapi

COPY --from=build --chown=studentapi:studentapi /out/student-api /student-api

USER studentapi
EXPOSE 8080
ENTRYPOINT ["/student-api"]