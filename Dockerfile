FROM golang:1.19.4-alpine3.17 AS builder

RUN apk add --update make build-base

WORKDIR /app

# Install dependencies
COPY go.mod ./
COPY go.sum ./
COPY Makefile ./
RUN make install

# Copy source code into image
COPY ./ ./

RUN make build

FROM alpine:3.17.0 AS runner

COPY --from=builder /app/build/golox /bin/golox
