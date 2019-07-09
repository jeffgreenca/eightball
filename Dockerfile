FROM golang:1.12.6-alpine3.9 as builder
WORKDIR ~/project
COPY . .
RUN go mod vendor && go build -o /eightball

FROM alpine:3.9
COPY --from=builder /eightball .
CMD ["/eightball"]
