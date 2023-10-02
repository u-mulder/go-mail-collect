package mailbox_protocol

type MailboxProtocol string

const (
	IMAP MailboxProtocol = "IMAP"
	POP3 MailboxProtocol = "POP3"
)