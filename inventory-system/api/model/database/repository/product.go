package repository

import (
	"inventory-system/api/model/database"
	"inventory-system/api/model/database/entity"
	inventory "inventory-system/gen/inventory"
	"log"
)

func SaveProduct(p *inventory.CreateProductPayload) {
	product := entity.Product{
		ProductName:        p.ProductName,
		ProductDescription: *p.ProductDescription,
		ProductMinStock:    *p.ProductMinStock,
	}
	database.Db.Create(&product)
}

func UpdateProduct(p *inventory.UpdateProductPayload) {
	// find
	var product entity.Product
	if err := database.Db.Where("product_id = ?", p.ProductID).First(&product).Error; err != nil {
		log.Fatal("NOT FOUND PRODUCT")
	}
	// update
	product.ProductName = *p.ProductName
	product.ProductDescription = *p.ProductDescription
	product.ProductMinStock = *p.ProductMinStock
	database.Db.Save(&product)
}

func FindProduct(id int) entity.Product {
	var product entity.Product
	if err := database.Db.Where("product_id = ?", id).First(&product).Error; err != nil {
		log.Fatal("NOT FOUND PRODUCT")
	}
	return product
}

func DeleteProduct(id int) {
	// find
	var product entity.Product
	if err := database.Db.Where("product_id = ?", id).First(&product).Error; err != nil {
		log.Fatal("NOT FOUND PRODUCT")
	}
	// delete
	database.Db.Delete(&product)

}
