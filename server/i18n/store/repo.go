package store

import (
	"sync"
	"truffle/utils"
)

type Store struct {
	dict map[string]*utils.Conf
}

func NewStore() *Store {
	return &Store{
		dict: make(map[string]*utils.Conf),
	}
}

// ? lazy singleton pattern

func GetStore() *Store {
	once.Do(func() {
		store = NewStore()
		store.Init(Lopt.Langs)
	})
	return store
}

func (s *Store) Get(lang string, msg string) (string, bool) {
	ok, dic := s.HasLang(lang)
	if !ok {
		return msg, false
	}
	ret := dic.GetString(msg)
	if len(ret) <= 0 {
		ret = msg
	}
	return ret, true
}

func (s *Store) Set(lang, k, v string) {
	s.dict[lang].Set(k, v)
}

func (s *Store) Init(langs []string) {
	for _, lang := range langs {
		s.dict[lang] = utils.NewConf(lang, "yaml", "./lang")
	}
}

func (s *Store) HasLang(lang string) (bool, *utils.Conf) {
	if dic, ok := s.dict[lang]; ok {
		return true, dic
	}
	return false, nil
}

var (
	store *Store
	once  sync.Once
	Lopt  LangOpt
)
