package main

import "testing"

const prefixEnglish = "hello, "
const prefixSpanish = "hola, "

func hello(str string, language string) string {
	if str == "" {
		return "hello, world"
	}
	prefix := prefixEnglish
	switch language {
	case "English":
		prefix = prefixEnglish
	case "Spanish":
		prefix = prefixSpanish
	}
	return prefix + str
}

func TestHello(t *testing.T) {
	str1 := hello("world", "English")
	str2 := "hello, world"

	if str1 != str2 {
		t.Errorf("str1: %s, str2: %s", str1, str2)
	}
}
