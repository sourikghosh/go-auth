FROM golang:1.17-alpine3.15 AS build_base

# Updating and Installing git make c/c++ compilers for CGO 
RUN apk add git make gcc g++

# GO ENVs
ENV GO111MODULE on
ENV GOBIN $GOPATH/bin

WORKDIR $GOPATH/src/tmp

# Extracting mod and sum files for installing package level dependencies
COPY go.mod ./
COPY go.sum ./

# Downloading dependencies for the package
RUN go mod download

# Coping the complete source code
COPY . .

# Build the package
RUN go build -o auth cmd/main.go

# Starting fresh with only binary
FROM alpine:3.15 

WORKDIR /app

# Copy the binary from `build_base` step
COPY --from=build_base /go/src/tmp/auth ./

EXPOSE 6969

CMD ["./auth"]