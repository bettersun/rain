package tree

// BuildTree 构建节点树
//
//	仅限各级数据之间构建
func BuildTree(node *Node, allData [][]Data, layer int) {
	if layer < 0 {
		return
	}

	var children []*Node
	for _, v := range allData[layer] {
		if v.ParentId == node.Id {

			var nd Node

			nd.Id = v.Id
			nd.Name = v.Name
			nd.Amount = v.Amount
			nd.ParentId = v.ParentId
			nd.ParentName = v.ParentName

			// 递归调用 继续构建子节点的树
			BuildTree(&nd, allData, layer-1)
			children = append(children, &nd)
		}
	}
	node.Children = children
}
