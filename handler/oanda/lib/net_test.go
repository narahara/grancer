package lib

import (
	"log"
	"testing"
)

func GetTest(t *testing.T) {
	b, n, err := Get("https://api-fxtrade.oanda.com/v1/accounts", "a768709c9055b58ec5ac9055ca406d35-d6440968f9154452ff16dda859788d48")
	if err != nil {
		t.Error(err)
	}
	log.Println(n)
	log.Println(string(b))
}
