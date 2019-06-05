package cache

import (
	"testing"
	)

func TestCache_Cache(t *testing.T) {
	cacher := NewCommonCacher()
	if success, err := cacher.Cache("a", Entry{RegionName:"region_a", CityName:"cityname_a"}); success == true {
		t.Log("cache a success")
	}else{
		t.Error(err)
	}
	if success, err := cacher.Cache("b", Entry{RegionName:"region_b", CityName:"cityname_b"}); success== true {
		t.Log("cache b success")
	}else{
		t.Error(err)
	}
	t.Log(cacher.PrintContainer())
}

func TestCache_Get(t *testing.T) {
	cacher := NewCommonCacher()
	if success, err := cacher.Cache("a", Entry{RegionName:"region_a", CityName:"cityname_a"}); success == true {
		t.Log("cache a success")
	}else{
		t.Error(err)
	}
	if success, err := cacher.Cache("b", Entry{RegionName:"region_b", CityName:"cityname_b"}); success== true {
		t.Log("cache b success")
	}else{
		t.Error(err)
	}
	if entry, err := cacher.Get("a"); err == nil {
		t.Log(entry.CityName, entry.RegionName)
	}
	t.Log("--------")
	if entry, err := cacher.Get("B"); err == nil {
		t.Log(entry.CityName, entry.RegionName)
	}

}

