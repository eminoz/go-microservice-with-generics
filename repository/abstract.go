package repository

func InsertOneEntity[I IBaseEntity](i I) I {
	return i
}

type IBaseEntity interface {
}
