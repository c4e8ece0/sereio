package doc

// check document structure and shows tags intersection

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
