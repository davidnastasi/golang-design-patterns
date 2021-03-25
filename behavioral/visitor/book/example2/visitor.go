package example2

import "fmt"

type ProductInfoRetriever interface {
	GetPrice() float32
	GetName() string
}

type Visitor interface {
	Visit(ProductInfoRetriever)
}

type Visitable interface {
	Accept(Visitor)
}

type Product struct {
	Price float32
	Name  string
}

func (p *Product) GetPrice() float32 {
	return p.Price
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) Accept(v Visitor) {
	v.Visit(p)
}

type Rice struct {
  Product
}

type Pasta struct {
  Product
}

type PriceVisitor struct {
  Sum float32
}

func (pv *PriceVisitor) Visit(p ProductInfoRetriever) {
	pv.Sum += p.GetPrice()
}

type NamePrinter struct {
	ProductList string
}

func (n *NamePrinter) Visit(p ProductInfoRetriever) {
	n.ProductList = fmt.Sprintf("%s\n%s", p.GetName(), n.ProductList)
}

type Fridge struct {
  Product
}

func (f *Fridge) GetPrice() float32 {
	return f.Product.Price + 20
}

/*func (f *Fridge) GetName() string {
	return f.Product.Name
}*/

func (f *Fridge) Accept(v Visitor) {
	v.Visit(f)
}