package main

import (
	"testing"
)

func TestHello2(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, str1, str2 string) {
		t.Helper()
		if str1 != str2 {
			t.Errorf("str1: %s str2: %s", str1, str2)
		}
	}
	t.Run("say hello to people", func(t *testing.T) {
		str1 := hello("world", "Spanish")
		str2 := "hello, world"
		//if str1 != str2 {
		//	t.Errorf("str1: %s, str2: %s", str1, str2)
		//}
		assertCorrectMessage(t, str1, str2)
	})
	t.Run("say hello to people when people is empty", func(t *testing.T) {
		str1 := hello("", "Spanish")
		str2 := "hello, world"
		if str1 != str2 {
			t.Errorf("str1: %s, str2: %s", str1, str2)
		}
		//assertCorrectMessage(t, str1, str2)
	})
}
