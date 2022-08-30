package clean

import (
	"github.com/aaaton/golem/v4"
	"github.com/aaaton/golem/v4/dicts/de"
	"github.com/aaaton/golem/v4/dicts/en"
	"github.com/aaaton/golem/v4/dicts/es"
	"github.com/aaaton/golem/v4/dicts/fr"
	"github.com/aaaton/golem/v4/dicts/it"
	"github.com/aaaton/golem/v4/dicts/sv"
	"github.com/mx79/go-nlp/utils"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

// Purger
type Purger struct {
	Language       string
	LemmInit       bool
	Lemmatizer     *golem.Lemmatizer
	Stopwords      []string
	NoPunct        bool
	Lemmatize      bool
	Lowercase      bool
	NoAccent       bool
	NoStopword     bool
	NoLowTfidfWord bool
}

// NewPurger
func NewPurger(lang string, initLemm bool, noPunct bool, lemmatize bool, lowercase bool, noAccent bool, noStopword bool, noLowTfidfWord bool) *Purger {
	var (
		langPack   golem.LanguagePack
		lemmatizer *golem.Lemmatizer
	)
	if initLemm {
		if lang != "fr" && lang != "en" && lang != "es" && lang != "it" && lang != "de" && lang != "sv" {
			panic("lang variable value is not recognized, use one of the followings: " +
				"'fr', 'en', 'es, 'it', 'de', 'sv'")
		} else if lang == "fr" {
			langPack = fr.New()
		} else if lang == "en" {
			langPack = en.New()
		} else if lang == "es" {
			langPack = es.New()
		} else if lang == "it" {
			langPack = it.New()
		} else if lang == "de" {
			langPack = de.New()
		} else if lang == "sv" {
			langPack = sv.New()
		}
		lemmatizer, _ = golem.New(langPack)
	}
	return &Purger{
		Language:       lang,
		LemmInit:       initLemm,
		Lemmatizer:     lemmatizer,
		Stopwords:      stopwordList(lang),
		NoPunct:        noPunct,
		Lemmatize:      lemmatize,
		Lowercase:      lowercase,
		NoAccent:       noAccent,
		NoStopword:     noStopword,
		NoLowTfidfWord: noLowTfidfWord,
	}
}

// PurgeText The function that allows to clean a given text in depth by applying several layers of treatment
// return: The sentence or word list based on boolean values
func (p *Purger) PurgeText(s string) string {
	if p.NoPunct {
		s = removePunctuation(s)
	}
	if p.Lemmatize {
		s = p.lemmatize(s)
	}
	if p.Lowercase {
		s = lower(s)
	}
	if p.NoAccent {
		s = removeAccent(s)
	}
	if p.NoStopword {
		s = stopword(s, p.Stopwords)
	}
	//if noLowTfidfWord {
	//	s = LowTfidfWord(s)
	//}
	return s
}

// slash trim the left side of a text containing ', useful to lem "j'ai" -> "ai" = "avoir"
func slash(word string) string {
	if strings.Contains(word, "'") {
		return strings.Split(word, "'")[1]
	}
	return word
}

// lemmatize The function that lemmatizes the words of a sentence or a word
func (p *Purger) lemmatize(s string) string {
	if p.LemmInit {
		var sent string
		for _, word := range Tokenize(s) {
			sent += p.Lemmatizer.Lemma(slash(word)) + " "
		}
		return sent
	} else {
		panic("Purger cannot use func lemmatize if attribute LemmInit is not set to true")
	}
}

// removePunctuation The function that allows to remove punctuation in a sentence
func removePunctuation(s string) string {
	punctuation := "!@#$%^&*()[]_+<>?:.,;"
	for _, c := range punctuation {
		s = strings.Replace(s, string(c), "", -1)
	}
	return s
}

// removeAccent The function that allows you to remove the accents in a sentence
func removeAccent(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, err := transform.String(t, s)
	if err != nil {
		panic(err)
	}
	return output
}

// stopwordList The function that retrieves a list of stopwords
func stopwordList(lang string) []string {
	var stopwords []string
	if lang != "fr" && lang != "en" && lang != "es" && lang != "it" && lang != "de" && lang != "sv" {
		panic("lang variable value is not recognized, use one of the followings: " +
			"'fr', 'en', 'es, 'it', 'de', 'sv'")
	} else if lang == "fr" {
		stopwords = []string{"alors", "au", "aucuns", "aussi", "autre", "avant", "avec", "avoir", "bon", "ce", "cela", "ces", "ceux", "chaque", "ci", "comme", "comment", "dans", "des", "du", "dedans", "dehors", "depuis", "devrait", "doit", "dos", "début", "elle", "elles", "en", "encore", "essai", "est", "eu", "fait", "faites", "fois", "font", "hors", "ici", "il", "ils", "je", "la", "le", "les", "leur", "là", "ma", "maintenant", "mes", "mien", "moins", "mon", "mot", "meme", "nommer", "notre", "nous", "ou", "par", "parce", "pas", "peut", "peu", "plupart", "pour", "pourquoi", "quand", "que", "quel", "quelle", "quelles", "quels", "qui", "sa", "sans", "ses", "seulement", "si", "sien", "son", "sont", "sous", "soyez", "sur", "ta", "tandis", "tellement", "tels", "tes", "ton", "tous", "tout", "trop", "très", "tu", "voient", "vont", "votre", "vous", "vu", "ça", "état", "être", "euh", "coup"}
	} else if lang == "en" {
		stopwords = []string{"a", "about", "above", "after", "again", "against", "all", "am", "an", "and", "any", "are", "aren't", "as", "at", "be", "because", "been", "before", "being", "below", "between", "both", "but", "by", "can't", "cannot", "could", "couldn't", "did", "didn't", "do", "does", "doesn't", "doing", "don't", "down", "during", "each", "few", "for", "from", "further", "had", "hadn't", "has", "hasn't", "have", "haven't", "having", "he", "he'd", "he'll", "he's", "her", "here", "here's", "hers", "herself", "him", "himself", "his", "how", "how's", "i", "i'd", "i'll", "i'm", "i've", "if", "in", "into", "is", "isn't", "it", "it's", "its", "itself", "let's", "me", "more", "most", "mustn't", "my", "myself", "no", "nor", "not", "of", "off", "on", "once", "only", "or", "other", "ought", "our", "ours \tourselves", "out", "over", "own", "same", "shan't", "she", "she'd", "she'll", "she's", "should", "shouldn't", "so", "some", "such", "than", "that", "that's", "the", "their", "theirs", "them", "themselves", "then", "there", "there's", "these", "they", "they'd", "they'll", "they're", "they've", "this", "those", "through", "to", "too", "under", "until", "up", "very", "was", "wasn't", "we", "we'd", "we'll", "we're", "we've", "were", "weren't", "what", "what's", "when", "when's", "where", "where's", "which", "while", "who", "who's", "whom", "why", "why's", "with", "won't", "would", "wouldn't", "you", "you'd", "you'll", "you're", "you've", "your", "yours", "yourself", "yourselves"}
	} else if lang == "es" {
		stopwords = []string{"un", "una", "unas", "unos", "uno", "sobre", "todo", "también", "tras", "otro", "algún", "alguno", "alguna", "algunos", "algunas", "ser", "es", "soy", "eres", "somos", "sois", "estoy", "esta", "estamos", "estais", "estan", "como", "en", "para", "atras", "porque", "por qué", "estado", "estaba", "ante", "antes", "siendo", "ambos", "pero", "por", "poder", "puede", "puedo", "podemos", "podeis", "pueden", "fui", "fue", "fuimos", "fueron", "hacer", "hago", "hace", "hacemos", "haceis", "hacen", "cada", "fin", "incluso", "primero \tdesde", "conseguir", "consigo", "consigue", "consigues", "conseguimos", "consiguen", "ir", "voy", "va", "vamos", "vais", "van", "vaya", "gueno", "ha", "tener", "tengo", "tiene", "tenemos", "teneis", "tienen", "el", "la", "lo", "las", "los", "su", "aqui", "mio", "tuyo", "ellos", "ellas", "nos", "nosotros", "vosotros", "vosotras", "si", "dentro", "solo", "solamente", "saber", "sabes", "sabe", "sabemos", "sabeis", "saben", "ultimo", "largo", "bastante", "haces", "muchos", "aquellos", "aquellas", "sus", "entonces", "tiempo", "verdad", "verdadero", "verdadera \tcierto", "ciertos", "cierta", "ciertas", "intentar", "intento", "intenta", "intentas", "intentamos", "intentais", "intentan", "dos", "bajo", "arriba", "encima", "usar", "uso", "usas", "usa", "usamos", "usais", "usan", "emplear", "empleo", "empleas", "emplean", "ampleamos", "empleais", "valor", "muy", "era", "eras", "eramos", "eran", "modo", "bien", "cual", "cuando", "donde", "mientras", "quien", "con", "entre", "sin", "trabajo", "trabajar", "trabajas", "trabaja", "trabajamos", "trabajais", "trabajan", "podria", "podrias", "podriamos", "podrian", "podriais", "yo", "aquel"}
	} else if lang == "it" {
		stopwords = []string{"a", "adesso", "ai", "al", "alla", "allo", "allora", "altre", "altri", "altro", "anche", "ancora", "avere", "aveva", "avevano", "ben", "buono", "che", "chi", "cinque", "comprare", "con", "consecutivi", "consecutivo", "cosa", "cui", "da", "del", "della", "dello", "dentro", "deve", "devo", "di", "doppio", "due", "e", "ecco", "fare", "fine", "fino", "fra", "gente", "giu", "ha", "hai", "hanno", "ho", "il", "indietro \tinvece", "io", "la", "lavoro", "le", "lei", "lo", "loro", "lui", "lungo", "ma", "me", "meglio", "molta", "molti", "molto", "nei", "nella", "no", "noi", "nome", "nostro", "nove", "nuovi", "nuovo", "o", "oltre", "ora", "otto", "peggio", "pero", "persone", "piu", "poco", "primo", "promesso", "qua", "quarto", "quasi", "quattro", "quello", "questo", "qui", "quindi", "quinto", "rispetto", "sara", "secondo", "sei", "sembra \tsembrava", "senza", "sette", "sia", "siamo", "siete", "solo", "sono", "sopra", "soprattutto", "sotto", "stati", "stato", "stesso", "su", "subito", "sul", "sulla", "tanto", "te", "tempo", "terzo", "tra", "tre", "triplo", "ultimo", "un", "una", "uno", "va", "vai", "voi", "volte", "vostro"}
	} else if lang == "de" {
		stopwords = []string{"aber", "als", "am", "an", "auch", "auf", "aus", "bei", "bin", "bis", "bist", "da", "dadurch", "daher", "darum", "das", "daß", "dass", "dein", "deine", "dem", "den", "der", "des", "dessen", "deshalb", "die", "dies", "dieser", "dieses", "doch", "dort", "du", "durch", "ein", "eine", "einem", "einen", "einer", "eines", "er", "es", "euer", "eure", "für", "hatte", "hatten", "hattest", "hattet", "hier \thinter", "ich", "ihr", "ihre", "im", "in", "ist", "ja", "jede", "jedem", "jeden", "jeder", "jedes", "jener", "jenes", "jetzt", "kann", "kannst", "können", "könnt", "machen", "mein", "meine", "mit", "muß", "mußt", "musst", "müssen", "müßt", "nach", "nachdem", "nein", "nicht", "nun", "oder", "seid", "sein", "seine", "sich", "sie", "sind", "soll", "sollen", "sollst", "sollt", "sonst", "soweit", "sowie", "und", "unser \tunsere", "unter", "vom", "von", "vor", "wann", "warum", "was", "weiter", "weitere", "wenn", "wer", "werde", "werden", "werdet", "weshalb", "wie", "wieder", "wieso", "wir", "wird", "wirst", "wo", "woher", "wohin", "zu", "zum", "zur", "über"}
	} else if lang == "sv" {
		stopwords = []string{"aderton", "adertonde", "adjö", "aldrig", "alla", "allas", "allt", "alltid", "alltså", "än", "andra", "andras", "annan", "annat", "ännu", "artonde", "artonn", "åtminstone", "att", "åtta", "åttio", "åttionde", "åttonde", "av", "även", "båda", "bådas", "bakom", "bara", "bäst", "bättre", "behöva", "behövas", "behövde", "behövt", "beslut", "beslutat", "beslutit", "bland", "blev", "bli", "blir", "blivit", "bort", "borta", "bra", "då", "dag", "dagar", "dagarna", "dagen", "där", "därför", "de", "del", "delen", "dem", "den", "deras", "dess", "det", "detta", "dig", "din", "dina", "dit", "ditt", "dock", "du", "efter", "eftersom", "elfte", "eller", "elva", "en", "enkel", "enkelt", "enkla", "enligt", "er", "era", "ert", "ett", "ettusen", "få", "fanns", "får", "fått", "fem", "femte", "femtio", "femtionde", "femton", "femtonde", "fick", "fin", "finnas", "finns", "fjärde", "fjorton", "fjortonde", "fler", "flera", "flesta", "följande", "för", "före", "förlåt", "förra", "första", "fram", "framför", "från", "fyra", "fyrtio", "fyrtionde", "gå", "gälla", "gäller", "gällt", "går", "gärna", "gått", "genast", "genom", "gick", "gjorde", "gjort", "god", "goda", "godare", "godast", "gör", "göra", "gott", "ha", "hade", "haft", "han", "hans", "har", "här", "heller", "hellre", "helst", "helt", "henne", "hennes", "hit", "hög", "höger", "högre", "högst", "hon", "honom", "hundra", "hundraen", "hundraett", "hur", "i", "ibland", "idag", "igår", "igen", "imorgon", "in", "inför", "inga", "ingen", "ingenting", "inget", "innan", "inne", "inom", "inte", "inuti", "ja", "jag", "jämfört", "kan \tkanske", "knappast", "kom", "komma", "kommer", "kommit", "kr", "kunde", "kunna", "kunnat", "kvar", "länge", "längre", "långsam", "långsammare", "långsammast", "långsamt", "längst", "långt", "lätt", "lättare", "lättast", "legat", "ligga", "ligger", "lika", "likställd", "likställda", "lilla", "lite", "liten", "litet", "man", "många", "måste", "med", "mellan", "men", "mer", "mera", "mest", "mig", "min", "mina", "mindre", "minst", "mitt", "mittemot", "möjlig", "möjligen", "möjligt", "möjligtvis", "mot", "mycket", "någon", "någonting", "något", "några", "när", "nästa", "ned", "nederst", "nedersta", "nedre", "nej", "ner", "ni", "nio", "nionde", "nittio", "nittionde", "nitton", "nittonde", "nödvändig", "nödvändiga", "nödvändigt", "nödvändigtvis", "nog", "noll", "nr", "nu", "nummer", "och", "också", "ofta", "oftast", "olika", "olikt", "om", "oss", "över", "övermorgon", "överst", "övre", "på", "rakt", "rätt", "redan", "så", "sade", "säga", "säger", "sagt", "samma", "sämre", "sämst", "sedan", "senare", "senast", "sent", "sex", "sextio", "sextionde", "sexton", "sextonde", "sig", "sin", "sina", "sist", "sista", "siste", "sitt", "sjätte", "sju", "sjunde", "sjuttio", "sjuttionde", "sjutton", "sjuttonde", "ska", "skall", "skulle", "slutligen", "små", "smått", "snart", "som", "stor", "stora", "större", "störst", "stort", "tack", "tidig", "tidigare", "tidigast", "tidigt", "till", "tills", "tillsammans", "tio", "tionde", "tjugo", "tjugoen", "tjugoett", "tjugonde", "tjugotre", "tjugotvå", "tjungo", "tolfte", "tolv", "tre", "tredje", "trettio", "trettionde", "tretton", "trettonde", "två", "tvåhundra", "under", "upp", "ur", "ursäkt", "ut", "utan", "utanför", "ute", "vad", "vänster", "vänstra \tvar", "vår", "vara", "våra", "varför", "varifrån", "varit", "varken", "värre", "varsågod", "vart", "vårt", "vem", "vems", "verkligen", "vi", "vid", "vidare", "viktig", "viktigare", "viktigast", "viktigt", "vilka", "vilken", "vilket", "vill"}
	}

	return stopwords
}

// stopword The function that removes the stopwords contained in the input sentence
func stopword(s string, stopwords []string) string {
	var finalSent []string
	for _, word := range Tokenize(s) {
		if !utils.SliceContains(stopwords, word) {
			finalSent = append(finalSent, word)
		}
	}
	return utils.ListToStr(finalSent)
}

//// stopwordList The function that retrieves a list of stopwords, here in French
//func lowTfidfWordList() []string {
//	var lowTfidfWord []string
//	dir := path.Join(utils.ChuchoDir, "stopwords", "lowTfidfWord.txt")
//	stopwords, e2 := os.Open(dir)
//	if e2 != nil {
//		log.Fatalf("Error when opening file: %s", e2)
//	}
//	defer stopwords.Close()
//	fileScanner := bufio.NewScanner(stopwords)
//	for fileScanner.Scan() {
//		lowTfidfWord = append(lowTfidfWord, fileScanner.Text())
//	}
//	if e3 := fileScanner.Err(); e3 != nil {
//		log.Fatalf("Error while reading file: %s", e3)
//	}
//	// stopwords.Close()
//	return lowTfidfWord
//}
//
//// lowTfidfWord
//func lowTfidfWord(s string) string {
//	var finalSent []string
//	stopList := lowTfidfWordList()
//	for _, word := range Tokenize(s) {
//		if !utils.SliceContains(stopList, word) {
//			finalSent = append(finalSent, word)
//		}
//	}
//	return utils.ListToStr(finalSent)
//}

// lower The function that put words of a sentence in lowercase mode
func lower(s string) string {
	return strings.ToLower(s)
}

// Tokenize The function that separates the words of a sentence to facilitate their study in NLP
func Tokenize(s string) []string {
	return strings.Split(strings.Trim(s, " "), " ")
}
