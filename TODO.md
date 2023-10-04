# IDOMPET
This is what I learned during the development of the idompet project

## Working with database
- [x] Design schema database
- [x] Install docker and image Postgresql
- [x] Write Database migration
- [x] Generate CRUD code from sqlc
- [x] Write golang unit test with random data
- [x] Database transaction implementation
- [x] Handle deadlock
- [x] Avoid database deadlock
- [x] Set up Github action for automated test

## Building Restful API
- [x] Implement Restful using Gin
- [x] Make and load config using Viper
- [ ] Mock database for test API
- [ ] Implement transfer money with custom parameter validation
- [ ] Add user table 
- [ ] Handle database error in golang
- [ ] Store and decrypt password using Bcrypt
- [ ] Write stronger unit test using gomock
- [ ] Create and verify JWT and PASETO
- [ ] Implement user login that return JWT or PASETO
- [ ] Implement authentication middleware and authorization using Gin

## Deploying app to prod
- [ ] Build minimal Golang docker image
- [ ] Use docker network to connect 2 stand-alone container
- [ ] Write docker compose file and control service start-up orders
- [ ] Create free tier AWS Acoount
- [ ] Auto build and push docker image to AWS ECR with Github action
- [ ] Create a production databse on AWS RDS
- [ ] Store & retrieve production secrets with AWS secrets manager
- [ ] Create an EKS cluster on AWS
- [ ] Use kubectl and k9s to connetc to kubernetes cluster on AWS EKS
- [ ] Deploy a web app to Kubernetes cluster on AWS EKS
- [ ] Register domain and set up A-record using Route53
- [ ] Use Ingress to route traffics to different service in Kubernetes
- [ ] Automatic issue TLS certificates in Kubernetes with Lest's Encrypt
- [ ] Automatic deploy to kubernetes with github actions

## Advance Backend Topic
- [ ] Manage user session with refresh token
- [ ] Generate documentation page and SQL schema dump from DBML
- [ ] gRPC
- [ ] Define gRPC API and generate Go code with protobuf
- [ ] Run gRPC Server and call it's API
- [ ] Implement gRPC API to create and user login
- [ ] Write code once, serve gRPC API and HTTP request
- [ ] Extract info from gRPC metadata
- [ ] Automatic generate and serve Swagger docs from Go server
- [ ] Embed static frontend files inside golang backend server's binary
- [ ] validate gRPC parameters and send humanmachine friendly response
- [ ] Run database migration directly inside Go code
- [ ] Partial update database record with sqlc nullable parameters
- [ ] build gRPC update API with optional parameter
- [ ] Add authorization to protect gRPC API
- [ ] Write structure logs for gRPC API
- [ ] Write HTTP logger middleware in Go

## Asyncronous
- [ ] Implement background worker in Go with Redis and async
- [ ] Integrate async workers to Go web server
- [ ] Send async task to Redis within in a database transaction
- [ ] How to handle error and print logs for Go async workers
- [ ] A bit of delay might be good for your async tasks
- [ ] How to send emails in Go via Gmail
- [ ] How to skip test in Go and config test flag in vscode
- [ ] Email verification in Go design DB and send email
- [ ] Implement email verification API in Go
- [ ] Unit test gRPC API with mock DB & Redis
- [ ] How to test a gRPC API that requires authentication

## Improve stability security server
- [ ] Config sqlc version 2 for Go and Postgres
- [ ] Switch DB driver from libpq to pgx
- [ ] How to handle DB errors with PGX driver