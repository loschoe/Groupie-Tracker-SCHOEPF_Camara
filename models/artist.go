package models

import "strings"

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Image        string   `json:"image"`
}

var Artists []Artist

func (a Artist) NbMembers() int {
	return len(a.Members)
}

func (a Artist) MembersList() string {
	return strings.Join(a.Members, ", ")
}
