package database

type Database interface {
	Get()
	Create()
	Update()
	Delete()
}
