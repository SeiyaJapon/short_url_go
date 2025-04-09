# Configuraciones
IMAGE_NAME = urlshortener
AWS_ACCOUNT_ID = 954976322765
AWS_REGION = us-east-1
REPOSITORY_URI = $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/$(IMAGE_NAME)
PROFILE = personal  # Cambia a tu perfil de AWS si es necesario

# Construye la imagen de Docker
build:
	docker build -t $(IMAGE_NAME) .

# Ejecuta el contenedor en modo interactivo
run:
	docker run -it --entrypoint /bin/sh $(IMAGE_NAME)

# Inicia sesión en ECR y realiza el push de la imagen a AWS
push: login-ecr
	docker tag $(IMAGE_NAME):latest $(REPOSITORY_URI):latest
	docker push $(REPOSITORY_URI):latest

# Inicia sesión en AWS ECR
login-ecr:
	aws ecr get-login-password --region $(AWS_REGION) --profile $(PROFILE) | \
	docker login --username AWS --password-stdin $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com

# Elimina el contenedor y la imagen local
clean:
	-docker rm -f $(shell docker ps -aq --filter ancestor=$(IMAGE_NAME))
	-docker rmi -f $(IMAGE_NAME)

# Elimina la imagen de ECR
clean-ecr:
	aws ecr batch-delete-image --repository-name $(IMAGE_NAME) --image-ids imageTag=latest --region $(AWS_REGION) --profile $(PROFILE)