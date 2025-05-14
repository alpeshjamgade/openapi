FROM golang:1.23-alpine AS builder

RUN mkdir /app
COPY . /app
RUN ls
WORKDIR /app
RUN CGO_ENABLED=0 go build -o openapi-client ./cmd/main.go
RUN chmod +x openapi-client

# Install migrate CLI
RUN go install -ldflags="-s -w" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

FROM alpine:latest
RUN mkdir /app

ENV ELASTIC_APM_RECORDING=false
ENV ELASTIC_APM_ACTIVE=false
ENV ELASTIC_APM_ENVIRONMENT=test
ENV ELASTIC_APM_SERVICE_NAME=openapi-client

COPY --from=builder /app/openapi-client /app
COPY --from=builder /go/bin/migrate /usr/local/bin/
COPY private /app/private

CMD [ "./app/openapi-client" ]