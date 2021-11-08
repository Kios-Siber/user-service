init:
	go mod init skeleton
	
gen:
	protoc --proto_path=proto --go_out=paths=source_relative,plugins=grpc:./pb proto/*/*.proto
	
.PHONY: init gen