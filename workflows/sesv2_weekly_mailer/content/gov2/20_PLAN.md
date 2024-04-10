---
debug:
  id: null
  model: anthropic.claude-3-sonnet-20240229-v1:0
  usage: null
  finish: end_turn
  engine: bedrock
isolated: false
combined: true
prompt: |
  Create a plan to implement this workflow in the Go programming language.

  Prefer putting SES API actions into an actions folder, main in a cmd folder,
  and interactive parts in a scenarios folders. Put tests in the same folder
  as code files, named with an _test suffix. Write test stubs for each API
  action in a stubs folder.

  Describe the files that we will need as a Windows batch script to create their outlines.

  Describe the files that we will need.
---
Here's a plan to implement the SES v2 Coupon Newsletter Workflow in Go, following the suggested folder structure and naming conventions:

File Creation Plan (Windows Batch Script):

```bat
@echo off

rem Create folders
mkdir actions
mkdir cmd
mkdir scenarios
mkdir stubs

rem Create main package
cd cmd
type nul > main.go
cd ..

rem Create actions package
cd actions
type nul > create_contact.go
type nul > create_contact_list.go
type nul > create_email_identity.go
type nul > create_email_template.go
type nul > delete_contact_list.go
type nul > delete_email_identity.go
type nul > delete_email_template.go
type nul > list_contacts.go
type nul > send_email.go
cd ..

rem Create scenarios package
cd scenarios
type nul > gather_subscribers.go
type nul > prepare_application.go
type nul > send_newsletter.go
cd ..

rem Create stubs package
cd stubs
type nul > create_contact_stub.go
type nul > create_contact_list_stub.go
type nul > create_email_identity_stub.go
type nul > create_email_template_stub.go
type nul > delete_contact_list_stub.go
type nul > delete_email_identity_stub.go
type nul > delete_email_template_stub.go
type nul > list_contacts_stub.go
type nul > send_email_stub.go
cd ..
```

Files Needed:

1. **cmd/main.go**: The entry point of the application, where the workflow is executed.

2. **actions/create_contact.go**: Implementation of the `CreateContact` API action.
   - **actions/create_contact_test.go**: Test cases for the `CreateContact` action.

3. **actions/create_contact_list.go**: Implementation of the `CreateContactList` API action.
   - **actions/create_contact_list_test.go**: Test cases for the `CreateContactList` action.

4. **actions/create_email_identity.go**: Implementation of the `CreateEmailIdentity` API action.
   - **actions/create_email_identity_test.go**: Test cases for the `CreateEmailIdentity` action.

5. **actions/create_email_template.go**: Implementation of the `CreateEmailTemplate` API action.
   - **actions/create_email_template_test.go**: Test cases for the `CreateEmailTemplate` action.

6. **actions/delete_contact_list.go**: Implementation of the `DeleteContactList` API action.
   - **actions/delete_contact_list_test.go**: Test cases for the `DeleteContactList` action.

7. **actions/delete_email_identity.go**: Implementation of the `DeleteEmailIdentity` API action.
   - **actions/delete_email_identity_test.go**: Test cases for the `DeleteEmailIdentity` action.

8. **actions/delete_email_template.go**: Implementation of the `DeleteEmailTemplate` API action.
   - **actions/delete_email_template_test.go**: Test cases for the `DeleteEmailTemplate` action.

9. **actions/list_contacts.go**: Implementation of the `ListContacts` API action.
   - **actions/list_contacts_test.go**: Test cases for the `ListContacts` action.

10. **actions/send_email.go**: Implementation of the `SendEmail` API action.
    - **actions/send_email_test.go**: Test cases for the `SendEmail` action.

11. **scenarios/gather_subscribers.go**: Implementation of the "Gather Subscriber Email Addresses" scenario.
    - **scenarios/gather_subscribers_test.go**: Test cases for the "Gather Subscriber Email Addresses" scenario.

12. **scenarios/prepare_application.go**: Implementation of the "Prepare the Application" scenario.
    - **scenarios/prepare_application_test.go**: Test cases for the "Prepare the Application" scenario.

13. **scenarios/send_newsletter.go**: Implementation of the "Send the Coupon Newsletter" scenario.
    - **scenarios/send_newsletter_test.go**: Test cases for the "Send the Coupon Newsletter" scenario.

14. **stubs/create_contact_stub.go**: Stub implementation for the `CreateContact` API action.

15. **stubs/create_contact_list_stub.go**: Stub implementation for the `CreateContactList` API action.

16. **stubs/create_email_identity_stub.go**: Stub implementation for the `CreateEmailIdentity` API action.

17. **stubs/create_email_template_stub.go**: Stub implementation for the `CreateEmailTemplate` API action.

18. **stubs/delete_contact_list_stub.go**: Stub implementation for the `DeleteContactList` API action.

19. **stubs/delete_email_identity_stub.go**: Stub implementation for the `DeleteEmailIdentity` API action.

20. **stubs/delete_email_template_stub.go**: Stub implementation for the `DeleteEmailTemplate` API action.

21. **stubs/list_contacts_stub.go**: Stub implementation for the `ListContacts` API action.

22. **stubs/send_email_stub.go**: Stub implementation for the `SendEmail` API action.

This structure follows the recommended organization, separating API actions, scenarios, and stubs into their respective packages. The `cmd` package contains the main entry point, and the `scenarios` package handles the high-level workflow scenarios by utilizing the API actions from the `actions` package. The `stubs` package provides stub implementations for testing purposes.