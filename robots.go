package robots

/*

$t = '

Allow: /
# test robots
user-agent: *# common rule
Disallow: *?

User-Agent: yandex
Disallow: * /yandex.html

Disallow: /* /yandex.php
Host: www.some.me

User-agent: googlebot0
User-agent: googlebot1
User-agent: googlebot2
User-agent: googlebot3

User-agent: googlebot4
User-agent: googlebot5
User-agent: google
Disallow: /* /google.php
Disallow: /* /googlesdvsdvsd.php
Disallow: /* /googlesdvsdvsd.php121


Crawl-delay: 0.2
';

$l = 'Disallow:';

$robots = new srfRobots($t, $l);

srf::r($robots->Get());

*/

/*
<?php

// --------------------------------------------------------------------------
// Разбор и обработка robots.txt
// --------------------------------------------------------------------------

class srfRobots
{
	protected $arr;
	protected $par;

	// ----------------------------------------------------------------------
	# Инициализация списком параметров
	// $robots - содержание файла
	// $local  - имя блока для обработки
	// ----------------------------------------------------------------------
	public function __construct($robots = null, $local = null) {

		$this->arr = array(
			'host' => '',
			'add'  => '',
			'src'  => '',
			'data' => array(),
		);

		$this->Set($robots, $local);
		TODO(__CLASS__, 'должен быть абстрактным');

		TODO(__METHOD__, 'не забывать писать алармы для неразрешённых полей и значений');
		TODO(__METHOD__, 'очень внимательно обрабатывать случаи alarm_same_len(какой инструкции отдаётся приоритет');
		TODO(__METHOD__, 'разрешение использовать межсекционные инструкции (ext_allow = 1?)');

		$this->par = array (

			// Максимальные размеры файла в байтах
			'max_size' => 20000, // в байтах
			'max_str'  => 20000, // в строках

			// Игнорировать allow/disallow без аргумента (иначе отрабатывать
			// по правилам, например, яндекс запрещает всё при "Allow: ")
			'arg_void_ignore' => 1,

			// Игнорировать аргументы disallow, содержащие только /
			'disallow_ignore_root' => 1,

			// Гугл разрешает использовать несколько юзер-агентов подряд для одного блока правил
			'useragent_several' => 0,

			// Создавать предупреждение, если длина правил d/a одинаковая
			'alarm_same_len' => 1, // яндекс сделает allow

			// Игнорировать пустые инструкции внутри блоков
			'empty_rule_ignore' => 1,

			// Требование пустой строки перед user-agent блока
			'empty_line_before' => 1,

			// Разрешить первый не_слеш в значениях d/a
			'first_not_slash' => 0,

			//  TODO
			'ext_allow' => 0, // 0 = все инструкции должны идти внзу блока d/a

			// Откуда брать {Host:}, если он поддерживается
			'ext_host'           => 0, // разрешить использовать ${}
			'ext_host_outer'     => 1, // можно брать первое определение (вне блоков, яндекс) иначе из первого блока
			'ext_host_block'     => 1, // брать ${} из первого блока
			'ext_host_block_my'  => 1, // брать ${} только из своего блока TODO уточнить

			// Откуда брать {Crawl-delay:}, если он поддерживается
			'ext_crawl'          => 0, // разрешить использовать ${}
			'ext_crawl_outer'    => 1, // можно брать первое определение (вне блоков, яндекс) иначе из первого блока
			'ext_crawl_block'    => 1, // брать ${} из первого блока
			'ext_crawl_block_my' => 1, // брать ${} только из своего блока TODO уточнить

			// Откуда брать {Clean-param:}, если он поддерживается
			'ext_clean'          => 0, // разрешить использовать ${}
			'ext_clean_outer'    => 1, // можно брать первое определение (вне блоков, яндекс) иначе из первого блока
			'ext_clean_block'    => 1, // брать ${} из первого блока
			'ext_clean_block_my' => 1, // брать ${} только из своего блока TODO уточнить
		);


		// игнорируемые параметры надо как-то дополнительно вверх передавать
		// и учитывать первое значение параметра, если страница не найдена
		// без него

		// статус файла интерпретируется по разному у гугла и яндекса?

		// is_200
		// is_txt
		// допустимы ли начальные пробелы у строки?
		// twice ua at file

		TODO(__METHOD__, 'Настройка: добавлять слеш в инструкции с начальной звёздочкой');
		// добавлять слеш в инструкции с начальной звёздочкой ? TODO - нет,
		// но нужно определить кто и почему их поддерживает

		// Нужна дополнительная проверка по нескольким юзер-агентам со
		// своими правилами, типа доступность страницы для яндекса И гугла
		// и своих правил

		// нужен системный список игнорируемых страниц типа флешей, пдф,
		// видео, док, эксель и т.п.
	}



	// ----------------------------------------------------------------------
	// Установка сайтового роботса и локального
	// ----------------------------------------------------------------------
	public function Set($robots = null, $local = null)
	{
		$this->arr = array(
			'host' => $robots === null ? ($this->arr['host'] ? '' : $this->arr['host']) : $robots,
			'add'  => $local  === null ? ($this->arr['add']  ? '' : $this->arr['add'])  : $local,
		);

		$this->arr['src'] = "\n"
			. 'User-agent: ' . $this->wrapper
			. "\n\n"
			. trim($this->arr['host'])
			. (
				empty($this->arr['add'])
				? ''
				: "\n\n"
					. "User-agent: " . $this->local
					. "\n"
					. trim($this->arr['add'])
			)
			. "\n";

		$this->Prepare();

		return $this;
	}

	// ----------------------------------------------------------------------
	// Получить текущие состояния
	// ----------------------------------------------------------------------
	public function Get() {
		return $this->arr;
	}

	// ----------------------------------------------------------------------
	// Подготовка инструкций к работе
	// ----------------------------------------------------------------------
	public function Prepare() {
		TODO(__METHOD, 'надо добавить ручной роботс со своим юзер-агентом, делеем b n/g/');

		$arr = explode("\n", $this->arr['src']);
		$res = array();
		$n   = count($arr);

		$ua    = '';
		$name  = '';
		$value = '';

		$f_void = 0;
		$f_ua   = 0;
		$f_next = 0;

		for($i=0; $i<$n; $i++) {
			$s = $arr[$i];
			$p = trim(srf::substr_before($s, '#')); // убираем комменты

			$name  = '';
			$value = '';

			if(mb_strpos($p, ':')) {
				$name  = trim(mb_strtolower(srf::substr_before($p, ':')));
				$value = trim(srf::substr_after($p, ':'));

				switch($name) {
					// ------------------------------------------------------
					case 'user-agent':
						$pr = strtolower($value); // current, $ua - previous

						if(empty($res['ua'][$pr])) {
							$res['ua'][$pr] = array(
								'add' => array(
									'nl'    => '',
									'next'  => '',
									'base'  => '',
									'twice' => '',
									'void'  => array()
								),
								'src'  => array(),
								'rule' => array(),
								'ext'  => array(),
							);
						} else {
							$res['ua'][$pr]['add']['twice'] = 1;
						}

						if($f_ua) {
							$res['ua'][$ua]['add']['next'] = $pr;
							$f_next = 1;
						}

						if($f_void) {
							$res['ua'][$pr]['add']['nl'] = 1;
						}
						$ua   = $pr;
						$f_ua = 1;
						if(empty($res['ua'][$pr])) {
							$res['ua'][$pr] = array();
						}
						break;
					// ------------------------------------------------------
					case 'allow':
					case 'disallow':
						// Trick for rsort 'l' > 'i' on same len
						$res['ua'][$ua]['src'][]  = $s;
						$res['ua'][$ua]['rule'][] = mb_substr($name, 1, 1) . ($value ? ' ' . $value : '');

						if($f_void) {
							$res['ua'][$ua]['add']['void'][] = $i - 1;
						}
						$f_ua = 0;
						break;
					// ------------------------------------------------------
					case 'clean-param':
					case 'crawl-delay':
					default:
						$res['ua'][$ua]['ext'][$name][] = $value;
						$res['ext'][$name][] = $value;
						$f_ua = 0;
						break;
					// ------------------------------------------------------
				}

				$f_void = 0;
			}
			// --------------------------------------------------------------
			elseif($p) {
				$this->Error($this->_me, 'BAD_FORMAT_AT_LINE(' . ($i+1) . ')');
				$f_void = 0;
				$f_ua   = 0;
			}
			// --------------------------------------------------------------
			else {
				$f_void = 1;
				$f_ua   = 0;
			}
			// --------------------------------------------------------------
		}

		// Зависимости для блочных юзер-агентов

		$i = 100; //inf_break
		while($f_next) {
			$f_next = 1;
			foreach($res['ua'] as $k=>$v)
			{
				$f = 0;
				if($v['add']['next'] && !$v['rule']) {
					if($res['ua'][ $v['add']['next'] ]['src']) {
						$res['ua'][$k]['src'] = $res['ua'][ $v['add']['next'] ]['src'];
					}

					if($res['ua'][ $v['add']['next'] ]['rule']) {
						$res['ua'][$k]['rule'] = $res['ua'][ $v['add']['next'] ]['rule'];
					}

					if($res['ua'][ $v['add']['next'] ]['ext']) {
						$res['ua'][$k]['ext']  = $res['ua'][ $v['add']['next'] ]['ext'];
					}

					// TODO : разобраться тут с логикой и f_next, всё же поменялось...
					if(1) {
						$res['ua'][$k]['add']['base'] = $v['add']['next'];
						if($res['ua'][ $v['add']['next'] ]['add']['next'] && $res['ua'][ $v['add']['next'] ]['add']['base']) {
							$res['ua'][ $v['add']['next'] ]['add']['base'];
						}
					}
					$f_next = 1;
				}
			}

			//inf_break
			if($i-- < 0) {
				break;
			}
		}

		srf::tab_print($res['ua'], 1);

		$res = srf::significant($res);

		$this->arr['data'] = $res;

		return $this;
	}

	// ----------------------------------------------------------------------
	// Сортировка списка инструкций
	// ----------------------------------------------------------------------
	public function Sort() {
	}

	// ----------------------------------------------------------------------
	// Проверка адреса на пригодность
	// ----------------------------------------------------------------------
	public function isAllow() {
	}

	// ----------------------------------------------------------------------
	public function ArrFilter() {
		// [ $this->isAllow() ]
	}

	// ----------------------------------------------------------------------
	public function UaExists($ua) {
		$arr = array_flip( $this->AgentList() );
		return isset($arr[$ua]);
	}

	// ----------------------------------------------------------------------
	public function UaList() {
		$arr = $this->Prepare();
		return array_keys($arr['ua']);
	}

	// ----------------------------------------------------------------------
	public function UaSelect($ua = '*') {
		$this->ua_best = $ua;
		return $this;
	}

	// ----------------------------------------------------------------------
	public function MultiMode($mode = self::MODE_AND) {
		$this->mode = $mode;
		return $this;
	}

	// ----------------------------------------------------------------------
	public function CacheSize($size = self::CACHE_NONE) {

	}
	// ----------------------------------------------------------------------
}




#############################################################################
#############################################################################
#############################################################################

/*

Как удобней обрабатывать список урлов?
Оставлять только годные?
Куда девать отфильтрованные?
Кешировать ли? Размер кеша? Может просто 2 функции?
Использовать ограничения на размеры с нотисом?
Нормализовать ли урл из роботса? _&amp;_param=232

Case-sensitive для инструкций и значений не в урл
User-agent: *
Crawl-delay:
SITEMAP:

UA: *
encoding
empty-line-before-ua
empty-line-ignore
maxsizekb
maxsizestr
field-case-sensetive
value-case-sensetive
same-len-defined
fail-fetch-eq-allow
no-me-or-wild-eq-allow
void-allow-eq-disallow
host-rulez


В тесты:
	в начале строк
	наличие пустых строк в блоках
	интерпретация *$ на конце (яндекс игнорирует$)


Можно ли использовать несколько юзер-агентов через запятую?


А что делать с национальными кодировками?? + разными роботсами ???


Проверить возможность пробела перед двоеточием


-----------------------------------------------------------------------------
MAIL.RU
-----------------------------------------------------------------------------
http://help.mail.ru/webmaster/indexing/robots.txt/rules/format

User-agent: Mail.ru

Если HTTP-код ответа сервера на URL не '200', либо формат полученных данных нарушен, а также в случае превышения файлом допустимого размера в 100кб считается, что робот не имеет ограничений для данного сайта.

_НАИМЕНОВАНИЕ_ ПОЛЯ ЯВЛЯЕТСЯ РЕГИСТРОНЕЗАВИСИМЫМ.

Для каждой записи обязательна строка с директивой Disallow.


USER-AGENT = Mail.ru, *


ALLOW/DISALLOW

same_len_dis_al = N/A


HOST
Внутри записи, начинающейся с user-agent.


CRAWL-DELAY
Дробные числа, внутри блока.



-----------------------------------------------------------------------------
GOOGLE
-----------------------------------------------------------------------------
https://developers.google.com/webmasters/control-crawl-index/docs/robots_txt
http://support.google.com/webmasters/bin/answer.py?hl=en&answer=156449

UTF-8
!UTF-8 = возможны ошибки при разборе

An optional Unicode BOM (byte order mark) at the beginning of the robots.txt file is ignored.
----->> Надо у себя проверить что за хрень

Spaces are optional (but recommended to improve readability).
Comments "#"

Несколько UA могут быть указаны один за другим, типа ок:
User-agent: first
User-agent: second
Disallow: /someurl

Googlebot will ignore white-space (in particular empty lines)and unknown directives in the robots.txt.


URL                           allow:   disallow:   Verdict
http://example.com/page.htm   /page    /*.htm      N/A


The <field> element is case-insensitive. The <value> element may be case-sensitive, depending on the <field> element.

Google currently enforces a size limit of 500kb.
Неправильный ответ или содержание не будут восприниматься роботов всерьёз (уточнить по видео из архива).

Файл привязывается к хосту, порту и протоколу.

User-agent:
	Googlebot, *

Else:
	Googlebot-Mobile, Googlebot, *
	Googlebot-Image
	Googlebot-News
	Googlebot-Video
	AdsBot-Google
	Mediapartners-Google


The entry should begin with a forward slash (/).



-----------------------------------------------------------------------------
Yandex
-----------------------------------------------------------------------------
Сессия начинается с закачки robots.txt сайта, если его нет, он не текстовый или на запрос робота возвращается HTTP-код отличный от '200', считается, что доступ роботу не ограничен.

Важно. В случае возникновения конфликта между двумя директивами с префиксами одинаковой длины приоритет отдается директиве Allow.

Если записи 'User-agent: Yandex' и 'User-agent: *' отсутствуют, считается, что доступ роботу не ограничен.

User-agent: YandexBot, Yandex, *

Yandex - общий для всех роботов яндекса
YandexBot — основной индексирующий робот;
YandexMedia — робот, индексирующий мультимедийные данные;
YandexImages — индексатор Яндекс.Картинок;
YandexCatalog — «простукивалка» Яндекс.Каталога, используется для временного снятие с публикации недоступных сайтов в Каталоге;
YandexDirect — робот Яндекс.Директа, особым образом интерпретирует robots.txt;  (игнор остальных)
YandexBlogs — робот поиска по блогам, индексирующий посты и комментарии;
YandexNews — робот Яндекс.Новостей;
YandexPagechecker — валидатор микроразметки;
YandexMetrika — робот Яндекс.Метрики;
YandexMarket— робот Яндекс.Маркета;
YandexCalendar — робот Яндекс.Календаря (игнор остальных)


Недопустимо наличие пустых переводов строки между директивами 'User-agent' и 'Disallow' ('Allow'), а также между самими 'Disallow' ('Allow') директивами. (в тесты)

Кроме того, в соответствии со стандартом перед каждой директивой 'User-agent' рекомендуется вставлять пустой перевод строки. (но расширения яндекса могут понять что и как)



User-agent: Yandex
Disallow: # то же, что и Allow: /

User-agent: Yandex
Allow: # то же, что и Disallow: /



HOST
добавлять в группе, начинающейся с записи 'User-Agent', непосредственно после директив 'Disallow'('Allow').

Аргументом директивы 'Host' является доменное имя с номером порта (80 по умолчанию), отделенным двоеточием.

Тем не менее директива Host является межсекционной, поэтому будет использоваться роботом вне зависимости от места в файле robots.txt, где она указана.

Важно: Директива Host в файле robots.txt может быть только одна. В случае указания нескольких директив, использоваться будет первая.

Важно: параметр директивы Host обязан состоять из одного корректного имени хоста (то есть соответствующего RFC 952 и не являющегося IP-адресом) и допустимого номера порта. Некорректно составленные строчки 'Host:' игнорируются.


CLEAN-PARAM

Clean-param: ref /some_dir/get_book.pl

Важно: директива Clean-Param является межсекционной, поэтому может быть указана в любом месте файла robots.txt. В случае, если директив указано несколько, все они будут учтены роботом.

Действует ограничение на длину правила — 500 символов.

Clean-param: s /forum/showthread.php
Clean-param: sid /index.php

Clean-param: s&ref /forum* /showthread.php

Clean-param: s /forum/index.php
Clean-param: s /forum/showthread.php


Слишком большие robots.txt (более 32 Кб) считаются полностью разрешающими, то есть рассматриваются аналогично:

Crawl-delay и Host должны добавляться после disallow и allow в группе User-agent

?>
*/
