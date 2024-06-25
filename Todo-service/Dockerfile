FROM golang:1.21.2-alpine3.18 AS build-stage
WORKDIR /todo_svc
COPY ./ /todo_svc
RUN mkdir -p /todo_svc/build
RUN go mod download
RUN go build -v -o /todo_svc/build/api ./cmd


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /todo_svc/build/api /
COPY --from=build-stage /todo_svc/.env /
EXPOSE 5050
CMD ["/api"]