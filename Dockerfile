# Start from golang v1.11 base image
FROM golang:1.11 as builder

# Add Maintainer Info
LABEL maintainer="BULGARE Gregory <gregory@beamery.com>"

# install xz
RUN apt update \
    && apt install -y xz-utils ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# install UPX
ADD https://github.com/upx/upx/releases/download/v3.94/upx-3.94-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.94-amd64_linux.tar.xz | \
    tar -xOf - upx-3.94-amd64_linux/upx > /bin/upx \
    && chmod a+x /bin/upx

# Download dependencies
RUN go get -u github.com/golang/dep/cmd/dep

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/leangeder/chatops2

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download dependencies
RUN dep ensure --vendor-only
# RUN go get -d -v ./...

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/chatops2 cmd/beambot/beambot.go

# use scratch (base for a docker image)
FROM scratch

# set working directory
WORKDIR /root

# copy the binary from builder
COPY --from=builder /go/bin/chatops2 .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

CMD ["./chatops2"] 
