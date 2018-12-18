from quay.io/jonnrb/go as build
add . /src
run cd /src && CGO_ENABLED=0 go get ./cmd/three_oh_two

from gcr.io/distroless/base
expose 8080
copy --from=build /go/bin/three_oh_two /three_oh_two
entrypoint ["/three_oh_two"]
