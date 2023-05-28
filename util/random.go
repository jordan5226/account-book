package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(strLen int) string {
	var sb strings.Builder
	alphabetSize := len(alphabet)

	for i := 0; i < strLen; i++ {
		c := alphabet[rand.Intn(alphabetSize)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Get random name in the random length
func RandomName() string {
	return RandomString(int(RandomInt(1, 20)))
}

// Get random user ID in the random length
func RandomUID() string {
	return RandomString(int(RandomInt(5, 10)))
}

// Get random password in the random length
func RandomPwd() string {
	return RandomString(int(RandomInt(8, 16)))
}

func RandomCurrency() string {
	arrCurrencies := []string{
		"TWD", "USD", "JPN", "ERU",
	}

	return arrCurrencies[rand.Intn(len(arrCurrencies))]
}

func RandomMoney() int64 {
	return RandomInt(0, 5000)
}

func RandomAccountName() string {
	arrAcct := []string{
		"Cash", "Credit Card", "Bank",
	}
	return arrAcct[rand.Intn(len(arrAcct))]
}

func RandomAccountIcon() string {
	arrIcon := []string{
		"cash.jpg", "creditcard.jpg", "bank.jpg",
	}
	return arrIcon[rand.Intn(len(arrIcon))]
}

func RandomProjectName() string {
	arrPrj := []string{
		"Personal", "Family", "Work", "Game",
	}
	return arrPrj[rand.Intn(len(arrPrj))]
}

func RandomProjectIcon() string {
	arrIcon := []string{
		"personal.jpg", "family.jpg", "work.jpg", "game.jpg",
	}
	return arrIcon[rand.Intn(len(arrIcon))]
}

func RandomStoreName() string {
	arrStore := []string{
		"7-11", "Family", "McDonald", "KFC", "ichiran",
	}
	return arrStore[rand.Intn(len(arrStore))]
}

func RandomStoreIcon() string {
	arrIcon := []string{
		"711.jpg", "family.jpg", "mcdonald.jpg", "kfc.jpg", "ichiran.jpg",
	}
	return arrIcon[rand.Intn(len(arrIcon))]
}

func RandomTypeName() string {
	arrType := []string{
		"Pay", "Income", "Transfer",
	}
	return arrType[rand.Intn(len(arrType))]
}

func RandomTypeIcon() string {
	arrIcon := []string{
		"pay.jpg", "income.jpg", "transfer.jpg",
	}
	return arrIcon[rand.Intn(len(arrIcon))]
}
