// Package meta fetch metatags from html io.Reader
package meta

// public function HE_vs_Meta(){} // поиск мет, которые должны быть "http-equiv", а не "name" и наоборот

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

var me = Sere.Module{
	name: "meta",
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

func init() {
	doc.Register("meta", me)
}

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
