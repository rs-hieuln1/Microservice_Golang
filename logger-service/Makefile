# Makefile for logger-service

PROTO_SRC_DIR = logs/proto
PROTO_OUT_DIR = logs

.PHONY: generate clean

# Generate Go code from proto files
generate:
	@echo "🧩 Generating Go code from proto files..."
	protoc \
		--go_out=$(PROTO_OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_OUT_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_SRC_DIR)/*.proto
	@echo "✅ Done generating protobuf files."

# Clean generated files
clean:
	@echo "🧹 Cleaning generated protobuf files..."
	rm -f $(PROTO_OUT_DIR)/*.pb.go
	@echo "✅ Cleaned."
