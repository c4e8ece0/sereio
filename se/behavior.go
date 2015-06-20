// Package se stores profiles with features of most popular search engines

package se

import (
	tags "sere.io/tag"
)

// Funcs and interface need for more concistense

type Behavior struct {
	IsSubExternal    bool // Считать поддомены внешними доменами
	IsPortExternal   bool // Считать сайт на другом порту внешним сайтом
	IgnoreIndexable  bool // Игнорировать правила индексации из мета-тегов
	IgnoreFollowable bool // Игнорировать правила прохода из мета-тегов
	UseTitleAsAnchor bool // Учитывать title= как анкор ссылки (продолжение? вторую ссылку?)
	UseAltAsText     bool // Подменять картинку текстом
}

func ShowMisses(sp SeProf) {
}

type Profile struct {
	Name string

	MaxBytes int // max bytes
	MaxLines int // max lines

	// Robots.txt fields
	// UA = UserAgent
	RobotsUserAgent          map[string]string // type(default, image, ad):name
	RobotsAllowExtension     bool
	RobotsCrawlDelay         bool
	RobotsCleanParam         bool
	RobotsCleanParamWide     bool // общая на весь файл
	RobotsHost               bool
	RobotsAllow              bool
	RobotsSitemap            bool
	RobotsMultipleUA         bool
	RobotsFirstSlash         bool
	RobotsWildcart           bool
	RobotsVoidIgnore         bool
	RobotsDisallowIgnoreRoot bool
	RobotsLenAlarm           bool // Сообщать об одинаковой длине конфликтующих суффиксов у allow и disallow
	RobotsLenAllow           bool // Приоритет у Allow при конфликте длинн суффиксов у allow и disallow
	RobotsEmptyBefore        bool // Требовать пустую строку перед User-agent

	// Tags processing
	TagCrop  []tag.Tag
	TagBlock []tag.Tag
	TagSpace []tag.Tag
}

var Default = Profile{
	Name:                 "Sere.io",
	RobotsUserAgent:      "sereio user-agent",
	RobotsAllowExtension: true,
	RobotsCleanParam:     true,
	RobotsCleanParamWide: true,
	RobotsHost:           true,
	RobotsAllow:          true,
	RobotsSitemap:        true,
	RobotsWildcart:       true,
	RobotsVoidIgnore:     true,
	RobotsLenAllow:       true,
	RobotsEmptyBefore:    true,
	RobotsSitemap:        true,
}
var Yandex = Profile{
	Name:                 "Yandex",
	RobotsUserAgent:      "yandex user-agent",
	RobotsAllowExtension: true,
	RobotsCleanParam:     true,
	RobotsCleanParamWide: true,
	RobotsHost:           true,
	RobotsAllow:          true,
	RobotsSitemap:        true,
	RobotsWildcart:       true,
	RobotsVoidIgnore:     true,
	RobotsLenAllow:       true,
	RobotsEmptyBefore:    true,
	RobotsSitemap:        true,
}

var Google = Profile{
	Name:                 "Google",
	RobotsUserAgent:      "googlebot user-agent",
	RobotsAllowExtension: true,
	RobotsCleanParam:     true,
	RobotsCleanParamWide: true,
	RobotsHost:           true,
	RobotsAllow:          true,
	RobotsSitemap:        true,
	RobotsWildcart:       true,
	RobotsVoidIgnore:     true,
	RobotsLenAllow:       true,
	RobotsEmptyBefore:    true,
	RobotsSitemap:        true,
}

var Yahoo = Profile{
	Name:                 "Yahoo",
	RobotsUserAgent:      "yahoo user-agent",
	RobotsAllowExtension: true,
	RobotsCleanParam:     true,
	RobotsCleanParamWide: true,
	RobotsHost:           true,
	RobotsAllow:          true,
	RobotsSitemap:        true,
	RobotsWildcart:       true,
	RobotsVoidIgnore:     true,
	RobotsLenAllow:       true,
	RobotsEmptyBefore:    true,
	RobotsSitemap:        true,
}

var Mail = Profile{
	Name:                 "Mail.ru",
	RobotsUserAgent:      "mail.ru",
	RobotsAllowExtension: true,
	RobotsCleanParam:     true,
	RobotsCleanParamWide: true,
	RobotsHost:           true,
	RobotsAllow:          true,
	RobotsSitemap:        true,
	RobotsWildcart:       true,
	RobotsVoidIgnore:     true,
	RobotsLenAllow:       true,
	RobotsEmptyBefore:    true,
	RobotsSitemap:        true,
}
