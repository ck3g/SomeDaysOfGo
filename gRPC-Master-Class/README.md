# gRPS Master Class

Hands-on exercises from [gRPC [Golang] Master Class: Build Modern API & Microservices](https://www.udemy.com/share/101Zo0A0QaeF1aTHg=/) course.

## Install dependencies

1. `brew install clang-format`
2. `brew install protobuf`
3. [gRPC-Go](https://github.com/grpc/grpc-go)

    ```
    go get -u google.golang.org/grpc
    ```

4. [protobuf/protoc-gen-go](https://github.com/golang/protobuf)

    ```
    go get -u github.com/golang/protobuf/protoc-gen-go
    ```


## Resources

* https://grpc.io
* https://developers.google.com/protocol-buffers
* https://imagekit.io/demo/http2-vs-http1

### Error handling

* https://grpc.io/docs/guides/error/
* http://avi.im/grpc-errors/

### Deadlines

* https://grpc.io/blog/deadlines/


### SSL

* https://grpc.io/docs/guides/auth/
* https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-auth-support.md


### Reflection

* https://github.com/grpc/grpc-go/tree/master/reflection
* https://github.com/ktr0731/evans
