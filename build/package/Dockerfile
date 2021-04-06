####
# Build the image with:
#
# docker build -f build/package/Dockerfile -t gocrudperson .
#
# Then run the container using:
#
# docker run -i --rm -p 8080:8080 gocrudperson
#
###

# imagem base para compilar a aplicacao
FROM golang:alpine3.13 as builder

WORKDIR /build

COPY go.mod go.sum ./
COPY cmd cmd
COPY internal internal

RUN echo "Building. GOOS: $(go env GOOS), GOARCH: $(go env GOARCH)" && \
    go build -o /build/bin/main /build/cmd/main.go

####################################################################

# imagem base final que executara a aplicacao compilada
FROM alpine:3.13

WORKDIR /app

COPY --from=builder /build/bin/main ./bin/main

COPY api api
COPY web web

EXPOSE 8080

CMD [ "/app/bin/main" ]
