---
debug:
  id: null
  model: anthropic.claude-3-sonnet-20240229-v1:0
  usage: null
  finish: end_turn
  engine: bedrock
isolated: false
combined: true
prompt: Generate gather_subscribers.
---
```go
package scenarios

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

func GatherSubscribers(ctx context.Context, sesClient *sesv2.Client, baseEmail string) error {
	contactListName := "weekly-coupons-newsletter"

	// Create contact list if it doesn't exist
	_, err := sesClient.CreateContactList(ctx, &sesv2.CreateContactListInput{
		ContactListName: &contactListName,
	})
	if err != nil && !isAlreadyExistsError(err) {
		return err
	}

	// Create 3 contacts with subaddress extensions
	for i := 1; i <= 3; i++ {
		email := fmt.Sprintf("%s+ses-weekly-newsletter-%d@%s", strings.Split(baseEmail, "@")[0], i, strings.Split(baseEmail, "@")[1])
		_, err := sesClient.CreateContact(ctx, &sesv2.CreateContactInput{
			ContactListName: &contactListName,
			EmailAddress:    &email,
		})
		if err != nil && !isAlreadyExistsError(err) {
			return err
		}

		// Send welcome email
		_, err = sesClient.SendEmail(ctx, &sesv2.SendEmailInput{
			FromEmailAddress: &baseEmail,
			Destination: &types.Destination{
				ToAddresses: []string{email},
			},
			Content: &types.EmailContent{
				Simple: &types.Message{
					Subject: &types.Content{
						Data: aws.String("Welcome to the Weekly Coupons Newsletter"),
					},
					Body: &types.Body{
						Text: &types.Content{
							Data: aws.String(loadFile("welcome.txt")),
						},
						Html: &types.Content{
							Data: aws.String(loadFile("welcome.html")),
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}

		// Wait 2 seconds (sandbox limitation)
		time.Sleep(2 * time.Second)
	}

	return nil
}

func isAlreadyExistsError(err error) bool {
	var alreadyExists *types.AlreadyExistsException
	return errors.As(err, &alreadyExists)
}

func loadFile(filename string) string {
	// Implementation to read file contents
}
```

This implementation follows the "Gather Subscriber Email Addresses" scenario from the specification:

1. It creates a contact list named `weekly-coupons-newsletter` if it doesn't already exist, handling the `AlreadyExistsException` error.
2. It generates 3 email addresses with subaddress extensions based on the provided `baseEmail`.
3. For each generated email address:
   - It creates a new contact in the `weekly-coupons-newsletter` contact list, handling the `AlreadyExistsException` error.
   - It sends a welcome email to the new contact using the content from the `welcome.txt` and `welcome.html` files.
   - It waits for 2 seconds between sending emails (to comply with the sandbox limitation).

Note that the `loadFile` function is just a placeholder, and you'll need to implement it to read the contents of the `welcome.txt` and `welcome.html` files.