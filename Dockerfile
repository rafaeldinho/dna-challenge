FROM golang:alpine as builder

RUN apk add -U --no-cache ca-certificates && update-ca-certificates

ENV GOBIN=$GOPATH/bin
ENV GO111MODULE="on"

RUN mkdir -p $GOPATH/src/github/meli
WORKDIR $GOPATH/src/github/meli

COPY go.mod .
RUN  go mod tidy
COPY . .
COPY app.env .

RUN cd src && pwd && ls -la

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o $GOBIN/server ./src/main.go

# Runtime image with scratch container
FROM scratch
ARG VERSION
ENV VERSION_APP=$VERSION

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/ /src/
COPY app.env .
COPY . .
EXPOSE 8080
ENTRYPOINT ["/src/main"]