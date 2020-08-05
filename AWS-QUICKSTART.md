# AWS Quickstart

## Access Architecture
When receiving AWS credentials (for an AWS user) from an Administrator, this user is simply a limited user that cannot do anything or access anything outside of managing your own security credentials (e.g. MFA, access tokens). From this user, you will be able to [assume roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html) which in turn have a set of permissions allowing you to do specific lines of work.

The current roles that are available are:
* Administrator (manage IAM Users)
* Developer (create compute machines via PowerUserAccess, create machine users)

Most likely, you will be given the permissions to assume the Developer role (and not the Administrator role). For more details, please see the [permissions delegated to the Developer role](infrastructure/developer-access.yaml).

## Console Access
In order to assume the Developer (or any other role) in the AWS Account, you will need to first [set up an MFA device](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_mfa_enable_virtual.html). *Then log out and log back in again (weird AWS quirk)*.

You can access the Developer role through this URL (replace `<SOME-AWS-ACCOUNT-ID>` with the 10-digit AWS account number provided to you with your credentials):
```bash
https://signin.aws.amazon.com/switchrole?account=<SOME-AWS-ACCOUNT-ID>&roleName=tortugas/developer
```

## CLI Access
We are using a development environment provided by [kelseyomk/assume-role](https://github.com/kelseymok/assume-role). Please follow the instructions there to set up your CLI.

Most importantly, the `base_profile` name we are using is `tortugas` (this will make sense after your read the docs).