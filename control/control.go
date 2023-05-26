package control

import "account-book/control/acctbook"

func NewAccountBook() acctbook.IAcctBook {
	ab := acctbook.New()
	return ab
}
