package data

import "testing"

func TestFetchCity(t *testing.T) {
	err := RefreshCity()
	city := GetCityRepo()
	if err == nil {
		for _, c := range *city {
			t.Logf("Found: %s", c)
		}
	} else {
		t.Error(err)
	}

}
