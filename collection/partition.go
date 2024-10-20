package collection

func MakeOrderedPartitions(sum, length int) [][]int {
	if length == 0 {
		if sum == 0 {
			return [][]int{{0}}
		} else {
			return [][]int{}
		}
	}
	if length == 1 {
		return [][]int{{sum}}
	}
	if sum < length {
		return [][]int{}
	}
	var partitions [][]int
	for first := 1; first <= sum-(length-1); first++ {
		subPartitions := MakeOrderedPartitions(sum-first, length-1)
		for i := 0; i < len(subPartitions); i++ {
			curPartition := make([]int, len(subPartitions[i])+1)
			curPartition[0] = first
			for j := 0; j < len(subPartitions[i]); j++ {
				curPartition[j+1] = subPartitions[i][j]
			}
			partitions = append(partitions, curPartition)
		}
	}
	return partitions
}
