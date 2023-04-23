FROM --platform=linux/amd64 golang:1.19 AS build-stage
WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# Deploy the application binary into a lean image
FROM --platform=linux/amd64 gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /main /main

EXPOSE 8081

USER nonroot:nonroot

ENTRYPOINT ["/main"]
