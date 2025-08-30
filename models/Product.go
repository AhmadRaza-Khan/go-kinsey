package models

type Product struct {
	ID                     uint   `gorm:"primaryKey"`
	NorthItemNumber        string `gorm:"size:50"`
	SouthItemNumber        string `gorm:"size:50"`
	Description1           string `gorm:"size:255"`
	Description2           string `gorm:"size:255"`
	Brand                  string `gorm:"size:100"`
	VendorItemNumber       string `gorm:"size:50"`
	MSRP                   float64
	MAPPrice               float64
	BarCode                string  `gorm:"size:50"`
	RHLH                   *string `gorm:"size:50"`
	Color1                 string  `gorm:"size:50"`
	Color2                 string  `gorm:"size:50"`
	Size                   string  `gorm:"size:50"`
	ProductLength          float64
	ProductWidth           float64
	ProductHeight          float64
	Weight                 float64
	BulletFeatures         string  `gorm:"type:text"`
	ExtendedText           string  `gorm:"type:text"`
	CountryOfOrigin        string  `gorm:"size:100"`
	Flammable              *string `gorm:"size:50"`
	Hazardous              *string `gorm:"size:50"`
	DateCreated            string  `gorm:"size:50"`
	ItemCategoryCode       string  `gorm:"size:50"`
	ProductGroupCode       string  `gorm:"size:50"`
	ProductSubGroup1       string  `gorm:"size:50"`
	ProductSubGroup2       string  `gorm:"size:50"`
	PackSize               string  `gorm:"size:50"`
	Inactive               string  `gorm:"size:10"`
	Blocked                string  `gorm:"size:10"`
	Name                   string  `gorm:"size:255"`
	IncludeExcludeGroup    *string `gorm:"size:50"`
	Prop65Applies          *string `gorm:"size:50"`
	Prop65CancerHarm       *string `gorm:"size:50"`
	Prop65ReproductiveHarm *string `gorm:"size:50"`
	Prop65Chemical         *string `gorm:"size:255"`
}
