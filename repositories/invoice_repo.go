package repositories

import (
	"backend-template/database"
	"backend-template/models"
)

//======GET======

// TO KNOW IF THE INVOICE HISTORY EXIST OR NOT
func IsInvoiceHistoryExist(query string, val ...interface{}) bool {
	var invoice_history models.InvoiceHistory
	database.DB.Where(query, val...).Last(&invoice_history)
	return invoice_history.ID != 0
}

// TO GET A INVOICE HISTORY
func GetInvoiceHistory(query string, val ...interface{}) (models.InvoiceHistory, error) {
	var invoice_history models.InvoiceHistory
	db := database.DB.Where(query, val...).Last(&invoice_history)
	return invoice_history, db.Error
}

// TO GET AN ARRAY OF INVOICE HISTORYS (NOT ALL BUT CAN ALL)
func GetInvoiceHistories(query string, val ...interface{}) ([]models.InvoiceHistory, error) {
	var invoice_histories []models.InvoiceHistory
	db := database.DB.Where(query, val...).Find(&invoice_histories)
	return invoice_histories, db.Error
}

// TO GET ALL INVOICE HISTORYS
func GetAllInvoiceHistories() ([]models.InvoiceHistory, error) {
	var invoice_histories []models.InvoiceHistory
	db := database.DB.Find(&invoice_histories)
	return invoice_histories, db.Error
}

// CREATE INVOICE HISTORY
func CreateInvoiceHistory(invoice_history models.InvoiceHistory) (uint, error) {
	db := database.DB.Create(&invoice_history)
	return invoice_history.ID, db.Error
}

//NO UPDATE & DELETE
