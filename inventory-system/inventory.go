package inventorysystem

import (
	"context"
	"inventory-system/api/model/database/repository"
	inventory "inventory-system/gen/inventory"
	"log"
)

// inventory service example implementation.
// The example methods log the requests and return zero values.
type inventorysrvc struct {
	logger *log.Logger
}

// NewInventory returns the inventory service implementation.
func NewInventory(logger *log.Logger) inventory.Service {
	return &inventorysrvc{logger}
}

func (s *inventorysrvc) Hello() string {
	return "HELLO WORLD"
}

// Create implements create.
func (s *inventorysrvc) Create(ctx context.Context, p *inventory.CreateProductPayload) (res string, err error) {
	repository.SaveProduct(p)
	s.logger.Print("inventory.create")
	return
}

// Update implements update.
func (s *inventorysrvc) Update(ctx context.Context, p *inventory.UpdateProductPayload) (res string, err error) {
	repository.UpdateProduct(p)
	s.logger.Print("inventory.update")
	return
}

// Find implements find.
func (s *inventorysrvc) Find(ctx context.Context, p *inventory.FindPayload) (res *inventory.Findproductresult, err error) {
	var product = repository.FindProduct(p.ProductID)
	res = &inventory.Findproductresult{
		ProductID:          &p.ProductID,
		ProductName:        &product.ProductName,
		ProductDescription: &product.ProductDescription,
		ProductMinStock:    &product.ProductMinStock,
	}
	s.logger.Print("inventory.find")
	return
}

// Delete implements delete.
func (s *inventorysrvc) Delete(ctx context.Context, p *inventory.DeletePayload) (res string, err error) {
	repository.DeleteProduct(p.ProductID)
	s.logger.Print("inventory.delete")
	return
}
