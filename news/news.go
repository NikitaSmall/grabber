package news

var (
	searchClause chan string
	reset        chan bool

	returnClause chan string
	result       chan []Topic
)

func init() {
	searchClause = make(chan string)
	reset = make(chan bool)

	returnClause = make(chan string)
	result = make(chan []Topic)
}

func SearchFor(category string) {
	searchClause <- category
}

func ResultFor(category string) []Topic {
	returnClause <- category
	return <-result
}

func ResetArchive() {
	reset <- true
}

func (a Archive) CollectNews() {
	for {
		select {
		case category := <-searchClause:
			a.collectNews(category)
		case category := <-returnClause:
			result <- a.data(category)
		case <-reset:
			a.resetData()
		}
	}
}
