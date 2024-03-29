package controllers

import (
	"backend-template/models"
	"backend-template/repositories"
)

// CREATE
func StoreCreate(store models.Store) (uint, error) {
	return repositories.CreateStore(store)
}

// UPDATE
func StoreIDUpdate(id uint, store models.Store) error {
	return repositories.UpdateStore(store, "ID = ?", id)
}

// DELETE
func StoreIDDelete(id uint) error {
	return repositories.DeleteStore("ID = ?", id)
}

// GET
func GetStoreByID(id uint) (models.Store, error) {
	return repositories.GetStore("ID = ?", id)
}

func GetAllStores() ([]models.Store, error) {
	return repositories.GetAllStores()
}
