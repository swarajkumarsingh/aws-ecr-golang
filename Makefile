deploy:
	aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin <aws_ecr_image_url>
	docker build -t test-name .
	docker tag test-name:latest <aws_ecr_image_url>/test-name:latest
	docker push <aws_ecr_image_url>/test-name:latest