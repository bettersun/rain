package tree

// Data 数据
type Data struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Amount     int    `json:"amount"`
	ParentId   string `json:"parent_id"`
	ParentName string `json:"parent_name"`
}

// Node 树节点
type Node struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Amount     int     `json:"amount"`
	ParentId   string  `json:"parent_id"`
	ParentName string  `json:"parent_name"`
	Children   []*Node `json:"children"`
	//Parent     *Node   `json:"parent"`
}

// RowData 数据行
type RowData struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`

	Id01     string `json:"id01"`
	Name01   string `json:"name01"`
	Amount01 int    `json:"amount01"`

	Id02     string `json:"id02"`
	Name02   string `json:"name02"`
	Amount02 int    `json:"amount02"`

	Id03     string `json:"id03"`
	Name03   string `json:"name03"`
	Amount03 int    `json:"amount03"`
}
