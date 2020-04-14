#
# Builder
#
FROM golang:1.14-alpine3.11 AS builder

# Define app directory
WORKDIR /build

# Copy app files
COPY go.mod go.sum cmd/exotic/exotic.go ./
COPY internal internal
COPY cmd cmd
COPY Taskfile.yml Taskfile.yml

# Install app deps
RUN apk --no-cache add git gcc && \
    # Get go-task
    TASK_VERSION=2.8.0 && \
    wget -q https://github.com/go-task/task/releases/download/v${TASK_VERSION}/task_linux_amd64.tar.gz && \
    tar -xzf task_linux_amd64.tar.gz task && \
    mv task /usr/local/bin/ && \
    chmod ugo+x /usr/local/bin/task && \
    rm task_linux_amd64.tar.gz && \
    # Build app
    task app:build

#
# App
#
FROM alpine:3.11

# Labels
LABEL Maintainer="Julien BREUX <julien.breux@gmail.com>"

# Copy app binary
COPY --from=builder /buld/dist/exotic /bin/exotic

# Install required deps
RUN apk add --update ca-certificates && \
    # Install temp deps
    # apk add --update -t deps ... && \
    # Purge deps
    # apk del --purge deps && \
    rm /var/cache/apk/*

# Define entry point
ENTRYPOINT [ "/bin/exotic" ]

# Define default command
CMD [ "serve" ]
