project_name = "sars-cov-2"
admin_role_name = "administrator"
developer_role_name = "developer"
group_name = "base-user"
org_name = "tortugas"
region = "eu-central-1"

admin_profile = "tortugas-administrator"
developer_profile = "tortugas-developer"

.PHONY: create-administrator-access-template-test
create-administrator-access-template-test:
	aws cloudformation validate-template --profile tortugas-base --template-body file://./administrator-access.yaml >> /dev/null

.PHONY: create-administor-access
create-administrator-access:
	aws cloudformation deploy \
		--template-file administrator-access.yaml \
		--stack-name "$(project_name)-administrator-access" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile tortugas-base \
		--parameter-overrides ProjectName=$(project_name) RoleName=$(admin_role_name) OrgName=$(org_name)

.PHONY: update-administrator-access
update-administrator-access:
	aws cloudformation create-change-set \
		--change-set-name "$(project_name)-administrator-access-$(shell date +%s)" \
		--change-set-type UPDATE \
		--template-body file://./administrator-access.yaml \
		--stack-name "$(project_name)-administrator-access" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile tortugas-base \
		--parameters ParameterKey=ProjectName,ParameterValue=$(project_name) ParameterKey=RoleName,ParameterValue=$(admin_role_name) ParameterKey=OrgName,ParameterValue=$(org_name)

.PHONY: create-developer-access-template-test
create-developer-access-template-test:
	aws cloudformation validate-template --profile $(admin_profile) --region $(region) --template-body file://./developer-access.yaml >> /dev/null

.PHONY: create-developer-access
create-developer-access:
	aws cloudformation deploy \
		--template-file developer-access.yaml \
		--stack-name "$(project_name)-developer-access" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile $(admin_profile) \
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
		--profile $(admin_profile) \
		--parameters ParameterKey=ProjectName,ParameterValue=$(project_name) ParameterKey=RoleName,ParameterValue=$(developer_role_name) ParameterKey=OrgName,ParameterValue=$(org_name)

.PHONY: create-users-template-test
create-users-template-test:
	aws cloudformation validate-template --profile $(admin_profile) --region $(region) --template-body file://./users.yaml >> /dev/null

.PHONY: create-users
create-users:
	aws cloudformation deploy \
		--template-file users.yaml \
		--stack-name "$(project_name)-users" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile $(admin_profile) \
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
		--profile $(admin_profile) \
		--parameters ParameterKey=ProjectName,ParameterValue=$(project_name) ParameterKey=GroupName,ParameterValue=$(group_name) ParameterKey=OrgName,ParameterValue=$(org_name)


.PHONY: create-base-user-group-template-test
create-base-user-group-template-test:
	aws cloudformation validate-template --profile $(admin_profile) --region $(region) --template-body file://./base-user-group.yaml >> /dev/null

.PHONY: create-base-user-group
create-base-user-group:
	aws cloudformation deploy \
		--template-file base-user-group.yaml \
		--stack-name "$(project_name)-base-user-group" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile $(admin_profile) \
		--parameter-overrides ProjectName=$(project_name) GroupName=$(group_name) OrgName=$(org_name)

.PHONY: update-base-user-group
update-base-user-group:
	aws cloudformation create-change-set \
		--change-set-name "$(project_name)-base-user-group-$(shell date +%s)" \
		--change-set-type UPDATE \
		--template-body file://./base-user-group.yaml \
		--stack-name "$(project_name)-base-user-group" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile $(admin_profile) \
		--parameters ParameterKey=ProjectName,ParameterValue=$(project_name) ParameterKey=GroupName,ParameterValue=$(group_name) ParameterKey=OrgName,ParameterValue=$(org_name)

.PHONY: create-billing-alarm-template-test
create-billing-alarm-template-test:
	aws cloudformation validate-template --profile $(developer_profile) --region "us-east-1" --template-body file://./base-user-group.yaml >> /dev/null

.PHONY: update-billing-alarm
update-billing-alarm:
	aws cloudformation create-change-set \
		--change-set-name "$(project_name)-billing-alarm-$(shell date +%s)" \
		--change-set-type UPDATE \
		--template-body file://./billing-alarm.yaml \
		--stack-name "$(project_name)-billing-alarm" \
		--region "us-east-1" \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile $(developer_profile) \
		--parameters ParameterKey=ProjectName,ParameterValue=$(project_name) ParameterKey=GroupName,ParameterValue=$(group_name) ParameterKey=AlarmThresholds,ParameterValue="10,25,50,100,200"

.PHONY: create-billing-alarm
create-billing-alarm:
	aws cloudformation deploy \
		--template-file billing-alarm.yaml \
		--stack-name "$(project_name)-billing-alarm" \
		--region "us-east-1" \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile $(developer_profile) \
		--parameter-overrides ProjectName=$(project_name) GroupName=$(group_name) AlarmThresholds="10,25,50,100,200"

.PHONY: templates-test
templates-test: create-base-user-group-template-test create-users-template-test create-developer-access-template-test create-administrator-access-template-test create-billing-alarm-template-test

.PHONY: create-infrastructure
create-infrastructure: create-base-user-group create-users create-developer-access create-administrator-access create-billing-alarm

.PHONY: update-infrastructure
update-infrastructure: update-base-user-group update-users update-developer-access update-administrator-access
