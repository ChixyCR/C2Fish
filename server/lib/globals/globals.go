package globals

// init globalsMap
var (
	GLOBALSMAP = make(map[string]interface{})
)

// set globalsMap
func Set(key string, value interface{}, force bool ) bool{
	if _,ok := GLOBALSMAP[key]; !ok && !force{
		return false
	}
	GLOBALSMAP[key] = value
	return true
}

// get globalsMap
func Get(key string) interface{} {
	if v,ok := GLOBALSMAP[key];ok{
		return v
	}
	return nil
}