package leetcode_hot_100

// TODO three - 课程表
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	inDeg := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
		inDeg[prerequisites[i][0]]++
	}

	queue := make([]int, 0)
	for i := 0; i < len(inDeg); i++ {
		if inDeg[i] == 0 {
			queue = append(queue, i)
		}
	}

	courses := 0
	for len(queue) > 0 {
		curCourse := queue[0]
		queue = queue[1:]
		courses++

		for i := 0; i < len(graph[curCourse]); i++ {
			inDeg[graph[curCourse][i]]--
			if inDeg[graph[curCourse][i]] == 0 {
				queue = append(queue, graph[curCourse][i])
			}
		}
	}

	return courses == numCourses
}
