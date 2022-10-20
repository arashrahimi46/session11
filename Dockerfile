FROM harbor.tasn.ir/library/docker-base-images/golang-alpine:latest  AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/merger

ADD . .

RUN go env -w GO111MODULE=on
RUN go mod tidy

RUN go build -o ./out/merger main.go

#start runner
FROM harbor.tasn.ir/library/alpine:latest

WORKDIR /app/merger
COPY --from=build_base /tmp/merger/out/merger .

CMD ["./merger"]