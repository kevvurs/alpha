package data

import 	"fmt"

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
}

func (p Publication) String() string {
	return fmt.Sprintf("{publisher:%s, home:%s, imgRef:%s, hits:%d, quality:%.2f, ycred:%d, ncred:%d, owner:%s, pubId:%d}",
		p.Publisher, p.Home, p.Imgref, p.Hits, p.Quality, p.Ycred, p.Ncred, p.Owner, p.PubId)
}
