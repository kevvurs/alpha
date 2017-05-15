package data

import "testing"

func TestRepoFetch(t *testing.T) {
	repo := GetRepo()
	if err := repo.Refresh(); err == nil {
		t.Log("Printing full cache...")
		t.Logf("%s", repo)
	} else {
		t.Error(err)
	}
}

func TestDBUpsert(t *testing.T) {
	repo := GetRepo()

	insPub := Publication{
		Publisher: "InfoWars",
		Home:      "www.infowars.com",
		Imgref:    "img/infowars.png",
		Hits:      11,
		Quality:   0.00,
		Ycred:     0,
		Ncred:     0,
		Owner:     "Alex Jones",
		PubId:     4,
		Exists:    true,
	}

	if err := repo.Refresh(); err == nil {
		t.Log("Initial cache...")
		t.Logf("%s", repo)
		t.Log("Pushing (deep)...")
		repo.PushDeep(&insPub)
		t.Log("New cache...")
		t.Logf("%s", repo)
	} else {
		t.Error(err)
	}
}

func TestConfig(t *testing.T)  {
	t.Logf("%s",sqlConf)
	t.Logf("%s",sqlConf.Stmnt.Use_database)
	t.Logf("%s",sqlConf.Stmnt.Select_publications)
	t.Logf("%s",sqlConf.Stmnt.Insert_update)
	t.Logf("%s",sqlConf.Stmnt.Insert_clobber)
}