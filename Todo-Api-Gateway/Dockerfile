FROM golang:1.21.2-alpine3.18 AS build-stage
WORKDIR /todo
COPY ./ /todo
RUN mkdir -p /todo/build
RUN go mod download
RUN go build -v -o /todo/build/api ./cmd


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /todo/build/api /
COPY --from=build-stage /todo/.env /
EXPOSE 5000
CMD ["/api"]

