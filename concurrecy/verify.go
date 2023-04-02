package concurrecy

type VerificaWebsite func(string) bool
type result struct {
	string
	bool
}

func VerifyWebsite(vw VerificaWebsite, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// send Syntax
			resultChannel <- result{u, vw(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// getting Syntax
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results

}
