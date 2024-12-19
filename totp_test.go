package main

import (
	"testing"
	"time"
)

func TestTOTPLengthOfEight(t *testing.T) {

	code, _ := generate("mysecret", 30, 8)
	if len(code) != 8 {
		t.Errorf("totp code must have length of 8 digits")
	}
}
func TestTOTPUnsupportedLength(t *testing.T) {

	rv, err := generate(string("asecretpassword"), 30, 9)
	if rv != "" && err.Error() != "Handling totp codes greater 8 is not implemented" {
		t.Errorf("Longer Totp Keys than 8 are not supported. Must return error message. \n%s", err.Error())
	}
}
func TestTOTPTimeIntervalDifferences(t *testing.T) {

	secret := "moresecret0391"
	rv, _ := generate(string(secret), 4, 8)
	time.Sleep(time.Duration(time.Millisecond * 600))
	rv2, _ := generate(string(secret), 4, 8)
	time.Sleep(time.Duration(time.Millisecond * 600))
	rv3, _ := generate(string(secret), 4, 8)
	if rv != rv2 && rv2 != rv3 {
		t.Errorf("Must have the same totp result in an interval of zwo seconds")

	}
}

func TestTotpInterfaceFour(t *testing.T) {
	var totp TotpI = Totp{Secret: "mySecret", Interval: 30, Wordlist: &Wordlist{filename: WORD_LIST_FILE_NAME}}

	r1 := totp.FourDigits()
	r2, _ := generate("mySecret", 30, 4)

	if r1 != r2 {
		t.Errorf("Error in TotpI 4")
	}
}
func TestTotpInterface(t *testing.T) {
	var totp TotpI = Totp{Secret: "12949mySecret", Interval: 30}

	r1 := totp.EightDigits()
	r2, _ := generate("12949mySecret", 30, 8)

	if r1 != r2 {
		t.Errorf("Error in TotpI 8")
	}
}

func TestTotpInterfaceDiffSecret(t *testing.T) {
	var totp TotpI = Totp{Secret: "AnotherSecret", Interval: 30}

	r1 := totp.EightDigits()
	r2, _ := generate("12949mySecret", 30, 8)

	if r1 == r2 {
		t.Errorf("Different Secrets should never give the same result")
	}
}

func TestZeroSecret(t *testing.T) {
	_, err := generate("", 30, 8)

	if err.Error() != "Please enter a secret!" {
		t.Errorf("Secret should never be empty")
	}
}

func TestPow10(t *testing.T) {
	type pair struct {
		X int
		Y int32
	}

	var exp = [5]pair{{1, 10}, {2, 100}, {3, 1000}, {4, 10000}, {5, 100000}}

	for _, i := range exp {
		y2 := pow10(i.X)
		if i.Y != y2 {
			t.Errorf("Error: Must be %d is %d", i.X, y2)
		}
	}

}

func TestExractPairOf32BitData(t *testing.T) {
	a, b, c, d, _ := extractPairOf32BitData(11223344)
	if a != 11 {
		t.Errorf("first must be 11 but is %d", a)
	}
	if b != 22 {
		t.Errorf("first must be 22 but is %d", b)
	}
	if c != 33 {
		t.Errorf("first must be 33 but is %d", c)
	}
	if d != 44 {

		t.Errorf("first must be 44 but is %d", c)
	}
}
func TestExractPairOf32BitData2(t *testing.T) {
	a, b, c, d, _ := extractPairOf32BitData(993344)
	if a != 00 {
		t.Errorf("first must be 00 but is %d", a)
	}
	if b != 99 {
		t.Errorf("first must be 99 but is %d", b)
	}
	if c != 33 {
		t.Errorf("first must be 33 but is %d", c)
	}
	if d != 44 {

		t.Errorf("first must be 44 but is %d", c)
	}
}
func TestExractPairOf32BitData3(t *testing.T) {
	a, b, c, d, _ := extractPairOf32BitData(0)
	if a != 00 {
		t.Errorf("first must be 00 but is %d", a)
	}
	if b != 00 {
		t.Errorf("first must be 00 but is %d", b)
	}
	if c != 00 {
		t.Errorf("first must be 00 but is %d", c)
	}
	if d != 00 {

		t.Errorf("first must be 00 but is %d", c)
	}
}
func TestExractPairOf32BitData4(t *testing.T) {
	_, _, _, _, err := extractPairOf32BitData(1234567890)
	if err.Error() != "Max supported positions is eight." {
		t.Errorf("Max Number allowed until 8 positions")
	}
}

func TestExractWords(t *testing.T) {
	// var totp TotpI = Totp{Secret: "mySecret", Interval: 30, Wordlist: &Wordlist{filename: WORD_LIST_FILE_NAME}}
}
