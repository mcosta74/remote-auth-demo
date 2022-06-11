#########################################
## BUILDER
FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o dist/coolsvc main.go

#####################################
## RUNTIME
FROM scratch AS prod

# Secret key for token encryption 
ENV SECRET_KEY ""

WORKDIR /app

COPY --from=builder /app/dist/coolsvc ./

EXPOSE 8080

ENTRYPOINT [ "/app/coolsvc" ]