package store

import (
	"truffle/utils"
)

type EmailRepo struct {
	timer *utils.Timer
}

func NewEmailRepo() *EmailRepo {
	return &EmailRepo{
		timer: utils.NewTimer(60, 100, 4, 0.7),
	}
}

func (repo *EmailRepo) Put(expire int) {

}
