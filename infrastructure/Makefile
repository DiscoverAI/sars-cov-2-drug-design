project_name = "sars-cov-2"
admin_role_name = "administrator"
developer_role_name = "developer"
group_name = "base-user"
org_name = "tortugas"
region = "eu-central-1"

admin_profile = "tortugas-administrator"
developer_profile = "tortugas-developer"
wait_for_status_change=5

now := $(shell date +%s)

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
		--change-set-name "$(project_name)-administrator-access-${now}" \
		--change-set-type UPDATE \
		--template-body file://./administrator-access.yaml \
		--stack-name "$(project_name)-administrator-access" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile tortugas-base \
		--parameters ParameterKey=ProjectName,ParameterValue=$(project_name) ParameterKey=RoleName,ParameterValue=$(admin_role_name) ParameterKey=OrgName,ParameterValue=$(org_name)

	sleep $(wait_for_status_change)

	aws cloudformation execute-change-set \
		--change-set-name "$(project_name)-administrator-access-${now}" \
		--stack-name "$(project_name)-administrator-access" \
		--region $(region) \
		--profile $(admin_profile)

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
		--change-set-name "$(project_name)-developer-access-${now}" \
		--change-set-type UPDATE \
		--template-body file://./developer-access.yaml \
		--stack-name "$(project_name)-developer-access" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile $(admin_profile) \
		--parameters ParameterKey=ProjectName,ParameterValue=$(project_name) ParameterKey=RoleName,ParameterValue=$(developer_role_name) ParameterKey=OrgName,ParameterValue=$(org_name)

	sleep $(wait_for_status_change)

	aws cloudformation execute-change-set \
		--change-set-name "$(project_name)-developer-access-${now}" \
		--stack-name "$(project_name)-developer-access" \
		--region $(region) \
		--profile $(admin_profile)

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
		--change-set-name "$(project_name)-users-${now}" \
		--change-set-type UPDATE \
		--template-body file://./users.yaml \
		--stack-name "$(project_name)-users" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile $(admin_profile) \
		--parameters ParameterKey=ProjectName,ParameterValue=$(project_name) ParameterKey=GroupName,ParameterValue=$(group_name) ParameterKey=OrgName,ParameterValue=$(org_name)

	sleep $(wait_for_status_change)

	aws cloudformation execute-change-set \
		--change-set-name "$(project_name)-users-${now}" \
		--stack-name "$(project_name)-users" \
		--region $(region) \
		--profile $(admin_profile)

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
		--change-set-name "$(project_name)-base-user-group-${now}" \
		--change-set-type UPDATE \
		--template-body file://./base-user-group.yaml \
		--stack-name "$(project_name)-base-user-group" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile $(admin_profile) \
		--parameters ParameterKey=ProjectName,ParameterValue=$(project_name) ParameterKey=GroupName,ParameterValue=$(group_name) ParameterKey=OrgName,ParameterValue=$(org_name)

	sleep $(wait_for_status_change)

	aws cloudformation execute-change-set \
		--change-set-name "$(project_name)-base-user-group-${now}" \
		--stack-name "$(project_name)-base-user-group" \
		--region $(region) \
		--profile $(admin_profile)

.PHONY: create-billing-alarm-template-test
create-billing-alarm-template-test:
	aws cloudformation validate-template --profile $(developer_profile) --region "us-east-1" --template-body file://./base-user-group.yaml >> /dev/null

.PHONY: update-billing-alarm
update-billing-alarm:
	aws cloudformation create-change-set \
		--change-set-name "$(project_name)-billing-alarm-${now}" \
		--change-set-type UPDATE \
		--template-body file://./billing-alarm.yaml \
		--stack-name "$(project_name)-billing-alarm" \
		--region "us-east-1" \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile $(developer_profile) \
		--parameters ParameterKey=AlarmThresholds,ParameterValue=\"10,25,50,100,200\"

	sleep $(wait_for_status_change)

	aws cloudformation execute-change-set \
		--change-set-name "$(project_name)-billing-alarm-${now}" \
		--stack-name "$(project_name)-billing-alarm" \
		--region "us-east-1" \
		--profile $(developer_profile)

.PHONY: create-billing-alarm
create-billing-alarm:
	aws cloudformation deploy \
		--template-file billing-alarm.yaml \
		--stack-name "$(project_name)-billing-alarm" \
		--region "us-east-1" \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile $(developer_profile) \
		--parameter-overrides ProjectName=$(project_name) GroupName=$(group_name) AlarmThresholds="10,25,50,100,200"

.PHONY: create-vpc
create-vpc-template-test:
	aws cloudformation validate-template --profile $(developer_profile) --region $(region) --template-body file://./vpc.yaml >> /dev/null

.PHONY: create-vpc
create-vpc:
	aws cloudformation deploy \
		--template-file vpc.yaml \
		--stack-name "$(project_name)-vpc" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile $(developer_profile) \
		--parameter-overrides ProjectName=$(project_name) EnvironmentName=$(project_name) VpcCIDR=10.1.0.0/16 PublicSubnet1CIDR=10.1.1.0/24 PrivateSubnet1CIDR=10.1.9.0/24 Region=$(region)

.PHONY: update-vpc
update-vpc:
	aws cloudformation create-change-set \
		--change-set-name "$(project_name)-vpc-${now}" \
		--change-set-type UPDATE \
		--template-body file://./vpc.yaml \
		--stack-name "$(project_name)-vpc" \
		--region $(region) \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile $(developer_profile) \
		--parameters ParameterKey=ProjectName,ParameterValue=$(project_name) ParameterKey=EnvironmentName,ParameterValue=$(project_name) ParameterKey=VpcCIDR,ParameterValue=10.1.0.0/16 ParameterKey=PublicSubnet1CIDR,ParameterValue=10.1.1.0/24 ParameterKey=PublicSubnet2CIDR,ParameterValue=10.1.2.0/24 ParameterKey=PrivateSubnet1CIDR,ParameterValue=10.1.9.0/24 ParameterKey=PrivateSubnet2CIDR,ParameterValue=10.1.8.0/24 ParameterKey=Region,ParameterValue=$(region)

	sleep $(wait_for_status_change)

	aws cloudformation execute-change-set \
		--change-set-name "$(project_name)-vpc-${now}" \
		--stack-name "$(project_name)-vpc" \
		--region $(region) \
		--profile $(developer_profile)



.PHONY: templates-test
templates-test: create-base-user-group-template-test create-users-template-test create-developer-access-template-test create-administrator-access-template-test create-billing-alarm-template-test

.PHONY: create-infrastructure
create-infrastructure: create-base-user-group create-users create-developer-access create-administrator-access create-billing-alarm

.PHONY: update-infrastructure
update-infrastructure: update-base-user-group update-users update-developer-access update-administrator-access
