project_name = "sars-cov-2"
admin_role_name = "administrator"
developer_role_name = "developer"
group_name = "base-user"
org_name = "tortugas"
region = "eu-central-1"

.PHONY: create-administrator-access-template-test
create-administrator-access-template-test:
	aws cloudformation validate-template --profile tortugas-base-administrator --template-body file://./administrator-access.yaml >> /dev/null

.PHONY: create-administor-access
create-administrator-access:
	aws cloudformation deploy \
		--template-file administrator-access.yaml \
		--stack-name "$(project_name)-administrator-access" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile tortugas-base-administrator \
		--parameter-overrides ProjectName=$(project_name) RoleName=$(admin_role_name)

.PHONY: update-administrator-access
update-administrator-access:
	aws cloudformation create-change-set \
		--change-set-name "$(project_name)-administrator-access-$(shell date +%s)" \
		--change-set-type UPDATE \
		--template-body file://./administrator-access.yaml \
		--stack-name "$(project_name)-administrator-access" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile tortugas-base-administrator \
		--parameters ParameterKey=ProjectName,ParameterValue=$(project_name) ParameterKey=RoleName,ParameterValue=$(admin_role_name)

.PHONY: create-developer-access-template-test
create-developer-access-template-test:
	aws cloudformation validate-template --profile tortugas-base-administrator --region $(region) --template-body file://./developer-access.yaml >> /dev/null

.PHONY: create-developer-access
create-developer-access:
	aws cloudformation deploy \
		--template-file developer-access.yaml \
		--stack-name "$(project_name)-developer-access" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile tortugas-base-administrator \
		--parameter-overrides ProjectName=$(project_name) RoleName=$(developer_role_name) OrgName=$(org_name)

.PHONY: update-developer-access
update-developer-access:
	aws cloudformation create-change-set \
		--change-set-name "$(project_name)-developer-access-$(shell date +%s)" \
		--change-set-type UPDATE \
		--template-body file://./developer-access.yaml \
		--stack-name "$(project_name)-developer-access" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile tortugas-base-administrator \
		--parameters ParameterKey=ProjectName,ParameterValue=$(project_name) ParameterKey=RoleName,ParameterValue=$(developer_role_name) ParameterKey=OrgName,ParameterValue=$(org_name)

.PHONY: create-users-template-test
create-users-template-test:
	aws cloudformation validate-template --profile tortugas-base-administrator --region $(region) --template-body file://./users.yaml >> /dev/null

.PHONY: create-users
create-users:
	aws cloudformation deploy \
		--template-file users.yaml \
		--stack-name "$(project_name)-users" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile tortugas-base-administrator \
		--parameter-overrides ProjectName=$(project_name) GroupName=$(group_name) OrgName=$(org_name)

.PHONY: update-users
update-users:
	aws cloudformation create-change-set \
		--change-set-name "$(project_name)-users-$(shell date +%s)" \
		--change-set-type UPDATE \
		--template-body file://./users.yaml \
		--stack-name "$(project_name)-users" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile tortugas-base-administrator \
		--parameters ParameterKey=ProjectName,ParameterValue=$(project_name) ParameterKey=GroupName,ParameterValue=$(group_name) ParameterKey=OrgName,ParameterValue=$(orgName)


.PHONY: create-base-user-group-template-test
create-base-user-group-template-test:
	aws cloudformation validate-template --profile tortugas-base-administrator --region $(region) --template-body file://./base-user-group.yaml >> /dev/null

.PHONY: create-base-user-group
create-base-user-group:
	aws cloudformation deploy \
		--template-file base-user-group.yaml \
		--stack-name "$(project_name)-base-user-group" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile tortugas-base-administrator \
		--parameter-overrides ProjectName=$(project_name) GroupName=$(group_name)

.PHONY: update-base-user-group
update-base-user-group:
	aws cloudformation create-change-set \
		--change-set-name "$(project_name)-base-user-group-$(shell date +%s)" \
		--change-set-type UPDATE \
		--template-body file://./base-user-group.yaml \
		--stack-name "$(project_name)-base-user-group" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile tortugas-base-administrator \
		--parameters ParameterKey=ProjectName,ParameterValue=$(project_name) ParameterKey=GroupName,ParameterValue=$(group_name)


.PHONY: templates-test
templates-test: create-base-user-group-template-test create-users-template-test create-developer-access-template-test create-administrator-access-template-test

.PHONY: create-infrastructure
create-infrastructure: create-base-user-group create-users create-developer-access create-administrator-access create-billing-alarm

.PHONY: update-infrastructure
update-infrastructure: update-base-user-group update-users update-developer-access update-administrator-access

.PHONY: create-billing-alarm
create-billing-alarm:
	aws cloudformation deploy \
		--template-file billing-alarm-template.yaml \
		--stack-name "$(project_name)-billing-alarm" \
		--region "us-east-1" \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile tortugas-base-developer \
		--parameter-overrides ProjectName=$(project_name) GroupName=$(group_name)
