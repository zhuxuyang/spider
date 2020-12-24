package biz

import (
	"log"
	"testing"
)

func TestParseSubjectID(t *testing.T) {
	log.Println(ParseSubjectID("https://book.douban.com/subject/1236778/"))
}
