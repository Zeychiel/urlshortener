# Stage 1: Build the Go app
FROM golang:latest AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Create appuser.
ENV USER=appuser
ENV UID=10001 
# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

# Copy the source from the current directory to the Working Directory inside the container
COPY . .
RUN go mod download
# Build the Go app
# RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app cmd/api/main.go
RUN --mount=type=cache,target="/root/.cache/go-build" --mount=type=cache,target="/go/pkg/mod" go build -a -o shortener cmd/api/main.go

# Command to run the executable
ENTRYPOINT ["/app/shortener"]


# Stage 2: Create a minimal image with the executable
# FROM scratch
# 
# # Import the user and group files from the builder.
# COPY --from=builder /etc/passwd /etc/passwd
# COPY --from=builder /etc/group /etc/group
# 
# # Copy the executable from the builder stage
# COPY --from=builder /app /app
# 
# # Use an unprivileged user.
# USER appuser:appuser
# 
# # Expose port 8000 to the outside world
# EXPOSE 8000
# 
# # Command to run the executable
# ENTRYPOINT ["/app/shortener"]

