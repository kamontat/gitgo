package model

import (
	"strconv"

	"github.com/kamontat/go-log-manager"
	"github.com/spf13/viper"
)

type List struct {
	key     string
	headers []Header
}

func (l *List) Setup(key string) *List {
	l.key = key
	return l
}

func (l *List) Load(vp *viper.Viper) *List {
	if vp == nil {
		om.Log.ToDebug("commit list", "viper is nil, cannot load list")
		return l
	}

	// reset list
	l.headers = []Header{}

	om.Log.ToDebug("commit list", "load commit list from "+vp.ConfigFileUsed())
	return l.Merge(vp)
}

func (l *List) Merge(vp *viper.Viper) *List {
	if vp == nil {
		om.Log.ToWarn("Merge list", "config object is nil, cannot merge")
		return l
	}

	if vp.Get(l.key) == nil {
		om.Log.ToWarn("Merge list", l.key+" key not exist @"+vp.ConfigFileUsed())
		return l
	}

	if l.headers == nil {
		l.headers = []Header{}
	}

	om.Log.ToVerbose("Merge list", vp.ConfigFileUsed())
	for i, element := range vp.Get(l.key).([]interface{}) {
		cm := element.(map[interface{}]interface{})

		if _, ok := cm["type"]; !ok {
			om.Log.ToError("Load list", "Have invalid type format")
			break
		}

		if _, ok := cm["value"]; !ok {
			om.Log.ToError("Load list", "Value of type="+cm["type"].(string)+" is not exist.")
		}

		commitHeader := Header{
			Type:  cm["type"].(string),
			Value: cm["value"].(string),
		}

		om.Log.ToVerbose("List header "+strconv.Itoa(i), commitHeader.String())
		l.headers = append(l.headers, commitHeader)
	}
	return l
}

// MakeList will return header list
func (l *List) MakeList() []string {
	var list []string
	for _, commits := range l.headers {
		list = append(list, commits.Format())
	}
	om.Log.ToVerbose("list size", len(list))
	return list
}

// IsContain check is list have element
func (l *List) IsContain() bool {
	return len(l.headers) > 0
}
