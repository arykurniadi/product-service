package transformers

type (
	Transformer struct {
		Data interface{} `json:"data"`
	}

	CollectionTransformer struct {
		Data []interface{} `json:"data"`
	}

	CollectionPagingTransformer struct {
		Meta interface{}   `json:"meta"`
		Data []interface{} `json:"data"`
	}

	SingleCollectionPagingTransformer struct {
		Meta interface{} `json:"meta"`
		Data interface{} `json:"data"`
	}
)
