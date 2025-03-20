package entities

import (
	"errors"
)

type GroupByEnum string

const (
	Income  GroupByEnum = "income"
	Expense GroupByEnum = "expense"
)

type Category struct {
	ID        string
	GroupType GroupByEnum
	Name      string
}

func NewCategory(groupType GroupByEnum, name string) (*Category, error) {
	if name == "" {
		return nil, errors.New("category name cannot be empty")
	}

	return &Category{
		ID:        generateUUID(),
		GroupType: groupType,
		Name:      name,
	}, nil
}

func (c *Category) UpdateName(name string) error {
	if name == "" {
		return errors.New("category name cannot be empty")
	}
	c.Name = name
	return nil
}
