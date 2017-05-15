package data

import (
	"fmt"
	"strings"
	"sync"
)

type Publication struct {
	Publisher string  `json:"publisher"`
	Home      string  `json:"home"`
	Imgref    string  `json:"imgRef"`
	Hits      int     `json:"hits"`
	Quality   float32 `json:"quality"`
	Ycred     int     `json:"ycred"`
	Ncred     int     `json:"ncred"`
	Owner     string  `json:"owner"`
	PubId     int     `json:"pubId"`
	Exists    bool
}

type PubCache struct {
	cache map[int]Publication
	mux   sync.Mutex
}

type cloudsql struct {
	Connection string `yaml:"instance"`
	UserName   string `yaml:"user"`
	Password   string `yaml:"paswd"`
	Stmnt      struct {
		Use_database        string `yaml:"use_database"`
		Select_publications string `yaml:"select_publications"`
		Insert_update       string `yaml:"insert_update"`
		Insert_clobber      string `yaml:"insert_clobber"`
		Delete_publication  string `yaml:"delete_publication"`
	}
}

func (cl cloudsql) String() string {
	return fmt.Sprintf("Google Cloud SQL Config:{conn:%s, user:%s, pass:%s}",
		cl.Connection, cl.UserName, cl.Password)
}

func (p Publication) String() string {
	return fmt.Sprintf("{publisher:%s, home:%s, imgRef:%s, hits:%d, quality:%.2f, ycred:%d, ncred:%d, owner:%s, pubId:%d}",
		p.Publisher, p.Home, p.Imgref, p.Hits, p.Quality, p.Ycred, p.Ncred, p.Owner, p.PubId)
}

func (pc PubCache) String() string {
	var body []string
	for k, v := range pc.cache {
		body = append(body, fmt.Sprintf("%d:%s", k, v))
	}
	contents := strings.Join(body, ",")
	return "[" + contents + "]"
}
