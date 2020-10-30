package main

import (
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func sendEmail() {
	var files []string     //O(1)-T O(1)-S
	var emails []string    //O(1)-T O(1)-S
	c := make(chan string) //O(1)-T O(1)-S

	root := "./output"                                                                //O(1)-T O(1)-S
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error { //O(n)-T O(n)-S
		emails = append(emails, strings.Split(path[7:], ".eml")...) //O(1)-T O(1)-S
		files = append(files, path)                                 //O(1)-T O(1)-S
		return nil
	})
	if err != nil { //O(1)-T O(1)-S
		panic(err)
	}
	for index, file := range files {
		fmt.Println(file) //O(1)-T O(1)-S
		if index != 0 {
			data, err := ioutil.ReadFile(file) //O(1)-T O(n)-S
			if err != nil {
				fmt.Println("File reading error", err) //O(1)-T O(1)-S
				return
			}

			go send(emails[index], string(data), c) //O(1)-T O(n)-S

		}
	} //O(n)-T O(n)-S
	for l := range c {
		e := l
		go func(l string, e string) {
			// time.Sleep(5*time.Second)
			send(l, e, c)
		}(l, e)
	} //O(n)-T O(n)-S

	//fmt.Println()
}
func send(email string, body string, c chan string) {
	timer1 := time.Now().UnixNano()
	from := "testuser.sret@gmail.com" //O(1)-T O(1)-S
	password := "testuser"            //O(1)-T O(1)-S

	to := []string{
		email,
	} //O(1)-T O(1)-S
	fmt.Println(to[0]) //O(1)-T O(1)-S

	smtpHost := "smtp.gmail.com" //O(1)-T O(1)-S
	smtpPort := "587"            //O(1)-T O(1)-S

	message := []byte(body) //O(1)-T O(n)-S

	auth := smtp.PlainAuth("", from, password, smtpHost) //O(1)-T O(1)-S

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message) //O(1)-T O(n)-S
	if err != nil {
		fmt.Println(err) //O(1)-T O(1)-S
		c <- "err"
		return
	}

	s := "Email to " + email + " Sent Successfully!" //O(1)-T O(1)-S
	fmt.Println(s)                                   //O(1)-T O(1)-S
	timer2 := time.Now().UnixNano()
	fmt.Println(timer2 - timer1)
	c <- s
	return
}
