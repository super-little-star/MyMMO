
protoc --proto_path=./protobuf --go_out=../Server/Src message.proto
protoc --proto_path=./protobuf --csharp_out=../Client\ClientProject\Assets\Scripts\Protobuf message.proto
@pause