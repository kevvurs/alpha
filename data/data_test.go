package data

import "testing"

func TestFetchCity(t *testing.T) {
	err := RefreshCity()
	publications := GetCityRepo()
	if err == nil {
		for _, c := range publications.cache {
			t.Logf("Found: %s", c)
		}
	} else {
		t.Error(err)
	}

}
