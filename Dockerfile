FROM golang:latest as builder

#Set Environment Variables
ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /app 
COPY go.mod go.sum ./

RUN go mod download

#Adds source code 
ADD cmd cmd
ADD config config
ADD database database
ADD internals internals
ADD router router

#Build the application 
RUN go build -a -installsuffix cgo -o main cmd/server/main.go

#Start a new stage from scratch 
FROM alpine:latest 

WORKDIR /root/

#Copy the prebuilt binary from previous stage.
COPY --from=builder /app/main .

EXPOSE 8080

#Command to run executable
CMD ["./main"]