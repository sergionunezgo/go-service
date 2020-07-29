FROM golang:1.14 as build

WORKDIR /go/src/app

RUN groupadd -g 999 appuser && useradd -r -u 999 -g appuser appuser

COPY . .

ENV CGO_ENABLED 0

RUN make build

RUN chown -R appuser:appuser /go/src/app/build/goservice

FROM scratch as release

COPY --from=build /etc/passwd /etc/passwd

COPY --from=build /go/src/app/build/goservice /goservice

USER appuser

ENTRYPOINT [ "/goservice" ]