FROM golang:1.18.3 as builder
WORKDIR /go/src/stubbygo/cmd
COPY . /go/src/stubbygo
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
WORKDIR /work
COPY --from=builder /go/src/stubbygo/cmd/main /work
ENTRYPOINT ["./main"]
CMD ["-c", "/config/endpoints.yaml"]
