FROM golang:1.16
WORKDIR /src
ENV CGO_ENABLED=0
ENV GO111MODULE=on
COPY . .
ARG SERVICE_NAME
RUN go build -o bin/main-exe cmd/${SERVICE_NAME}/main.go
CMD ["bin/main-exe"]