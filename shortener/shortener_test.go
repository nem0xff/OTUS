package shortener

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

var mystruct *Shorten

var src = rand.NewSource(time.Now().UnixNano())

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyz/ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = int(63 / letterIdxBits)
)

func init() {
	var err error
	mystruct, err = NewShorten(true, true) // первый параметр устанавливает значение дедупликации по умолчанию, второй параметр включает валидацию URL
	if err != nil {
		panic("Ошибка создания объекта")
	}
}

func TestAddLinks(t *testing.T) {

	for j := 0; j < 100000; j++ { // Добавляем 100 000 URL'ов
		_, err := mystruct.Shorten(makePseudoURL())
		if err != nil {
			t.Log("Был получен не валидный URL")
		}
	}

	PrintMemUsage()
}

func TestShorten(t *testing.T) {

	myURL := makePseudoURL()
	shortLinkKey, err := mystruct.Shorten(myURL)
	if err != nil {
		t.Error("Был получен не валидный URL")
	}
	if shortLinkKey == "" {
		t.Error("При добавлении элемента вернулся пустой ключ")
	}
}

func TestResolve(t *testing.T) {
	myURL := makePseudoURL()
	shortLinkKey, err := mystruct.Shorten(myURL)
	if err != nil {
		t.Error("Был получен не валидный URL")
	}
	if myURL != mystruct.Resolve(shortLinkKey) {
		t.Error("Вернулся не правильный результат")
	}
	if mystruct.Resolve("4Ydo3") != "" {
		t.Error("Вернулась не пустая строка при запросе несуществующего элемента")
	}
}

func TestBase10ToNewBase(t *testing.T) {
	if base10ToNewBase(73456487) != "4Ydo3" {
		t.Error("Неверный результат преобразования в новую систему счисления")
	}
}

func TestNewBaseToBase10(t *testing.T) {
	// 4Ydo3 - 73456487
	if newBaseToBase10("4Ydo3") != 73456487 {
		t.Error("Неверный результат преобразования в десятичную систему счисления")
	}
}

//Todo: починить дедупликацию
func TestDeduplication(t *testing.T) {
	myURL := makePseudoURL()
	shortLinkKey1, err := mystruct.Shorten(myURL)
	if err != nil {
		t.Error("Был получен не валидный URL")
	}
	shortLinkKey2, err := mystruct.Shorten(myURL)
	if err != nil {
		t.Error("Был получен не валидный URL")
	}
	if shortLinkKey1 != shortLinkKey2 {
		t.Error("Ошибка дедупликация не сработала")
	} else {
		t.Log("Дедупликация прошла успешно")
	}

}

func TestValidatorOfURL(t *testing.T) {
	if mystruct.enableValidation {

		errorPath := make([]string, 7)
		errorPath[0] = "http//testsite/"
		errorPath[1] = "http:/testsite/"
		errorPath[2] = "//testsite/"
		errorPath[3] = "http:\\testsite/"
		errorPath[4] = "testsite"
		errorPath[5] = "htp://testsite/"
		errorPath[6] = "http:///testsite/"

		for _, myURL := range errorPath {

			_, err := mystruct.Shorten(myURL)
			if err == nil {
				t.Error("Валидация пропустила адрес: " + myURL)
			}
		}
	}

}

func makePseudoURL() string {
	var lenght int
	for lenght = 0; lenght < 10; {
		lenght = int(src.Int63() & letterIdxMask) // from 1 to 63
	}
	return "http://random" + genString(lenght)
}

func genString(n int) string {
	b := make([]byte, n)

	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}

		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
