package plan

// package url

// 50+ parameters from php-api

// package doc
//
// Отчёт по ошибкам (контекст по ПС)

type TagRule struct {
	Name string
	Pair bool
}

type ParseRule struct {
	IsTag        bool     // Флаг для обработки текста по тегам
	TagRule      *TagRule // Список правил для каждого тега
	IsContextTag bool     // Контекстный тег, i.e. <td>
	Strings      string   //
	RegExp       string   //
}

type ParsePlan struct {
	Rules map[string]ParseRule
}

type Token struct {
}

func (t *Token) Parse() {
}

type NLP struct {
	bytes     []byte
	runes     []rune
	words     []word
	sentences []sentence
	passages  []passage
	paragraph  []passage
	blocks    []paragraph
	texts     []text
}


func TagSplit() {

}

func PassageSplit() {
	html.Scanner().Tokens(tok.ByTags)
}

type Client struct {
	NewStorage(Storage)
	NewFetcher(Fetcher)
	NewProxy(Proxy)
	TestWwwWide()
	Param().Conc(N) // Live Time, Keep Alive
	Delay() //
	Cookie(Storage)
	WebDriver(Storage)
	Timeout() // seacrch, resolve, connect, receive, keep-alive
	TimeLimit() // <- time.Limit(10*time.Sec)
	RequestNumLimit();
}

func ClientSyncAsync() {
	s := new Sere
	s.Chain() 
	b := sere.ProcessControl(); // stop, replace, restore, restart
}