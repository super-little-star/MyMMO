
protoc --proto_path=./protobuf --go_out=../Server/Src message.proto

"Protogen/net462/protogen" --proto_path="protobuf" message.proto --csharp_out="../Client\ClientProject\Assets\Scripts\Protobuf" 
@pause