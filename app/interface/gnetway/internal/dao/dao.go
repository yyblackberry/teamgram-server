// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/config"
)

type Dao struct {
	Session *Session
}

func New(c config.Config) *Dao {
	return &Dao{
		Session: NewSession(c),
	}
}
