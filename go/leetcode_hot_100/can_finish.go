package leetcode_hot_100

// TODO again
func canFinish(numCourses int, prerequisites [][]int) bool {
	grap := make([][]int, numCourses)
	inDeg := make([]int, numCourses)

	// 先将图构建好，并统计七每个节点的入度
	for i := 0; i < len(prerequisites); i++ {
		grap[prerequisites[i][1]] = append(grap[prerequisites[i][1]], prerequisites[i][0])
		inDeg[prerequisites[i][0]]++
	}

	// 找出入度为0的节点
	queue := make([]int, 0)
	for i := 0; i < len(inDeg); i++ {
		if inDeg[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 每次都从入度为0的节点开始，将其指向的节点，挨个拆除
	courses := 0
	for len(queue) > 0 {
		curCourse := queue[0]
		queue = queue[1:]
		courses++

		for i := 0; i < len(grap[curCourse]); i++ {
			inDeg[grap[curCourse][i]]--
			if inDeg[grap[curCourse][i]] == 0 {
				queue = append(queue, grap[curCourse][i])
			}
		}
	}

	return courses == numCourses
}
