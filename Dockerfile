FROM golang:1.20.3-alpine AS build-env
WORKDIR /app
COPY . .
RUN go build -ldflags "-s -w" -o main .

FROM scratch
COPY --from=build-env /app /app
WORKDIR /app
ENTRYPOINT [ "/app/main" ]