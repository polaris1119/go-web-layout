package model

import "strconv"

type Demo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (d *Demo) String() string {
	return "id:" + strconv.Itoa(d.ID) + ";name:" + d.Name
}
