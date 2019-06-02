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
	mystruct, err = NewShorten(true) // первый параметр устанавливает значение дедупликации по умолчанию
	if err != nil {
		panic("Ошибка создания объекта")
	}
}

func TestAddLinks(t *testing.T) {

	for j := 0; j < 100000; j++ { // Добавляем 100 000 URL'ов
		_ = mystruct.Shorten(makePseudoURL())
	}

	PrintMemUsage()
}

func TestShorten(t *testing.T) {

	myURL := makePseudoURL()
	shortLinkKey := mystruct.Shorten(myURL)
	if shortLinkKey == "" {
		t.Error("При добавлении элемента вернулся пустой ключ")
	}
}

func TestResolve(t *testing.T) {
	myURL := makePseudoURL()
	shortLinkKey := mystruct.Shorten(myURL)
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

func TestDeduplication(t *testing.T) {
	mystruct.SetDeduplicationStatus(true) // включаем дедупликацию если не включена
	myURL := makePseudoURL()
	shortLinkKey1 := mystruct.Shorten(myURL)
	shortLinkKey2 := mystruct.Shorten(myURL)
	if shortLinkKey1 != shortLinkKey2 {
		t.Error("Ошибка дедупликация не сработала")
	} else {
		t.Log("Дедупликация прошла успешно")
	}
	mystruct.SetDeduplicationStatus(false)
	shortLinkKey1 = mystruct.Shorten(myURL)
	shortLinkKey2 = mystruct.Shorten(myURL)
	if shortLinkKey1 == shortLinkKey2 {
		t.Error("Ошибка два одинаковых ключа с отключенной дедупликацией")
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
