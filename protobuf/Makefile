all: gen-ticket gen-participation gen-comment gen-common gen-profile gen-relation gen-story gen-events gen-search

gen-ticket:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	ticket/*.proto

gen-participation:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	participation/*.proto

gen-comment:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	comment/*.proto

gen-common:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	common/*.proto

gen-profile:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	profile/*.proto

gen-relation:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	relation/*.proto

gen-story:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	story/*.proto

gen-events:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	events/*.proto

gen-search:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	search/*.proto
