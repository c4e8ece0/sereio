//
package sere

import (
	"io"

	"golang.org/x/net/html"
)

// TODO: Полный рерайт? <- Нужна проверка пробелов между атрибутом и значением
// TODO: Define canonical behavior in diff SE
// TODO: Не забывать про множественные варианты тегов и содержаний для мета-тегов
// TODO: Не забывать про "nofollow external" и "dofollow-wtf?", "next" и "prev"
// TODO: Favicon может быть двух видов - найти нужный
// TODO: В Path.Scheme нужен свой разборщик на каждый протокол
// No space between attributes. mortgage_submit"><button type="button"class="
// <select></select> ? <- Нужна проверка на отсутствие необходимых элементов
// Doc.HasMicroFormats == [class=vcard],[meta^=og:...],[<attr>==schema.org]
// Рендер документа и полей стиля где и когда?
// HOWTO: append to Qolumn? A nahua?
//
// Sere.Register() + Module + mods - это же View ?!!!!

// FROM PHP:
// public function HttpEquiv(){}
// public function Crop(){} // обрезка документа по какому-то признаку (<body>, <html>)
// public function Report(){} // Отчёт об ошибках и добавлениях в код

// Создание нового объекта
func New(r io.Reader, f sere.Targeter) *Doc {
	return &Doc{}
}

// --------------------------------------------------------------------------
// --------------------------------------------------------------------------
// --------------------------------------------------------------------------

// Структура для описания документа
type Doc struct {
	blocks     []Block
	paragraphs []Paragraph // is br-br == p in Behavior?
	sentences  []Sentence
	metahe_robots

	HasMicroFormats bool // Maybe has microformats flag

	links []Link
	words []Word
}

// Создание представления из документа:
// - Мешок слов
// - Мешок слов с нормализацией по AOT|Yandex.HL|Google.HL|Yandex.Colloc||...
// - Блоки + слова []->[]
// - Блоки + пассажи []->[]->[]
// - Блоки|пассажи + слова
// - Полные классы стилей
// - block vs inline?
// - Приведение классов стилей к однословным
// - ... Теги, меты, бу бу бу, все поля и аттрибуты
func (s *Doc) View(func(Tr) Stat) {
	// ...
}

// Выдернуть все ссылки из документа
func (s *Doc) extractlinks() {
}

// Нужны же имена классов для поиска авторства сайтов?!
func (s *Doc) extractstructure() {
}

// Статистика по тегам, классам, словам
func (s *Doc) makestats() {
}

// // Module description
// type Module interface {
// 	Name() string
// 	Instance() int
// 	Require() []ModuleName
// }

// var mods map[string][]ModuleWhaaaaat

// func init() {
// 	mods = make(map[string][]ModuleWhaaaaat)
// }

// //
// func Register(s Module) {
// 	mods[s.Name()] = Module
// }

//
const (
	BLOCK = 1 << iota
	PARAGRAPH
	PASSAGE
	SENTENCE
	TITLE_INCLUDE
	ALL = BLOCK | PARAGRAPH | PASSAGE | SENTENCE
)

// --------------------------------------------------------------------------
// Выбрать только заголовки
func (s *Doc) Headers() []string {

}

// --------------------------------------------------------------------------

// Проверка разрешения на индексацию
func (s *Doc) Indexable() bool {
	return true
}

// Проверка разрешения на проход по ссылкам
func (s *Doc) Followable() bool {
	return true
}

// --------------------------------------------------------------------------

// Описание тега
type Attr struct {
	Href       []string
	HasHref    bool
	Src        []string
	HasSrc     bool
	Title      []string
	HasTitle   bool
	Alt        []string
	HasAlt     bool
	Rel        []string
	HasRel     bool
	Content    []string
	HasContent bool
}

// --------------------------------------------------------------------------

// <a><a><a> ?
const (
	IN_DOC iota
	IN_BLOCK
	IN_A
)
