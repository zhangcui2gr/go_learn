package main

import (
	"bytes"
	"fmt"
	"strings"
)

const BLOG = "http://www.baidu.org/"

func initStrings(N int) []string {
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = BLOG
	}
	return s
}

func initStringi(N int) []interface{} {
	s := make([]interface{}, N)
	for i := 0; i < N; i++ {
		s[i] = BLOG
	}
	return s
}

func stringPlus(p []string) string {
	var s string
	l := len(p)
	for i := 0; i < l; i++ {
		s += p[i]
	}
	return s
}

func stringFmt(p []interface{}) string {
	return fmt.Sprint(p...)
}

func stringJoin(p []string) string {
	return strings.Join(p, "")
}

func stringBuffer(p []string) string {
	var b bytes.Buffer
	l := len(p)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}

func stringBuilder(p []string) string {
	var b strings.Builder
	l := len(p)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}
