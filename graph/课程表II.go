package graph

// 现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。
// 给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。
//
// 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
// 返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。
func findOrder(numCourses int, prerequisites [][]int) []int {
	//节点入度 表示课程的前置依赖课程数量
	inDegree := make([]int, numCourses)
	//节点出度 表示课程的下一门课程 多个
	outDegree := make([][]int, numCourses)
	//可以学习的课程 也就是课程的前置依赖课程数量为0
	canLearn := make([]int, 0, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		//[bi] -> [ai]
		ai := prerequisites[i][0]
		bi := prerequisites[i][1]
		inDegree[ai]++
		outDegree[bi] = append(outDegree[bi], ai)
	}
	//如果课程的前置依赖课程数量为0 则加入到可以学习的课程列表中
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			canLearn = append(canLearn, i)
		}
	}

	for i := 0; i < len(canLearn); i++ {
		//学习课程
		learn := canLearn[i]
		//课程的下一门课程
		outCourses := outDegree[learn]
		for _, outCourse := range outCourses {
			//下一门课程前置依赖减一
			inDegree[outCourse]--
			//下一门课程能否开始学
			if inDegree[outCourse] == 0 {
				canLearn = append(canLearn, outCourse)
			}
		}
	}
	//已经学习完成
	if len(canLearn) == numCourses {
		return canLearn
	}
	return []int{}
}
