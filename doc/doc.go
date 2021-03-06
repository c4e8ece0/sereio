// Package doc collect data about all fragments
package doc

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
// USE SOMEWHERE: https://github.com/fatih/pool

// FROM PHP:
// public function HttpEquiv(){}
// public function Crop(){} // обрезка документа по какому-то признаку (<body>, <html>)
// public function Report(){} // Отчёт об ошибках и добавлениях в код

var Set = struct{}{}

// Создание нового объекта
func New(r io.Reader) *Doc {
	return &Doc{src: ""}
}

// CYRUTF ADDITION
func utf1_main() {
	data1 := "ABC"
	fmt.Println(utf8.ValidString(data1)) //prints: true

	data2 := "A\xfeC"
	fmt.Println(utf8.ValidString(data2)) //prints: false
}

// DON'T FORGET NFC
func utf2_main() {
	data := "é"
	fmt.Println(len(data))                    //prints: 3
	fmt.Println(utf8.RuneCountInString(data)) //prints: 2
}

// --------------------------------------------------------------------------
// HTML.TOKENIZER Looses spaces of attrs (and text?) -- REQUIRE SOMETHING (html.Node?)
// --------------------------------------------------------------------------

// --------------------------------------------------------------------------
// --------------------------------------------------------------------------
// --------------------------------------------------------------------------

// Const maybe better? How much total needed?
// Options setting? Func? Flags? Field assign?
type ParamDel struct { // i.e. what to fetch?
}

// Структура для описания документа
func usage() {
	// Var 1
	todo = doc.New(html.body())
	todo.Ahref = true
	todo.ImgSrc = true
	todo.ScriptSrc = true
	todo.LinksSrc = true
	todo.IframeSrc = true
	t := todo.Parse() // return Qolumn = [*.View("name", data)]?

	// Var 2
	todo = doc.New(html.body())
	todo.Resources() // .attr
	todo.Images()    // .attr + PrepareImages(implode attrs by tag index)
	todo.Links()
	s := todo.Parse() // return Qolumn = [*.View("name", data)]?
}

type Doc struct {
	// blocks     []Block
	// paragraphs []Paragraph // is br-br == p in Behavior?
	// sentences  []Sentence
	// metahe_robots
	// links []Link
	// words []Word

	data map[int16]interface{}

	// Report bools
	Has struct {
		HTML         bool
		MicroFormats bool // Maybe has microformats flag
	}

	// Parser params
	Param struct {
		// Move it to helper-methods of doc.
		// ScriptSrc     bool
		// ScriptContent bool
		// AHref         bool
		// AFull         bool // new Link[] as result
		// CssClass      bool
		// CssStyle      bool
		// ImgAlt        bool
		// AttrTitle     bool
		// IframeSrc     bool
		// ImgSrc        bool
		// LinkSrc       bool
		// Title         bool
		// Text          bool
	}
}

// Упаковка документа для хранения
// Свой формат для каждого поля описания
// text = id(words[10e3+] global_ids, delims[>10e3] ids, signs[>10e3] ids)
func (d *Doc) Pack() io.Reader {
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
//
// Расширение слов, где и как должно происходить? SOA? Inline? DbSharing?
// - Транслитерации (N вариантов)
// - Переводы
// - Аббревиатуры
// - Синонимы
// - Расширения по типам речи?
// - Джойны
// - Коллокации
func (s *Doc) View(func(Tr) Qol) {
	// ...
}

// Produces classic bag of words
func (s *Doc) BagOfWords() {
	// ...
}

// Produces vector view (i.e. with coordinates)
func (s *Doc) ListOfWords() {
	// ...
}

// Produces list of block-vectors (i.e. array of coordinates in document fragments block + paragraph + sentence + word)
func (s *Doc) ViewListOfBags() {
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

// Just good practice
func withLockContext(fn func()) {
	mu.Lock
	defer mu.Unlock()

	fn()
}
func foo() {
	withLockContext(func() {
		// foo related stuff
	})
}

// same + db
func withDBContext(fn func(db DB)) error {
	// get a db connection from the connection pool
	dbConn := NewDB()

	return fn(dbConn)
}
func boo() {
	withDBContext(func(db *DB) error {
		// foo related stuff
	})
}

// Safe multiple storages
type Storage interface {
	Delete(key string)
	Get(key string) string
	Put(key, value string)
}

func Delete(key string) {
	mu.Lock()
	delete(m, key)
	mu.Unlock()
}
func Put(key, value string) {
	mu.Lock()
	m[key] = value
	mu.Unlock()
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
	BLOCK = 1 << (iota + 1)
	PARAGRAPH
	PASSAGE
	SENTENCE
	TITLE_INCLUDE
	ALL = BLOCK | PARAGRAPH | PASSAGE | SENTENCE
)

// Getting sereio/attr objectject
func (s *Doc) Attr() *Attr {

}

// --------------------------------------------------------------------------
// Выбрать только заголовки
func (s *Doc) Headers() []string {
	// TODO doc.Tag().Headers()
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
	IN_DOC = iota + 1
	IN_BLOCK
	IN_A
)

// KISS THAT CRAP
func ParseLinks() { // []Attr
}

func ParseMeta() { // [name][]content
}

func ParseText() { // [block_id][passage_id][]word_id
}

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

//
func LoadFiles(dirpages string) map[int]string {
	doc := make(map[int]string)

	// Список урлов и содержание в них слов из подсветки
	arr, _ := ioutil.ReadDir(dirpages)
	for _, v := range arr {
		if strings.Contains(v.Name(), ".url") {
			continue
		}
		f, err := os.Open(dirpages + v.Name())
		check(err)
		defer f.Close()

		s, err := NewCyrReader(f)
		check(err)

		d, err := ioutil.ReadAll(s)
		check(err)

		if len(d) < 1 {
			continue
		}

		c, err := strconv.Atoi(v.Name())
		check(err)

		doc[c] = string(d)
		if len(doc) >= DOCLIMIT {
			break
		}
	}
	return doc
}

//
func NewCyrReader(r io.ReadSeeker) (io.Reader, error) {
	enc := DetermineCyrEncoding(r)
	return charset.NewReaderByName(enc, r)
}

//
func DetermineCyrEncoding(r io.ReadSeeker) string {
	r.Seek(0, 0)
	str, _ := ioutil.ReadAll(r)
	r.Seek(0, 0)
	s, _, _ := cyrutf.DetermineEncoding(str)
	a := string(s)
	if len(s) == 0 {
		_, p, e := charset.DetermineEncoding(str, "text/html") // works only on utf-8
		s = p
		fmt.Printf("2=%v %v\n", s, e)
	}
	if a == "windows-1252" {
		a = "windows-1251"
	}
	if a == "" {
		a = "utf-8" // in the name of universe
	}
	return a
}
