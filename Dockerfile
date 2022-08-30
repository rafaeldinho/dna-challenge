FROM golang:alpine as builder

ENV GOBIN=$GOPATH/bin
ENV GO111MODULE="on"

RUN mkdir -p $GOPATH/src/github.com/rafaeldinho131613/envs
WORKDIR $GOPATH/src/github.com/rafaeldinho131613/envs

COPY go.mod .
COPY src src/
COPY app.env .

RUN cd src
RUN ls -l


RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -ldflags '-extldflags "-static"' -o $GOBIN/main ./src/main.go

# Runtime image with scratch container
FROM scratch
ARG VERSION
ENV VERSION_APP=$VERSION

COPY --from=builder /go/bin/ /src/
COPY app.env .

EXPOSE 8080
ENTRYPOINT ["/src/main"]