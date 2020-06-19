# Accept the Go version for the image to be set as a build argument.
ARG GO_VERSION=1.13.3

# Execute the build using an alpine linux environment so that it will execute properly in the final environment
FROM golang:${GO_VERSION}-alpine AS builder

ARG PACKAGE_MAIN_PATH
RUN test -n "$PACKAGE_MAIN_PATH" || (echo "PACKAGE_MAIN_PATH build argument not set, required for building binary" && false)
ENV CGO_ENABLED=0

# Create the user and group files that will be used in the running container to
# run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
# Git is required for fetching the dependencies.
RUN apk add ca-certificates

# Set a directory to contain the go app to be compiled (this directory will work for go apps making use of go modules as long as go 1.13+ is used)
WORKDIR /go/src/app

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./go.sum ./
RUN go mod download

# Import the code from the context.
COPY . .

WORKDIR /go/src/app/${PACKAGE_MAIN_PATH}

# Build the executable to `/app`. Mark the build as statically linked.
RUN go build \
    -installsuffix 'static' \
    -o /app .

# Final stage: the running container.
FROM alpine AS final

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the compiled executable from the first stage.
COPY --from=builder /app /app

# Perform any further action as an unprivileged user.
USER nobody:nobody

# Run the compiled binary.
ENTRYPOINT ["/app"]