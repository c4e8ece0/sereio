// Package url makes work on small details about url
//
// STANDART PACKAGE ALLREADY HAS MOST FUNCTIONS, HERE MUST BE ONLY EXTENSIONS
//
package url

// Function New() makes new object from string
func New(s string) (URL, error) {
}

// Function Plus() makes new object from current one and string addition
func Plus(c URL, n string) (URL, error) {
}

// NewVoid make void object
func NewVoid() {
	return &URL{
		Path: '/',
		Dir:  '/',
	}
}

// Детали представления имени
type Host struct {
	Host      string   // полное имя
	Project   string   // .HOST без www
	Name      string   // имя хоста без зоны и www
	Slugs     []string // слаги домена
	DashCount int      // флаг/кол-во дефисов в .Host
	TLD       string   // зона
}

type Field struct {
	UrlReal string // начальный урл
	UrlGood string // полный правильный адрес для запроса
	UrlView string // приглядный адрес для людей

	Real string // реальный полученный хост
	Puny Host   // детали представления имени домена
	IDN  Host   //

	Path      string   // путь до "?"
	Dir       string   // каталог
	Slugs     []string // слаги полного пути, без первого пустого (разделитель - "/")
	File      string   // файл
	FileSlugs []string // слаги файла (разделитель - ".")
	FileName  string   // имя файла без расширения
	FileExt   string   // расширение файла

	QueryString string              // QUERY_STRING
	QueryList   []map[string]string // разобранный QUERY_STRING
	QueryAmp    bool                // QUERY_STRING has &amp;

	Scheme   string // http || https
	Port     int    // порт для соединения
	Username string // Пользователь для Basic-аутентификации
	Password string // Пароль для Basic-аутентификации
	Frag     string // Часть ссылки после хеша

	ManualPort int    // manual port
	IsSecure   bool   // секьюрность
	HasDouble  bool   // флаг наличия дубля слеша в пути
	HasDots    bool   // флаг наличия точек в начале и конце имени
	HasWWW     bool   // флаг наличия www
	Region     string // региональность (заодно и значение, соответствие городу, области и т.п.)
	Service    string // сервисность/государственность (заодно и значение) ex: livejournal.com narod.ru ucoz.ru
}

