package main

func NewDB() *DB {
	return &DB{
		store: make(map[string]int),
	}
}

type DB struct {
	store map[string]int
}

func (db *DB) GetReceiptPointsById(id string) (int, bool) {
	points, ok := db.store[id]
	if !ok {
		return 0, false
	}

	return points, true
}

func (db *DB) SetReceiptPointsById(id string, points int) {
	db.store[id] = points
}
