NAME := streaming-service
REGION := us-east-2

all: validate sync deploy
update: validate sync update-stack

.PHONY: validate
validate:
	@echo "+ $@"
	validator/validate-templates.sh

.PHONY: sync
sync:
	@echo "+ $@"
	aws s3 sync infrastracture s3://streaming-service-cloudformation --delete

.PHONY: deploy
deploy:
	@echo "+ $@"
	aws --region $(REGION) cloudformation create-stack --stack-name $(NAME) --template-body file://master.yaml --capabilities CAPABILITY_NAMED_IAM

.PHONY: update-stack
update-stack:
	@echo "+ $@"
	aws --region $(REGION) cloudformation update-stack --stack-name $(NAME) --template-body file://master.yaml --capabilities CAPABILITY_NAMED_IAM