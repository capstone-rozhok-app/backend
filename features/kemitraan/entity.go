package kemitraan

type MitraCore struct{
	ID				int
	StatusKemitraan	string
}

type UsecaseInterface interface{
	PutKemitraan (id int, data MitraCore)(row int, err error)
}

type DataInterface interface{
	UpdateKemitraan (id int, data MitraCore)(row int, err error)
}
