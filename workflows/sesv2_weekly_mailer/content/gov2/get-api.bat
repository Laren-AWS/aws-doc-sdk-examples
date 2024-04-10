@echo off

set ACTIONS=CreateContact CreateContactList CreateEmailIdentity CreateEmailTemplate DeleteContactList DeleteEmailIdentity DeleteEmailTemplate ListContacts SendEmail

for %%a in (%ACTIONS%) do (
    echo Downloading signature for %%a...
    curl -s "https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sesv2#Client.%%a" ^
        | findstr /c:"func (c *Client) %%a(" > 30_%%a.md
    if %ERRORLEVEL% equ 0 (
        echo Signature for %%a downloaded successfully.
    ) else (
        echo Error downloading signature for %%a.
    )
)