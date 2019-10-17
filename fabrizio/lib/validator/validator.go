package validator


type (
	validator struct {

	}
)

func NewValidator() *validator {
	return &validator{}
}

func (v *validator) ValidateFields(rows [][]string) (bool, error)  {
	headers := rows[0]
	for rowNumber, row := range rows {
		if rowNumber == 0 {
			continue
		}
		for k, fieldValue := range row {
			isValid, err := v.ValidateField(headers[k], fieldValue)

			if err != nil {
				return isValid, err
			}
		}

	}
	return true, nil
}
/**
1. Validate size
2. Mandatory
3. Positive integer
4. Max number
 */

func (v *validator) ValidateField(name string, value interface{}) (bool, error) {
	return true, nil
}

/**
		'SellerSku' => [
            'validators' => [
                'Length' => [
                    'min' => SellerCenter_Model_Catalog_Product::MIN_SKU_LENGTH,
                    'max' => SellerCenter_Model_Catalog_Product::MAX_SKU_LENGTH,
                ],
            ],
            'mandatory' => true,
        ],
        'ParentSku' => [
            'validators' => [
                'Length' => [
                    'min' => ParentSkuValidator::CONSTRAINT_LENGTH_MIN,
                    'max' => ParentSkuValidator::CONSTRAINT_LENGTH_MAX,
                ],
            ],
            'mandatory' => false,
        ],
        'Status' => [
            'validators' => [
                'Status' => [],
            ],
            'mandatory' => false,
        ],
        'Name' => [
            'validators' => [
                'Length' => [
                    'min' => SellerCenter_Model_Catalog_Product::MIN_NAME_LENGTH,
                    'max' => SellerCenter_Model_Catalog_Product::MAX_NAME_LENGTH,
                ],
            ],
            'mandatory' => true,
        ],
        'Variation' => [
            'validators' => [
                'Length' => [
                    'min' => SellerCenter_Model_Catalog_Product::MIN_VARIATION_LENGTH,
                    'max' => SellerCenter_Model_Catalog_Product::MAX_VARIATION_LENGTH,
                ],
            ],
            'mandatory' => false,
        ],
        'PrimaryCategory' => [
            'validators' => [
                'PositiveInteger' => [],
            ],
            'mandatory' => true,
        ],
        'Categories' => [
            'validators' => [
                'Categories' => [],
            ],
            'mandatory' => false,
        ],
        'BrowseNodes' => [
            'validators' => [
                'BrowseNodes' => [],
            ],
            'mandatory' => false,
        ],
        'Description' => [
            'validators' => [
                'Length' => [
                    'min' => self::TEXTAREA_MIN_LENGTH,
                    'max' => self::TEXTAREA_MAX_LENGTH,
                ],
            ],
            'mandatory' => true,
        ],
        'Brand' => [
            'validators' => [],
            'mandatory' => true,
        ],
        'Price' => [
            'validators' => [
                'Price' => [
                    'max' => SellerCenter_Model_Catalog_Product::MAX_PRICE_LENGTH,
                ],
            ],
            'mandatory' => true,
        ],
        'SalePrice' => [
            'validators' => [
                'SpecialPrice' => [
                    'max' => SellerCenter_Model_Catalog_Product::MAX_PRICE_LENGTH,
                ],
            ],
            'mandatory' => false,
        ],
        'SaleStartDate' => [
            'validators' => [
                'SpecialStartDate' => [],
            ],
            'mandatory' => false,
        ],
        'SaleEndDate' => [
            'validators' => [
                'specialEndDate' => [],
            ],
            'mandatory' => false,
        ],
        'TaxClass' => [
            'validators' => [],
            'mandatory' => true,
        ],
        'ProductId' => [
            'validators' => [
                'Barcode' => [],
            ],
            'mandatory' => false,
        ],
        'Quantity' => [
            'validators' => [
                'PositiveInteger' => [],
            ],
            'mandatory' => false,
        ],
        'ProductSin' => [
            'validators' => [],
            'mandatory' => false,
        ],
        'ProductGroup' => [
            'validators' => [
                'ProductGroup' => [],
                'Length' => [
                    'lengthMin' => SellerCenter_Model_Catalog_Product::MIN_PRODUCT_GROUP_LENGTH,
                    'lengthMax' => SellerCenter_Model_Catalog_Product::MAX_PRODUCT_GROUP_LENGTH,
                ],
            ],
            'mandatory' => false,
        ],
 */
