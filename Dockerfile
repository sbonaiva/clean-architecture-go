FROM golang:1.18 AS build

WORKDIR /app

COPY ./src ./

RUN go mod download

RUN GOOS=linux go build -o /clean-architecture-go

FROM gcr.io/distroless/base-debian11 AS deploy

WORKDIR /

COPY --from=build /clean-architecture-go /clean-architecture-go
COPY --from=build /app/local.env /local.env

EXPOSE 8080

USER nonroot:nonroot

CMD [ "/clean-architecture-go" ]