package domain

import (
	_ "pkg/types" // TODO correct path
)

type MailboxAccountId int

// MailboxAccount holds data about ....
type MailboxAccount struct {
	//  id MailboxAccountId ?

	// login is the email
	login string

	// password is password for accessing the mailbox
	password string

	active bool

	// last error happened during loading
	lastError string

	// uid of the last message that was processed
	lastUid Uid

	createdAt Time
	updatedAt Time

	mailbox Mailbox

	// future TODOs:
	// - ignoreOwnLetters - ignore letters sent from `login` to `login`
	// - removeLoadedMessages - remove messages from mailbox after loading
}

// TODO - Or to pkg/types?
type Mailbox struct {
	host     string
	port     string
	protocol MailboxProtocol
	useSSL   bool
}

func (m Mailbox) Url() string {
	return "https://" + "s1" + "s2"
}
