package patreon

import (
	"net/url"
	"strings"
	"time"
)

func (dt *DateTime) UnmarshalJSON(buffer []byte) error {
	var err error
	dt.Time, err = time.Parse(time.RFC3339, strings.Trim(string(buffer), `"`))
	return err
}

func (u *URL) UnmarshalJSON(buffer []byte) error {
	var err error
	u.URL, err = url.Parse(strings.Trim(string(buffer), `"`))
	return err
}

func (attr *Attributes) AmmountDollars() float64 {
	return float64(attr.AmountCents) / 100
}
