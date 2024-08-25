FROM golang:alpine as builder

WORKDIR /app
RUN apk add --no-cache make nodejs npm

COPY . ./
RUN make build
RUN make install
RUN > /app/.env

FROM scratch
COPY --from=builder /ftgodev-tut /ftgodev-tut
COPY --from=builder /app/.env .env

EXPOSE 3000
ENTRYPOINT [ "./ftgodev-tut" ]
