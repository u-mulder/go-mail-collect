package mailbox_account

type MailboxAccount struct {
	login string
	password string
	mailbox struct {
		Address string 
		Port string 
		Protocol string 
	}
	// useSsl bool
	// removeLetters bool
	currentUid int
	// lastRun time
}

// TODO fill more fields
func New(string login, password, int currentUid) Account {
	return Account{
		// login: login,
		// password: password,
		// currentUid: currentUid,
	}
}

// func (a Account) Login() string {
// 	return a.login
// }

// func (a Account) Password() string {
// 	return a.password
// }

// func (a Account) CurrentUid() int {
// 	return a.currentUid
// }