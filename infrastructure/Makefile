project_name = "sars-cov-2"
admin_role_name = "administrator"
develeper_role_name = "developer"
region = "eu-central-1"

.PHONY: administrator-access-template-test
administrator-access-template-test:
	aws cloudformation validate-template --profile base --template-body file://./administrator-access.yaml >> /dev/null

.PHONY: administor-access
administrator-access:
	aws cloudformation deploy \
		--template-file administrator-access.yaml \
		--stack-name "$(project_name)-administrator-access" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile base \
		--parameter-overrides ProjectName=$(project_name) RoleName=$(admin_role_name)

#.PHONY: datalake-template-test
#datalake-template-test:
#	aws cloudformation validate-template --template-body file://./datalake.yaml >> /dev/null
#
#.PHONY: datalake
#datalake:
#	aws cloudformation deploy \
#		--template-file datalake.yaml \
#		--region $(region) \
#		--stack-name "$(project_name)-datalake" \
#		--parameter-overrides BucketName=$(bucket_name) ProjectName=$(project_name)
#
#.PHONY: templates-test
#templates-test: datalake-template-test deploy-user-template-test deploy-user-key-template-test
#
#.PHONY: infrastructure
#infrastructure: datalake deploy-user deploy-user-key
