package helper

import (
	"bufio" //mail
	"fmt"
	"os"                 //mail
	"strconv"            //mail
	"strings"            //mail

	"gopkg.in/gomail.v2" //mail
)

//for mail
var (
	wd = func() string {
		d, _ := os.Getwd()
		return d + "/"
	}()
)

func ReadEmailConfig() map[string]string {
	ret := make(map[string]string)
	file, err := os.Open(wd + "conf/app.conf")
	if err == nil {
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			line, _, e := reader.ReadLine()
			if e != nil {
				break
			}

			sval := strings.Split(string(line), "=")
			ret[sval[0]] = sval[1]
		}
	} else {
		fmt.Println(err.Error())
	}

	return ret
}

func SendingEmailNotif(emailTo string, emailSubject string, emailMessage string) {
	//send email
	mailconfig := ReadEmailConfig()

	m := gomail.NewMessage()
	m.SetHeader("From", mailconfig["MailAddress"])
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", emailSubject)
	m.SetBody("text/html", emailMessage)

	port, _ := strconv.Atoi(mailconfig["Port"])

	gm := gomail.NewPlainDialer(mailconfig["Host"], port, mailconfig["MailAddress"], mailconfig["MailAddressPassword"])
	err := gm.DialAndSend(m)
	if err != nil {
		fmt.Println(err)
	}

	return
}

func SendEmail(to map[string]string, subject string, message string, isHtml bool) {
	SendEmailWithAttachment(to, subject, message, isHtml, "")
}

func SendEmailWithAttachment(to map[string]string, subject string, message string, isHtml bool, attachmentLoc string) {
	mailconfig := ReadEmailConfig()

	m := gomail.NewMessage()

	var senders []string
	senders = append(senders, mailconfig["MailAddress"])
	if mailconfig["MailAddressName"] != "" {
		senders = append(senders, mailconfig["MailAddressName"])
	}

	var recipients []string
	recipients = append(recipients, to["ToMail"])
	if to["ToName"] != "" {
		recipients = append(recipients, to["ToName"])
	}

	m.SetHeader("From", senders...)
	m.SetHeader("To", recipients...)
	m.SetHeader("Subject", subject)
	if isHtml {
		m.SetBody("text/html", message)
	} else {
		m.SetBody("text", message)
	}

	if attachmentLoc != "" {
		m.Attach(attachmentLoc)
	}

	port, _ := strconv.Atoi(mailconfig["Port"])

	gm := gomail.NewPlainDialer(mailconfig["Host"], port, mailconfig["MailAddress"], mailconfig["MailAddressPassword"])
	err := gm.DialAndSend(m)
	if err != nil {
		fmt.Println(err)
	}

	return
}
