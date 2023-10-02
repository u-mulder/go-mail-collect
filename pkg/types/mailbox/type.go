package mailbox

import (
	"mailbox_protocol"
)

type Mailbox struct {
	host     string
	port     string
	protocol MailboxProtocol
	useSSL   bool
}

func New(host, port string, useSSL bool, protocol mailbox_protocol.MailboxProtocol) Mailbox {
	return Mailbox{
		host: host,
		port: port,
		protocol: protocol, 
		useCase: useSSL,
	}
}

func NewIMAP(host, port string, useSSL bool) Mailbox {
	return Mailbox{
		host: host,
		port: port,
		protocol: mailbox_protocol.IMAP, 
		useCase: useSSL,
	}
}

func NewPOP3(host, port string, useSSL bool) Mailbox {
	return Mailbox{
		host: host,
		port: port,
		protocol: mailbox_protocol.POP3, 
		useCase: useSSL,
	}
}
