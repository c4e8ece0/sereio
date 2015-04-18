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

// FROM PHP:
// public function HttpEquiv(){}
// public function HEvsMeta(){} // поиск мет с неправильным названием
// public function Crop(){} // обрезка документа по какому-то признаку (<body>, <html>)
// public function Report(){} // Отчёт об ошибках и добавлениях в код

type Behavior struct {
	IsSubExternal    bool // Считать поддомены внешними доменами
	IsPortExternal   bool // Считать сайт на другом порту внешним сайтом
	IgnoreIndexable  bool // Игнорировать правила индексации из мета-тегов
	IgnoreFollowable bool // Игнорировать правила прохода из мета-тегов
	UseTitleAsAnchor bool // Учитывать title= как анкор ссылки (продолжение? вторую ссылку?)
	UseAltAsText     bool // Подменять картинку текстом
}

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

type Factor struct {
	name  string
	value float32
	ok    bool // ?
}

func (f *Factor) Set(string) error {
	//... проверять имя?
}

type Stat struct {
	name string
	rows int
	cols int
	arr  []int
}

func (s *Stat) Show() map[Factor]float32 {
	// Набор данных для эксорта в Qolumn
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

const (
	BLOCK = 1 << iota
	PARAGRAPH
	PASSAGE
	SENTENCE
	ALL = BLOCK | PARAGRAPH | PASSAGE | SENTENCE
)

// Индекс слов документе
func (s *Doc) makeindex(depth int) {
}

func (s *Doc) extract_vcard() {
}

func (s *Doc) extract_schemaorg() {
}

func (s *Doc) extract_og() { // ++ etc microformats
}

// Выдернуть ссылки из документа
func (s *Doc) extractsingles() {
	// -------------------------
	// auditor/__include.php
	// -------------------------
	// function easysingletag($content, $tname, $aname, $avalue = '([^"]*)', $atarget, $complex = '') {
	// 	$r = array();
	// 	preg_match_all('/<' . $tname . '[^>]*' . preg_quote($aname) . '="' . preg_quote($avalue) . '"[^>]*>/is' , $content, $arr);

	// 	$r['val']  = array('-1' => '', '0' => '');
	// 	$r['num']  = isset($arr[0]) ? count($arr[0]) : 0;
	// 	$r['arr']  = array();
	// 	$r['kv']   = array();

	// 	if($r['num']) {
	// 		$i = 0;
	// 		foreach($arr[0] as $k=>$v) {
	// 			preg_match('/<' . $tname . '[^>]*' . preg_quote($atarget) . '="([^"]*)"[^>]*>/', $v, $p);
	// 			$r['val'][$i]  = '';
	// 			if(isset($p[1])){
	// 				$r['val'][$i]  = $p[1];
	// 			}

	// 			if($complex) {
	// 				$t = $r['val'][$i];
	// 				$t = explode($complex, $t . $complex);
	// 				unset($t[count($t) - 1]);

	// 				foreach($t as $a=>$b) {
	// 					$b = trim($b);
	// 					$r['arr'][]  = $b;
	// 					$r['kv'][$b] = $b;
	// 				}
	// 			}
	// 			$i++;
	// 		}
	// 	}
	// 	return $r;
	// }
}

const (
	FOLLOW = 1 << iota
	NOFOLLOW
	INTERNAL // for Anchors()
	EXTERNAL // for Anchors()
	NOMATTER = FOLLOW | NOFOLLOW | INTERNAL | EXTERNAL
)

// Список внутренних ссылок
func (s *Doc) Internal(flags int) []Link {
	// Find links
	// Return only internal
}

// Список внешних ссылок
func (s *Doc) External(flags int) []Link {
	// Find links
	// Return only external
}

// Выбрать только тексты ссылок
func (s *Doc) Anchors(flags int) []string {
}

// --------------------------------------------------------------------------
// Выбрать только заголовки
func (s *Doc) Headers() []string {

}

// --------------------------------------------------------------------------

const (
	TITLE
	ROBOTS
	NOARCHIVE
	NOODP
	NOYACA
	CANONICAL
	KEYWORDS
	DESCRIPTION
	BASE
	FAVICON
	CONTENTTYPE
)

// Проверка ошибок для метатегов:
// 1. Единственность
// 2. Расположение в секции head
// ...
func (s *Doc) MetaErrors(flags int) []string {
}

// Выбрать только заголовки
func (s *Doc) Title() []string {
}

// Метки управления индексацией и присутствием в каталогах
// (name=yaca, noodp..., value=all, none, index, follow, no&)
func (s *Doc) Robots() []string {
}

// Canonical <link rel-canonical> из <head/>
func (s *Doc) Canonical() string {
}

// <meta name="keywords">
func (s *Doc) Keywords() string {
}

// <meta name="description">
func (s *Doc) Description() []Tag {
}

// Получение <base-href=> из <head/>
func (s *Doc) Base() []Attr {
	return true
}

// Favicon
func (s *Doc) Favicon() string {

}

// <meta name="content-type" content="..."> => mime + charset
func (s *Doc) ContentType() string {
}

// Проверка корректности структуры документа
// => <!doctype><html><head><title></title></head><body></body></html>
// Errors:
// - Текст до <head>
// - Текст после <head>
// - Отсутствие и неединственность открывающих и закрывающих основных тегов
const (
	HEAD = 1 << iota
	HTML
	TITLE
	BODY
)

func (s *Doc) CheckStuct() {

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

// Описание пути в ссылке или ресурсе
type Path struct {
	External  bool
	Relative  bool // first "/" not exists
	DotPrefix int8 // number of dots in prefix,
}

// Структура для описания документа
type Link struct {
	Pre              string
	Anchor           string
	Post             string
	Scheme           string // http, https, tel, javascript, mailto, call, any...
	Image            []Attr // Index of image + details
	Nofollow         bool
	NofollowExternal bool
	Rel              string
	path             string
	frag             string
}

// Проверка наличия аттрибута rel=nofollow
func (a *Link) Nofollow() bool {
	return a.Nofollow || a.NofollowExternal
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

// 4 scorpion example
func ParseHtml(r io.Reader) {
	t := html.NewTokenizer(r)
	t.AllowCDATA(true)

	for {
		// t.NextIsNotRawText() // To find later, else text goes with tags: <title>s<b>s</b>ss</title>
		tt := t.Next()
		if tt == html.ErrorToken {
			return
		}
		u, _ := t.TagName()
		s := t.Text()
		for {
			v, x, e := t.TagAttr()

			tag := string(u)
			attr := string(v)
			value := string(x)
			text := string(s)

			if tag == "script" && attr == "src" {
				if strings.HasPrefix(value, "//") || strings.HasPrefix(value, "http://") || strings.HasPrefix(value, "https://") {
					fmt.Printf("%v %v = %v|\n", tag, attr, value)
				}
			}
			if strings.TrimSpace(text) != "" {
				fmt.Printf("=== %v=%v|\n", tag, strings.TrimSpace(text))
			}

			if e == false || tag == "" {
				break
			}
		}
	}
}
