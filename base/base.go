package base

const LangError = "lang variable value is not recognized, check the lang code you want, " +
	"for example: 'de' for German, 'en' for English, 'fr' for French"

type (
	// Lang is an alias for string but is more likely accepting 2 letters code for lang
	Lang string

	// Stemms and StemmDict are types required to create a Stemmer object
	Stemms    map[Lang]StemmDict // every lang
	StemmDict map[string]string  // local lang

	// GlobalStopwords and StopList are the types required to create a Stopwords object
	GlobalStopwords map[Lang]StopList // every lang
	StopList        []string          // local lang
)
