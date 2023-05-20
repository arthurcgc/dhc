package handlers

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/mail"

	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"net/url"

	echo "github.com/labstack/echo/v4"
)

// Index handles the index page.
func Index(c echo.Context) error {
	data := struct {
		Title string
	}{
		Title: "MyApp",
	}
	return c.Render(http.StatusOK, "index.html", data)
}

// Email handles an email request from the client.
func Email(c echo.Context) error {
	type Contact struct {
		Name string 	`json:"name"`
		Email string 	`json:"email"`
		Phone string 	`json:"phone"`
	}

	contact := Contact{}
	if err := c.Bind(&contact); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, contact)
}

const (
	apiKey       = "YOUR_API_KEY"
	recipient    = "RECIPIENT_EMAIL"
	senderEmail  = "YOUR_SENDER_EMAIL"
	subject      = "Test Subject"
	messageBody  = "Hello, this is a test email!"
	gmailAPIURL  = "https://www.googleapis.com/gmail/v1/users/me/messages/send"
)

func send() {
	// Create the email message
	msg := createMessage(senderEmail, recipient, subject, messageBody)

	// Encode the message as base64 URL-safe
	rawMsg := base64.URLEncoding.EncodeToString([]byte(msg))

	// Create the API request URL
	apiURL := fmt.Sprintf("%s?key=%s", gmailAPIURL, apiKey)

	// Send the email
	err := sendEmail(apiURL, rawMsg)
	if err != nil {
		log.Fatalf("Unable to send email: %v", err)
	}

	fmt.Println("Email sent successfully!")
}

// Helper function to create a formatted email message
func createMessage(sender, recipient, subject, body string) string {
	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		mail.Address{Address: sender}.String(),
		mail.Address{Address: recipient}.String(),
		subject,
		body)
	return msg
}

// Helper function to send the email using HTTP POST
func sendEmail(apiURL, rawMsg string) error {
	data := url.Values{}
	data.Set("raw", rawMsg)

	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send email: %s", resp.Status)
	}

	return nil
}