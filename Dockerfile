FROM golang:1.23-alpine AS builder

RUN mkdir /app
COPY . /app
RUN ls
WORKDIR /app
RUN CGO_ENABLED=0 go build -o gocms ./cmd/main.go
RUN chmod +x gocms

FROM alpine:latest
RUN mkdir /app

ENV ELASTIC_APM_RECORDING=false
ENV ELASTIC_APM_ACTIVE=false
ENV ELASTIC_APM_ENVIRONMENT=test
ENV ELASTIC_APM_SERVICE_NAME=gocms

COPY --from=builder /app/gocms /app
CMD [ "./app/gocms" ]