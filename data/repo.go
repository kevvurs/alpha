package data

import "log"

// In-memory, thread-safe cache
var pubCache PubCache

// Load the data from DB to cache
func RefreshCity() error {
	err := fetchCity(&pubCache)
	if err == nil {
		log.Println("Pub cache refreshed")
	} else {
		log.Printf("Error: refreshing cache: %v", err)
	}
	return err
}

func GetCityRepo() *PubCache {
	return &pubCache
}

func (pc *PubCache) Push(p *Publication) {
	pc.mux.Lock()
	defer pc.mux.Unlock()
	k, v := pc.cache[p.PubId];
	if v {
		log.Print("Overwriting <A> with <B>\n  A:%s\n  B:%s",k,p)
	}
	pc.cache[p.PubId] = *p
}

func (pc *PubCache) Pull(id *int) *Publication {
	pc.mux.Lock()
	defer pc.mux.Unlock()
	return &pc.cache[id]
}