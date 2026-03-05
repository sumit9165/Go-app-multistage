# Get the base image of golang from dockerhub
FROM golang:1.22

# make directory for workspace
WORKDIR /app
# copy dependency packege installer file
COPY go.mod .
# it install all dependencies.
RUN go mod download
# copy all source codes
COPY . .
# it compiles go app into binary
RUN go build -o server .
# indicate port
EXPOSE 3000
# it run your app
CMD ["./server"]
