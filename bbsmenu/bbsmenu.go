package bbsmenu

type BBSmenu struct {
	Description      string     `json:"description"`
	LastModify       int64      `json:"last_modify"`
	LastModifyString string     `json:"last_modify_string"`
	MenuList         []MenuList `json:"menu_list"`
}

type MenuList struct {
	CategoryContent []CategoryContent `json:"category_content"`
	CategoryTotal   uint64            `json:"category_total"`
	CategoryName    string            `json:"category_name"`
	CategoryNumber  string            `json:"category_number"`
}

type CategoryContent struct {
	BoardName     string `json:"board_name"`
	Category      uint64 `json:"category"`
	CategoryName  string `json:"category_name"`
	CategoryOrder uint64 `json:"category_order"`
	DirectoryName string `json:"directory_name"`
	Url           string `json:"url"`
}
