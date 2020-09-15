package string

import (
	"crypto/md5"
	"encoding/hex"
	//"fmt"
	"net/url"
	"sort"
)

type Encoder struct {
	Params map[string]string
	Secret string
}

func NewEncoder(secret string) *Encoder {
	return &Encoder{Secret: secret}
}

func (e *Encoder) Add(k string, v string) {
	e.Params[k] = v
}

func (e *Encoder) Token() string {
	var stringSlice []string
	for key := range e.Params {
		stringSlice = append(stringSlice, key)
	}

	sort.Strings(stringSlice)

	//fmt.Printf("stringSlice:%v", stringSlice)

	s := ""
	for _, val := range stringSlice {
		s += e.Params[val]
	}

	return Md5(Md5(s) + e.Secret)
}

func Md5(str string) string {
	md5ctx := md5.New()
	md5ctx.Write([]byte(str))
	return hex.EncodeToString(md5ctx.Sum(nil))
}

func (e *Encoder) GetURLEncodeParams() string {
	return e.GetURLValues().Encode()
}

func (e *Encoder) GetURLValues() url.Values {
	e.Params["t"] = e.Token()
	v := url.Values{}
	for key, val := range e.Params {
		v.Add(key, val)
	}
	return v
}
