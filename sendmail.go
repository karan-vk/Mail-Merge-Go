package main

import (
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
)

func sendEmail() {
	var files []string
	var emails []string
	c := make(chan string)

	root := "./output"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		emails = append(emails, strings.Split(path[7:], ".eml")...)
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for index, file := range files {
		fmt.Println(file)
		if index != 0 {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Println("File reading error", err)
				return
			}
			send(emails[index], string(data),c)
			//fmt.Println(string(data))
		}
	}
	for l := range c {
		e := l
		go func(l string,e string) {
			// time.Sleep(5*time.Second)
			send(l,e, c)
		}(l,e)
	}

	//fmt.Println()
}
func send(email string, body string, c chan string ) {
	from := "testuser.sret@gmail.com"
	password := "testuser"

	// Receiver email address.
	to := []string{
		email,
	}
	fmt.Println(to[0])

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(body)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		c <- "err"
		return
	}

	s := "Email to " + email + " Sent Successfully!"
	fmt.Println(s)
	c <- s
	return
}
