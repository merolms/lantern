FROM golang:1.13-alpine AS lantern-build

RUN apk add --no-cache make

LABEL com.meroedu.name="lantern" \
  com.meroedu.maintainer="Trang Mai <trangmaiq@gmail.com>"

COPY . /src/lantern
WORKDIR /src/lantern

RUN go mod download

RUN make all

EXPOSE 9090

ENTRYPOINT ["./lantern"]
CMD ["serve"]