/*<?php

// --------------------------------------------------------------------------
// Разбор и сбор урла
// --------------------------------------------------------------------------

class srfURL
{



	protected $struct = array(
		'url_real'   => '',      // начальный урл
		'url_good'   => '',      // полный правильный адрес для запроса
		'url_view'   => '',      // приглядный адрес для запроса

		'real'       => '',      // реальный полученный хост

		'puny'       => '',      // полный хост в punycode (без www)
		'puny_slugs' => array(), // puny-домен по слагам
		'f_punydash' => 0,       // флаг/кол-во дефисов в puny-домене
		'puny_name'  => '',      // имя хоста без зоны и www
		'puny_tld'   => '',      // зона

		'idn'        => '',      // полный хост в idn (без www)
		'idn_slugs'  => array(), // idn-домен по слагам
		'f_idndash'  => 0,       // флаг/кол-во дефисов в idn-домене
		'idn_name'   => '',      // имя хоста без зоны и www
		'idn_tld'    => '',      // зона

		'path'       => '/',     // путь до ?
		'dir'        => '/',     // каталог
		'slugs'      => array(), // слаги полного пути, без первого пустого (разделитель - "/")
		'file'       => '',      // файл
		'file_slugs' => array(), // слаги файла (разделитель - ".")
		'file_name'  => '',      // имя файла без расширения
		'file_ext'   => '',      // расширение файла

		'qs'         => '',      // QUERY_STRING
		'qs_arr'     => array(), // разобранный QUERY_STRING
		'f_amp'      => 0,       // QUERY_STRING has &amp;

		'proto'      => '',      // http || https
		'port'       => 0,       // порт для соединения
		'user'       => '',      // Пользователь для Basic-аутентификации
		'pass'       => '',      // Пароль для Basic-аутентификации
		'frag'       => '',      // Часть ссылки после хеша

		'f_port'     => 0,       // manual port
		'f_secure'   => 0,       // секьюрность
		'f_double'   => 0,       // флаг наличия дубля слеша в пути
		'f_dot'      => 0,       // флаг наличия точек в начале и конце имени
		'f_www'      => 0,       // флаг наличия www
		'f_region'   => '',      // региональность (заодно и значение, соответствие городу, области и т.п.)
		'f_service'  => '',      // сервисность/государственность (заодно и значение) ex: livejournal.com narod.ru ucoz.ru
	);



	protected $item;
	protected $idna;
	protected $name;

	// Конструктор
	public function __construct($url) {
		$this->idna = new idna_convert();
		$this->item = new srfStrictStruct($this->struct, 'srfURL', DIR_LOG . 'url.' . strftime('%Y-%m-%d') . '.txt');
		$this->name = $url;
		$this->_parse_url($url);
	}

	// Дубликатор
	public function __clone() {
		$this->item = clone $this->item;
		// idna - static, nobody care
	}

	// Получение имени урла (или полного урла)
	public function getName() {
		return $this->name;
	}

	// Установка имени для урла (для идентификации в групповых операциях)
	public function SetName($name) {
		$this->name = $name;
		return $this;
	}

	// ----------------------------------------------------------------------
	// Сборка простого урла для запроса
	public function buildSimple() {
		list($proto, $f_www, $host, $dir, $file, $qs) = $this->item->Chain()
			->getFieldList('proto', 'f_www', 'puny', 'dir', 'file', 'qs');

		return ''
			. $proto
			. '://'
			. ($f_www ? 'www.' : '')
			. $host
			. ($dir ? $dir:'/')
			. $file
			. ($qs ? '?' . $qs : '')
		;
	}

	// ----------------------------------------------------------------------
	// Установка значения поля
	public function Set($field, $value) {
		$this->item->SetField($field, $value);
		return $this;
	}

	// ----------------------------------------------------------------------
	// Проверка наличия ошибок
	public function hasErrors() {
		return $this->item->Check()->hasErrors();
	}

	// ----------------------------------------------------------------------
	// Получение списка ошибок
	public function getErrors() {
		return $this->item->Check()->getErrors();
	}

	// ----------------------------------------------------------------------
	// Получение полной структуры
	public function getStruct() {
		return $this->item->getItem();
	}

	// ----------------------------------------------------------------------
	// Получение полной структуры
	public function getField($name) {
		return $this->item->getField($name);
	}

	// ----------------------------------------------------------------------
	// Разбор урла на поля
	protected function _parse_url($url) {

		if(strncasecmp($url, 'http://', 7)
			&& strncasecmp($url, 'https://', 8)
		) {
			$url = 'http://' . $url;
		}
		$this->item->SetField('url_real', $url);

		// Приведение схемы к нижнему регистру
		$pos = mb_strpos($url.'////////', '/', 8, 'UTF-8');
		if($pos === false) {
			$url = mb_strtolower($url, 'UTF-8') . '/';
		} else {
			$url = mb_strtolower(mb_substr($url, 0, $pos, 'UTF-8')) . mb_substr($url, $pos, 100000, 'UTF-8');
		}

		// Удаление &amp; из урла
		$this->item->SetDefault('f_amp');
		if(false !== mb_strpos($url, '&amp;', 0, 'UTF-8')) {
			$url = str_replace('&amp;', '&', $url);
			$this->item->SetField('f_amp', 1);
		}

		// Разбор стандартных полей
		$parsed = @parse_url($url);
		if($parsed === false) {
			return array(array(), 'parse_url() fails');
		}
		if(empty($parsed['path'])) {
			$parsed['path'] = '/';
		}

		// Логин и пароль
		$this->item->Chain()
			->SetField('user', (string) @$parsed['user'])
			->SetField('pass', (string) @$parsed['pass'])
		;

		// Разбор схемы и портов
		$this->item->SetDefault('f_port');
		// http
		if($parsed['scheme'] == 'http') {
			$this->item->Chain()
				->SetField('f_secure', 0)
				->SetField('proto', 'http')
				->SetField('port', 80)
			;
			if(isset($parsed['port']) && $parsed['port'] != 80) {
				$this->item->Chain()
					->SetField('port', $parsed['port'])
					->SetField('f_port', 1)
				;
			}
		}
		// https
		elseif($parsed['scheme'] == 'https') {
			$this->item->Chain()
				->SetField('f_secure', 1)
				->SetField('proto', 'https')
				->SetField('port', 443)
			;
			if(isset($parsed['port']) && $parsed['port'] !== 443) {
				$this->item->Chain()
					->SetField('port', $parsed['port'])
					->SetField('f_port', 1)
				;
			}
		}

		// ----------------------------------------------------------------------
		// Домен и поддомен
		// ----------------------------------------------------------------------
		$host = mb_strtolower($parsed['host']);

		// Точки на концах имени
		$this->item->Chain()
			->SetField('real', $parsed['host'])
			->SetDefault('f_dot')
		;
		if('.' == $host[0]) {
			$host = substr($host, 1);
			$this->item->Inc('f_dot');
		}
		if('.' == $host[strlen($host)-1]) {
			$host = substr($host, 0, -1);
			$this->item->Inc('f_dot');
		}

		// Проверка наличия префикса www.
		$this->item->SetDefault('f_www');
		while(substr($host, 0, 4) == 'www.') {
			$host = mb_substr($host, 4, 10000, 'UTF-8');
			$this->item->Inc('f_www');
		}

		// puny-представление
		$puny = $this->idna->encode($host);

		// idn-представление
		$idn = $this->idna->decode($puny);

		foreach(array('puny'=>$puny, 'idn'=>$idn) as $type => $str) {
			$slg = $arr = explode('.', $str);
			$cnt = substr_count($str, '-');
			$tld = array_pop($arr);
			$nam = implode('.', $arr);
			$this->item->Chain()
				->SetField($type, $str)
				->SetField($type.'_slugs', $slg)
				->SetField('f_'.$type.'dash', $cnt)
				->SetField($type.'_tld', $tld)
				->SetField($type.'_name', $nam)
			;
		}

		// Дубль слеша в пути
		$this->item->Chain()
			->SetDefault('f_double')
			->IfTrue(false !== strpos($parsed['path'], '//'))
			->SetField('f_double', 1)
		;
		$path  = str_replace('//', '/', $parsed['path']);
		$dir   = (string) substr($path, 0, strrpos($path, '/') + 1);
		$file  = (string) substr($path, strlen($dir));

		// Подготовка списка слагов
		$arr = array_slice(explode('/', $path), 1);
		if(!$arr[count($arr)-1]) {
			$arr = array_slice($arr, 0, -1);
		}
		$slugs = $arr;

		// Путь, каталог, файл и их слаги
		$this->item->Chain()
			->SetField('slugs', $slugs)
			->SetField('path', $path)
			->SetField('dir', $dir)
			->SetField('file', $file)
			->SetField('file_slugs', $file ? explode('.', $file) : array())
			->SetField('file_name', $file)
			->SetDefault('file_ext')
		;
		if(false !== ($ex = strrpos($file, '.'))) {
			$this->item->Chain()
				->SetField('file_name', substr($file, 0, $ex))
				->SetField('file_ext', substr($file, $ex + 1))
			;
		}

		// Get-параметры
		$this->item->Chain()
			->SetDefault('qs')
			->SetDefault('qs_arr')
		;
		if(isset($parsed['query'])) {
			mb_parse_str($parsed['query'], $arr);
			$this->item->Chain()
				->SetField('qs', $parsed['query'])
				->SetField('qs_arr', $arr)
			;
		}

		// Якорь
		$this->item->SetField('frag', (string) @$parsed['fragment']);

		// Представления адреса
		$idn = array('http://');
		$puny = array('http://');
		if($this->item->getField('f_www')) {
			$idn[] = $puny[] = str_repeat('www.', $this->item->getField('f_www'));
		}
		$idn[] = $this->item->getField('idn');
		$puny[] = $this->item->getField('puny');
		$idn[] = $puny[] = $this->item->getField('path');
		if($this->item->getField('qs')) {
			$idn[] = $puny[] = '?' . $this->item->getField('qs');
		}
		if($this->item->getField('frag')) {
			$idn[] = $puny[] = '#' . $this->item->getField('frag');
		}
		$this->item->Chain()
			->SetField('url_view', implode('', $idn))
			->SetField('url_good', implode('', $puny))
		;
		// TODO
		$this->item->Chain()
			->TODO('f_region')
			->TODO('f_service')
		;
	}
}

?>*/
