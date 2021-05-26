FROM golang:1.16 AS build-env

RUN mkdir /app
WORKDIR /app

COPY ./go.mod /app
COPY ./go.sum /app

RUN go mod download

COPY ./ /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /app/pig

EXPOSE 8000

CMD /app/pig

FROM alpine

RUN mkdir /app
WORKDIR /app
COPY --from=build-env /app/pig/ /app/pig
RUN mkdir /app/resources
COPY --from=build-env /app/resources/index.html /app/resources/index.html

EXPOSE 8000
CMD /app/pig

