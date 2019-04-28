package main

import gomail "gopkg.in/gomail.v2"

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "###USERNAME@comcast.net")
	m.SetHeader("To", "###USERNAMEOTHER@EMAIL.com")
	m.SetHeader("Subject", "Config File")
	m.SetBody("text/html", "Config Files by me")
	m.Attach("###FILE")

	d := gomail.NewDialer("smtp.comcast.net", 587, "###USERNAME", "###USERNAMECREDIT")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
