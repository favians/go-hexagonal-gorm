package paginator

type Meta struct {
	Page         int  `json:"page"`
	RowPerPage   int  `json:"row_per_page"`
	NextPage     bool `json:"next_page"`
	PreviousPage bool `json:"previous_page"`
}

func (meta *Meta) BuildMeta(dataLength int, page int, rowPerPage int) {

	rowPerPage = rowPerPage - 1

	meta.Page = page
	meta.RowPerPage = rowPerPage
	meta.NextPage = false

	if dataLength > rowPerPage {
		meta.NextPage = true
	}

	if (dataLength-1 <= rowPerPage) && (page != 1) {
		meta.PreviousPage = true
	}
}
