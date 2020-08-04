project_name = "sars-cov-2"
admin_role_name = "administrator"
developer_role_name = "developer"
group_name = "base-user"
org_name = "tortugas"
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

.PHONY: developer-access-template-test
developer-access-template-test:
	aws cloudformation validate-template --profile base --region $(region) --template-body file://./developer-access.yaml >> /dev/null

.PHONY: developer-access
developer-access:
	aws cloudformation deploy \
		--template-file developer-access.yaml \
		--stack-name "$(project_name)-developer-access" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile base \
		--parameter-overrides ProjectName=$(project_name) RoleName=$(developer_role_name)

.PHONY: users-template-test
users-template-test:
	aws cloudformation validate-template --profile base --region $(region) --template-body file://./users.yaml >> /dev/null

.PHONY: users
users:
	aws cloudformation deploy \
		--template-file users.yaml \
		--stack-name "$(project_name)-users" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile base \
		--parameter-overrides ProjectName=$(project_name) GroupName=$(group_name) OrgName=$(org_name)


.PHONY: base-user-group-template-test
base-user-group-template-test:
	aws cloudformation validate-template --profile base --region $(region) --template-body file://./base-user-group.yaml >> /dev/null

.PHONY: base-user-group
base-user-group:
	aws cloudformation deploy \
		--template-file base-user-group.yaml \
		--stack-name "$(project_name)-base-user-group" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile base \
		--parameter-overrides ProjectName=$(project_name) GroupName=$(group_name)


.PHONY: templates-test
templates-test: base-user-group-template-test users-template-test developer-access-template-test administrator-access-template-test

.PHONY: infrastructure
infrastructure: base-user-group users developer-access administrator-access