package level1tasks

func createHugeString(size int) string {
	return string(make([]byte, size))
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10) // создаем большую строку (1 кб)
	justString = string([]rune(v)[:100]) // используем rune, чтобы избежать проблем с многобайтными символами
	// justString - теперь не ссылается на v, что позволит GC очистить память большой строки и избежать утечки памяти
}

func Task15() {
	someFunc()
}