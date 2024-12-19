package main

import (
	"fmt"
)

func main() {
	// Beispielhafter geheimer Schlüssel (Base32-kodiert)
	secret := "JVQXEY3PJNUXI5DFNQYTA===" // Dieser Schlüssel muss Base32 kodiert sein

	// TOTP-Generierung mit 30-Sekunden-Schritten und 8 Ziffern
	code, err := generate(secret, 30, 8)
	if err != nil {
		fmt.Println("Fehler:", err)
		return
	}

	fmt.Println("TOTP Code:", code)

	var t = &Totp{Secret: "102jf0j023jf0jf023f003f0hscdjljsdlkjfdslj299", Interval: 30, Wordlist: &Wordlist{filename: WORD_LIST_FILE_NAME}}
	v, err := t.FourWords()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}
