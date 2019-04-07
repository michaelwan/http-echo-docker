FROM amd64/golang:alpine as build
WORKDIR /go/src/app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /echo .

FROM scratch
COPY --from=build /echo .
COPY test.html .
CMD ["/echo"]
