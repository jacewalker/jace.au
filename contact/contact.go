package contact

import (
	"fmt"

	"github.com/smtp2go-oss/smtp2go-go"
)

type Form struct {
	Name    string
	Email   string
	Phone   string
	Message string
}

func SendEmail(form *Form) error {
	to := fmt.Sprintf("Jace <%s>", "me@jace.au")
	subject := fmt.Sprintf("Contact Form Submission by %s", form.Name)
	textbody := fmt.Sprintf("Message: %s", form.Message)
	htmlbody := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<title>Contact Form Submission on jace.au</title>
		</head>
		<body style="font-family: sans-serif;">
			<h1>Contact Form Submission on jace.au</h1>
			<p>Below are the details of a contact form submission:</p>
			<ul>
				<li><strong>Name:</strong> %s</li>
				<li><strong>Email:</strong> %s</li>
				<li><strong>Phone:</strong> %s</li>
				<li><strong>Message:</strong> %s</li>
			</ul>
			<p>Kind regards,</p>
			<p>Contact Form on jace.au</p>
		</body>
	</html>

`, form.Name, form.Email, form.Phone, form.Message)

	email := smtp2go.Email{
		From: "Website <website@jace.au>",
		To: []string{
			to,
		},
		Subject:  subject,
		TextBody: textbody,
		HtmlBody: htmlbody,
	}
	result, err := smtp2go.Send(&email)
	if err != nil || result.Data.Error != "" {
		fmt.Println("An Error Occurred:", err)
		fmt.Println("SMTP2Go Data Error: " + result.Data.Error)
		fmt.Println("SMTP2Go Data Error Code: " + result.Data.ErrorCode)
		fmt.Println("SMTP2Go Field Validation Error - Field Name: " + result.Data.FieldValidationErrors.FieldName)
		fmt.Println("SMTP2Go Field Validation Error - Message: " + result.Data.FieldValidationErrors.Message)
		return err
	}
	fmt.Println("Email sent successfully.")
	return nil
}
