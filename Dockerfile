FROM golang:alpine AS build-env

ENV MONGODB_URL localhost:27017

RUN mkdir /hello
WORKDIR /hello
COPY go.mod .
COPY go.sum .

RUN apk add --update --no-cache ca-certificates git

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/hello
FROM scratch 
COPY --from=build-env /go/bin/hello /go/bin/hello
EXPOSE 8000
ENTRYPOINT ["/go/bin/hello"]