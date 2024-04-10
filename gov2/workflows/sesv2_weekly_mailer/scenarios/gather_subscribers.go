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