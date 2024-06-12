package model

type ConventionalCommit struct {
	Id          int
	Type        ConventionalType
	Scope       ConventionalScope
	Description string
}
