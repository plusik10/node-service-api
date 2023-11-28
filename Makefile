.PHONY: generate

generate:
	mkdir -p pkg/note_v1
	protoc --proto_path vendor.protogen --proto_path api/note_v1 \
				--go_out=pkg/note_v1 --go_opt=paths=source_relative \
				--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative \
				--grpc-gateway_out=pkg/note_v1 \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=source_relative \
				-I . --openapiv2_out ./api/note_v1 \
				--validate_out lang=go:pkg/note_v1 \
				--validate_opt=paths=source_relative \
				api/note_v1/note.proto


LOCAL_MIGRATION_DIR = ./migrations
LOCAL_MIGRATION_DSN = "host=localhost port=54321 dbname=note-service user=postgres password=qwerty sslmode=disable"

.PHONY: install-goose
.install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: local-migration-status
local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v


.PHONY: local-migration-up
local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v


.PHONY: local-migration-down
local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v


PHONY: vendor-proto
vendor-proto: .vendor-proto

PHONY: .vendor-proto
.vendor-proto:
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google/protobuf ]; then \
			git clone https://github.com/protocolbuffers/protobuf vendor.protogen/protobuf &&\
			mkdir -p  vendor.protogen/google/protobuf &&\
			mv vendor.protogen/protobuf/src/google/protobuf/*.proto vendor.protogen/google/protobuf &&\
			rm -rf vendor.protogen/protobuf ;\
		fi


.PHONY: server-run
server-run:
	PG_URL='postgresql://postgres:qwerty@localhost:54321/note-service' go run cmd/server/main.go

.PHONY: client-run
client-run:
	go run cmd/client/main.go


