FROM golang:1.19.4-alpine AS builder

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

FROM golang:1.19.4-alpine AS runner

COPY --from=builder /app/build/golox /bin/golox
