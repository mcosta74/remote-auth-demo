#########################################
## BUILDER
FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o dist/authsvc main.go

#####################################
## RUNTIME
FROM scratch AS prod

WORKDIR /app

COPY --from=builder /app/dist/authsvc ./

EXPOSE 8080

ENTRYPOINT [ "/app/authsvc" ]