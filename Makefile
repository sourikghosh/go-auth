up:
	docker-compose -f ${GOPATH}/src/personalProjects/go-auth/deployments/docker-compose.yml --env-file .env up

down:
	docker-compose -f ${GOPATH}/src/personalProjects/go-auth/deployments/docker-compose.yml --env-file .env down