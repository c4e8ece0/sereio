// Package se stores profiles with features of most popular search engines

package se

import (
	tags "sere.io/tag"
)

// Funcs and interface need for more concistense

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

var Yandex = Profile{
	Name:                 "Yandex",
	RobotsUserAgent:      "yandex",
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
	RobotsUserAgent:      "googlebot",
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
	RobotsUserAgent:      "yahoo",
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
