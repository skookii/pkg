package mail

import (
	"crypto/tls"
	"io"
	"log"
	"net"
	"net/smtp"
	"os"

	"github.com/domodwyer/mailyak/v3"
)

var (
	from string
	addr string
	pass string
)

func SetConfig() {
	from = os.Getenv("fromEmail")
	addr = os.Getenv("smtpAddr")
	pass = os.Getenv("smtpPass")
}

func SendAttach(fromName, to, subject, txt, fileName string, rd io.Reader) {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		log.Println(err)
		return
	}
	m, err := mailyak.NewWithTLS(addr,
		smtp.PlainAuth("", from, pass, host),
		&tls.Config{
			ServerName: host,
		})
	if err != nil {
		log.Println(err)
		return
	}

	m.To(to)
	m.From(from)
	m.FromName(fromName)

	m.Subject(subject)
	m.HTML().Set(txt)
	m.Attach(fileName, rd)
	if err := m.Send(); err != nil {
		log.Println(err)
	}
}

func Send(fromName, to, subject, txt string) {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		log.Println(err)
		return
	}
	m, err := mailyak.NewWithTLS(addr,
		smtp.PlainAuth("", from, pass, host),
		&tls.Config{
			ServerName: host,
		})
	if err != nil {
		log.Println(err)
		return
	}

	m.To(to)
	m.From(from)
	m.FromName(fromName)

	m.Subject(subject)
	m.HTML().Set(txt)
	if err := m.Send(); err != nil {
		log.Println(err)
	}
}

func SendInline(fromName, to, subject, txt, imageName string, rd io.Reader) {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		log.Println(err)
		return
	}
	m, err := mailyak.NewWithTLS(addr,
		smtp.PlainAuth("", from, pass, host),
		&tls.Config{
			ServerName: host,
		})
	if err != nil {
		log.Println(err)
		return
	}

	m.To(to)
	m.From(from)
	m.FromName(fromName)

	m.Subject(subject)
	
	//<img src="cid:myimage"/>
	m.HTML().Set(txt)
	m.AttachInline(imageName, rd)
	if err := m.Send(); err != nil {
		log.Println(err)
	}
}
