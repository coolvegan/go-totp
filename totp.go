package main

import (
	"crypto/hmac"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"time"

	"golang.org/x/crypto/sha3"
)

type TotpI interface {
	FourDigits() string
	SixDigits() string
	EightDigits() string
	FourWords() ([]string, error)
}

type Totp struct {
	Secret   string
	Interval int64
	Wordlist WordlistI
}

func (t Totp) FourDigits() string {
	result, _ := generate(t.Secret, t.Interval, 4)
	return result
}
func (t Totp) SixDigits() string {
	result, _ := generate(t.Secret, t.Interval, 6)
	return result
}

func (t Totp) EightDigits() string {
	result, _ := generate(t.Secret, t.Interval, 8)
	return result
}

func (t Totp) FourWords() ([]string, error) {
	if t.Wordlist == nil {
		return nil, fmt.Errorf("WordList not defined in Totp")
	}

	number, err := strconv.Atoi(t.EightDigits())
	if err != nil {
		return nil, err
	}
	idx, err := getIndexInWordList(int(number))
	if len(idx) != 4 {
		return nil, fmt.Errorf("We need four indexes")
	}
	if err != nil {
		return nil, err
	}
	var result []string
	wl, err := t.Wordlist.Loadlist()
	if err != nil {
		return nil, err
	}
	for _, i := range idx {

		result = append(result, wl[i])
	}
	return result, nil

}

// TOTP generiert ein zeitbasiertes Einmalpasswort.
func generate(secret string, interval int64, digits int) (string, error) {
	if len(secret) == 0 {
		return "", fmt.Errorf("Please enter a secret!")
	}
	// 1. Base32-decodierten geheimen Schlüssel dekodieren
	if digits > 8 {
		return "", fmt.Errorf("Handling totp codes greater 8 is not implemented")
	}

	// 2. Berechne die Schrittnummer (aktueller Unix-Zeitstempel / Intervall)
	counter := time.Now().Unix() / interval

	// 3. Konvertiere die Schrittnummer in einen Byte-Array (Big Endian)
	// 8x8 Bits = 64 Bits
	var counterBytes [8]byte
	binary.BigEndian.PutUint64(counterBytes[:], uint64(counter))

	h := hmac.New(sha3.New256, []byte(secret))
	h.Write(counterBytes[:])
	hash := h.Sum(nil)

	// 5. Verwende Dynamic Truncation, um eine 8-stellige Ziffernfolge zu extrahieren
	offset := hash[len(hash)-1] & 0x0F
	code := (int64(hash[offset]&0x7F) << 32) |
		(int64(hash[offset+1]&0xFF) << 24) |
		(int64(hash[offset+2]&0xFF) << 16) |
		(int64(hash[offset+3]&0xFF) << 8) |
		(int64(hash[offset+4] & 0xFF))

	// 6. Modulo-Operation, um die Ziffern auf die gewünschte Länge zu beschränken
	otp := code % int64(pow10(digits))
	return fmt.Sprintf("%0*d", digits, otp), nil
}

// Hilfsfunktion zur Berechnung von 10^digits
func pow10(n int) int32 {
	result := int32(math.Pow10(n))
	return result
}

func extractPairOf32BitData(number uint) (int, int, int, int, error) {
	if int(math.Log10(float64(number)))+1 > 8 {
		return 0, 0, 0, 0, fmt.Errorf("Max supported positions is eight.")
	}
	//number ist immer achtstellig
	//extrahieren der ersten beiden stellen
	//1      1    2    2    3   3    4     4
	//10^7 10^6 10^5 10^4 10^3 10^2 10^1 10^0
	position1und2 := number / 1000000 % 100
	position3und4 := number / 10000 % 100
	position5und6 := number / 100 % 100
	position7und8 := number % 100
	return int(position1und2), int(position3und4), int(position5und6), int(position7und8), nil
}

// Get Index between 0 and WordList (4999)
func getIndexInWordList(number int) ([]int, error) {
	if int(math.Log10(float64(number)))+1 > 8 {
		return nil, fmt.Errorf("Max supported positions is eight.")
	}
	number1 := number / 1000000 % 100
	number2 := number / 10000 % 100
	number3 := number / 100 % 100
	number4 := number % 100

	var hash []byte = make([]byte, 8*32)
	result := ((((number1 * 1000) ^ number2) * number3) + 1) / number4
	str := strconv.Itoa(result)
	sha3.New224()
	//Generating PseudoRandom Values out of 256 Bit Sha3.224
	sha3.ShakeSum256(hash, []byte(str))
	x1 := int(hash[12]) << 6
	x2 := int(hash[16]) << 4
	x3 := int(hash[18]) << 2
	x4 := int(hash[24])

	x5 := int(hash[8]) << 10

	r2 := x1 | x2 | x3 | x4
	r3 := x1 + x2 + x3 + x4
	r4 := x1 + x2*x5*x4

	var resultArr []int
	resultArr = append(resultArr, (result^r2)%WORD_LIST_SIZE)
	resultArr = append(resultArr, (result*r3)%WORD_LIST_SIZE)
	resultArr = append(resultArr, (result+(x5&r3))%WORD_LIST_SIZE)
	resultArr = append(resultArr, (result*r4)%WORD_LIST_SIZE)
	return resultArr, nil
}
