all: participation auth comment common party profile relation story events search

participation:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	participation/*.proto

auth:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	auth/*.proto

comment:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	comment/*.proto

common:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	common/*.proto

party:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	party/*.proto

profile:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	profile/*.proto

relation:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	relation/*.proto
	
story:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	story/*.proto

events:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	events/*.proto

search:
	protoc \
	--go_out . --go_opt paths=source_relative \
	--go-grpc_out . --go-grpc_opt paths=source_relative  \
	search/*.proto
