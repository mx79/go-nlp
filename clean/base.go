package clean

const LangError = "lang variable value is not recognized, check the lang code you want, " +
	"for example: 'DE' for German, 'EN' for English, 'FR' for French"

// Language constant are here for users who don't know the exact code of their country,
// because there are many ones: 2 or 3 letters.
//
// The use of this const avoids any error in language matching, an error can still be raised if
// the target object has no language data for its usage.
const (
	FR lang = "fr"
	ES lang = "es"
	IT lang = "it"
	EN lang = "en"
)

type (
	// Lang is an alias for string but is more likely accepting 2 letters code for lang.
	lang string

	// Stemms and StemmDict are types required to create a Stemmer object.
	stemms    map[lang]stemmDict // every lang
	stemmDict map[string]string  // local lang

	// GlobalStopwords and StopList are the types required to create a Stopwords object.
	globalStopwords map[lang]stopList // every lang
	stopList        []string          // local lang
)
