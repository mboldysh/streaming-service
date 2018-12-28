include .env

NAME := streaming-service
CLOUDFORMATION_BUCKET_STACK_NAME = streaming-service-cloudformation-bucket
DATA_BUCKET_STACK_NAME = streaming-service-data-bucket

all: validate sync deploy
full-deploy: validate create-cloudformation-bucket create-data-bucket sync deploy
update: validate sync update-stack

.PHONY: validate
validate:
	@echo "+ $@"
	validator/validate-templates.sh

.PHONY: sync
sync:
	@echo "+ $@"
	aws s3 sync infrastracture s3://$(CLOUDFORMATIONS_BUCKET) --delete

.PHONY: deploy
deploy:
	@echo "+ $@"
	aws --region $(REGION) cloudformation create-stack --stack-name $(NAME) --template-body file://master.yaml \
	--capabilities CAPABILITY_NAMED_IAM \
	--parameters ParameterKey=DataBucketName,ParameterValue=$(DATA_BUCKET) \
	ParameterKey=Image,ParameterValue=$(IMAGE_ADDRESS)

.PHONY: update-stack
update-stack:
	@echo "+ $@"
	aws --region $(REGION) cloudformation update-stack --stack-name $(NAME) --template-body file://master.yaml \
	--capabilities CAPABILITY_NAMED_IAM \
	--parameters ParameterKey=DataBucketName,ParameterValue=$(DATA_BUCKET) \
	ParameterKey=Image,ParameterValue=$(IMAGE_ADDRESS)

.PHONY: create-cloudformation-bucket
create-cloudformation-bucket:
	@echo "+ $@"
	aws --region $(REGION) cloudformation create-stack --stack-name $(CLOUDFORMATION_BUCKET_STACK_NAME) \
	--template-body file://cloudformation-stack-bucket.yaml \
	--parameters ParameterKey=BucketName,ParameterValue=$(CLOUDFORMATIONS_BUCKET)
	aws --region $(REGION) cloudformation wait stack-create-complete --stack-name $(CLOUDFORMATION_BUCKET_STACK_NAME)

.PHONY: create-data-bucket
create-data-bucket:
	@echo "+ $@"
	aws --region $(REGION) cloudformation create-stack --stack-name $(DATA_BUCKET_STACK_NAME) \
	--template-body file://data-bucket.yaml \
	--parameters ParameterKey=BucketName,ParameterValue=$(DATA_BUCKET)
	aws --region $(REGION) cloudformation wait stack-create-complete --stack-name $(DATA_BUCKET_STACK_NAME)