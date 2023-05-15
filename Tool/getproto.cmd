
protoc --proto_path=./protobuf --go_out=../Server/Src message.proto request.proto response.proto error.proto enum.proto model.proto core.proto

"Protogen/net462/protogen" --proto_path="protobuf"  --csharp_out="../Client\ClientProject\Assets\Scripts\Protobuf"  message.proto request.proto response.proto error.proto enum.proto model.proto core.proto
@pause