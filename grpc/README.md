一个简单的gRPC样例，来自Go高级编程

分client和server两部分，两个文件夹下的hello.proto需要保持一致。

protoc --go_out=plugins=grpc:. hello.proto