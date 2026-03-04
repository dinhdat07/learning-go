package models

type Mode string

const (
	ExpressionMode   Mode = "expression"
	EquationMode     Mode = "equation"
	LinearSystemMode Mode = "linear_system"
)
