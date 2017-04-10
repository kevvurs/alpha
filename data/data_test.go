package data

import "testing"


func TestFetchCity(t *testing.T) {
	city, err := FetchCity()
	if err == nil {
		for _, c := range city {
			t.Logf("Found: %s", c)
		}
	} else {
		t.Error(err)
	}

}