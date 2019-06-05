package cache

/*
func TestLRUCache_Cache(t *testing.T) {
	cache, err  := NewLRUCache(10)
	if err != nil {
		t.Error(err)
	}
	if success, err := cache.Cache("a", Entry{RegionName:"region_a", CityName:"city_a"}); success == true {
		t.Log("lru cache a success")
	}else{
		t.Error(err)
	}
	if success, err := cache.Cache("b", Entry{RegionName:"region_b", CityName:"city_b"}); success == true {
		t.Log("lru cache b success")
	}else{
		t.Error(err)
	}
	t.Log(cache.PrintLRUCacheList())
	entry := cache.PrintLRUCacheContainer()["a"].Value
	t.Log(entry.(Entry).RegionName, entry.(Entry).RegionName)
	if success, err := cache.Cache("a", Entry{RegionName:"region_a", CityName:"city_a"}); success == true {
		t.Log("lru cache a success")
	}else{
		t.Error(err)
	}
	t.Log(cache.PrintLRUCacheList())
}


func TestLRUCache_Get(t *testing.T) {
	cache, err  := NewLRUCache(10)
	if err != nil {
		t.Error(err)
	}
	cache.Cache("a", Entry{RegionName:"region_a", CityName:"city_a"})
	t.Log(cache.PrintLRUCacheList().Len())
	cache.Cache("b", Entry{RegionName:"region_b", CityName:"city_b"})
	t.Log(cache.PrintLRUCacheList().Len())
	cache.Cache("c", Entry{RegionName:"region_c", CityName:"city_c"})
	t.Log(cache.PrintLRUCacheList().Len())

	cache.Get("a")
	t.Log("last: ", cache.PrintLRUCacheList())
	t.Log("last: ", cache.PrintLRUCacheList())
	cache.Get("b")
	cache.Get("c")
	cache.Get("c")
	t.Log("last: ", cache.PrintLRUCacheList())
	elem, _ := cache.Get("c")
	value := cache.PrintLRUCacheList().Remove(elem)
	t.Log("last: ", cache.PrintLRUCacheList(), "removed: ", value.(Entry).RegionName)
}

*/
