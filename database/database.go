package database

// Database will be the interface used to handle persistent data
type Database interface {
	Get()
	Create()
	Update()
	Delete()
}
