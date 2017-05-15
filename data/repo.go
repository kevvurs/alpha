package data

import "log"

// In-memory, thread-safe cache
var pubCache = initCache()

// Initialize the in-memory cache
func initCache() PubCache {
	pc := new(PubCache)
	pc.cache = make(map[int]Publication)
	return *pc
}

// Returns the address of the in-memory cache
func GetRepo() *PubCache {
	return &pubCache
}

// Load the data from DB to cache
func (pc *PubCache) Refresh() error {
	pc.mux.Lock()
	defer  pc.mux.Unlock()

	// Re-flash cache contents
	err := fetchAll(pc)
	if err == nil {
		log.Println("Pub cache refreshed")
	} else {
		log.Printf("Error: refreshing cache: %v\n", err)
	}
	return err
}

// Relinquish the cache's contents
func (pc *PubCache) Clean() {
	pc.mux.Lock()
	defer pc.mux.Unlock()
	pc.cache = make(map[int]Publication)
}

// Clobbering update a cache element
func (pc *PubCache) Push(p *Publication) {
	pc.mux.Lock()
	defer pc.mux.Unlock()
	k, v := pc.cache[p.PubId];
	if v {
		log.Printf("Overwriting <A> with <B>\n  A:%s\n  B:%s\n",k,p)
	}
	pc.cache[p.PubId] = *p
}

// Send element to DP as update or insert and populate the cache
func (pc *PubCache) PushDeep(p *Publication) {
	if err := upsert(p); err != nil {
		log.Printf("Aborting push: %v\n",err)
	} else {
		pc.Push(p)
	}
}

//
func (pc *PubCache) Pull(id *int) *Publication {
	pc.mux.Lock()
	defer pc.mux.Unlock()
	p, _ := pc.cache[*id]
	return &p
}