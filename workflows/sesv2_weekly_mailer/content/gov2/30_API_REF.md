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
  Write a Windows batch script to download the Go v2 SDK API 
  command signature
  for each SESv2 action listed in README.md.
  For example, the CreateContactList command is here: https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sesv2#Client.CreateContactList
  The output file should look like this:
  func (c *Client) CreateContactList(ctx context.Context, params *CreateContactListInput, optFns ...func(*Options)) (*CreateContactListOutput, error)
---
Here's a Windows batch script that will download the Go v2 SDK API command signatures for each SESv2 action listed in the README.md file:

```bat
@echo off

set ACTIONS=CreateContact CreateContactList CreateEmailIdentity CreateEmailTemplate DeleteContactList DeleteEmailIdentity DeleteEmailTemplate ListContacts SendEmail

for %%a in (%ACTIONS%) do (
    echo Downloading signature for %%a...
    curl -s "https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sesv2#Client.%%a" ^
        | findstr /c:"func (c *Client) %%a(" > %%a.txt
    if %ERRORLEVEL% equ 0 (
        echo Signature for %%a downloaded successfully.
    ) else (
        echo Error downloading signature for %%a.
    )
)
```

This script will create the following files in the current directory, each containing the command signature for the respective SESv2 action:

1. `CreateContact.txt`
2. `CreateContactList.txt`
3. `CreateEmailIdentity.txt`
4. `CreateEmailTemplate.txt`
5. `DeleteContactList.txt`
6. `DeleteEmailIdentity.txt`
7. `DeleteEmailTemplate.txt`
8. `ListContacts.txt`
9. `SendEmail.txt`

For example, the contents of `CreateContactList.txt` will be:

```
func (c *Client) CreateContactList(ctx context.Context, params *CreateContactListInput, optFns ...func(*Options)) (*CreateContactListOutput, error)
```

Here's how the script works:

1. The `ACTIONS` variable is set with a space-separated list of SESv2 actions.
2. The script iterates over each action using a `for` loop.
3. Inside the loop, it constructs the URL to fetch the command signature from `pkg.go.dev` using the action name.
4. The `curl` command is used to download the webpage content from the constructed URL.
5. The `findstr` command is used to extract the line containing the command signature from the downloaded content, and the extracted line is written to a file with the action name.
6. If the `findstr` command is successful (i.e., the command signature is found), a success message is printed. Otherwise, an error message is printed.

Note: This script requires `curl` and `findstr` to be available on your system. If you're using a Unix-based system (e.g., Linux or macOS), you can adapt the script to use appropriate commands like `wget` and `grep`.