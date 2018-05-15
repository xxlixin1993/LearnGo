# 如何使用golang的protocolbuf

```
go get -u github.com/golang/protobuf/protoc-gen-go
protoc -I=./ --go_out=./ ./Person.proto //protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto

```