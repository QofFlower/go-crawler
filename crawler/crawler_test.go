package crawler

import "testing"

// @Author: Hananoq
func TestGetHttpHtmlContent(t *testing.T) {
	_, err := GetHttpHtmlContent("http://baidu.com", "", `document.querySelector("body")`)
	if err != nil {
		t.Fatal("TestGetHttpHtmlContent: Not expected result or error encountered")
	}
}

func TestGetSpecialData(t *testing.T) {
	html, _ := GetHttpHtmlContent("http://baidu.com", "", `document.querySelector("body")`)
	_, err := GetSpecialData(html, []*SelectorData{
		{JobTypeName, ".selector_one"},
		{JobTypeBaseAndType, ".selector_two"},
		{JobTypeDescription, ".selector_three"}}...)
	if err != nil {
		t.Fatal("TestGetSpecialData: Not expected result or error encountered")
	}
}

func TestJoinSelector(t *testing.T) {
	selector := JoinSelector([]*SelectorData{
		{JobTypeName, ".selector_one"},
		{JobTypeBaseAndType, ".selector_two"},
		{JobTypeDescription, ".selector_three"}}...)
	want := ".selector_one,.selector_two,selector_three"
	if selector != want {
		t.Fatalf("TestGetHttpHtmlContent: Not expected result, want: %s, got: %s", want, selector)
	}
}
