# Deploying a Go Server to AWS Lambda

This repository provides a simple guide and example code for deploying a Dockerized Go server to AWS ECR & ECS. By leveraging AWS ECR & ECS, you can host your Go application with minimal infrastructure management and cost.

![a](https://github.com/swarajkumarsingh/aws-ecr-golang/assets/89764448/ef94f67b-3fa7-4c87-a097-1a43120b7897)

## Table of Contents

- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [Workflow](#workflow)
- [Run Go Server](#building-the-go-server)
- [Deploying to AWS ECR](#deploying-to-aws-lambda)
- [Testing](#testing)
- [Clean-Up](#clean-up)
- [Contributing](#contributing)
- [License](#license)

## Prerequisites

Before you begin, make sure you have the following installed and set up:

- [Go](https://golang.org/doc/install)
- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)

## Getting Started

Clone this repository to your local machine:

```bash
git clone https://github.com/swarajkumarsingh/aws-ecr-golang.git
cd aws-ecr-golang
```

## Run Go Server
```
go run main.go
```

## Configuration
Makefile - file contains all configs related to image deployment
```
deploy:
	aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin <aws_ecr_image_url>
	docker build -t test-name .
	docker tag test-name:latest <aws_ecr_image_url>/test-name:latest
	docker push <aws_ecr_image_url>/test-name:latest
```


## Workflow
- First we build our docker image
- Then we push it to AWS ECR
- Then setup AWS ECS cluster
- Create a task defining your OS, expose your port carefully
- Inside AWS ECS cluster create a service, and add your task with other configurations
- Create a Load-Balancer and set the values accordingly
- After everything succeeds, your Load-Balancer ARM url, will be url backend server url  

## Deploying to AWS ECR
Use serverless cli for deploying
```
make deploy
``` 

## Testing
After you deploy, you will be given a url
```
curl <given-url>
``` 

## Clean-Up
```
- Head to your AWS console
- Search 'ECR'
- Delete the repository
- Search 'ECS'
- Delete the cluster & tasks
- Head to Load-Balancer
- Delete the load-balancer
- And all set :)
```

## Contributing
If you find any issues or have improvements, feel free to open an issue or submit a pull request. Contributions are welcome!

## License
This project is licensed under the MIT License - see the LICENSE file for details.
