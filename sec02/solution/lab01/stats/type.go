package stats

import "github.com/egbordzor/learn-glft/shared/sec02/db"

type RecordProvider func() (name, gender string, income db.Currency, err error)
