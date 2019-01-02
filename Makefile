build:
	protoc 	--micro_out=. --go_out=. \
		proto/auth/auth.proto
	protoc-go-inject-tag -input=./proto/auth/auth.pb.go
	protoc-go-inject-field -input=./proto/auth/auth.pb.go
	docker build -t shippy-user-service .
buildproto:
	protoc 	--micro_out=. --go_out=. \
		proto/auth/auth.proto
	protoc-go-inject-tag -input=./proto/auth/auth.pb.go
	protoc-go-inject-field -input=./proto/auth/auth.pb.go
	
run:
	docker run --net="host" \
		-p 50051 \
		-e DB_HOST=localhost:54321 \
		-e DB_PASSWORD=postgres \
		-e DB_USER=postgres \
		-e MICRO_SERVER_ADDRESS=:50055 \
		-e MICRO_REGISTRY=mdns \
		shippy-user-service

deploy:
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yml
	kubectl replace -f ./deployments/deployment.yml
