package fkopts

import "golang.org/x/text/language"

type LangWrapper struct {
	TagList []language.Tag
}

func (l *LangWrapper) UseLang(lang string) {
	tag, err := language.Parse(lang)
	if err != nil {
		panic(err)
	}

	l.TagList = append(l.TagList, tag)
}

func (l *LangWrapper) GetFirstLang() language.Tag {
	if l == nil || len(l.TagList) == 0 {
		return language.English
	}
	return l.TagList[0]
}
