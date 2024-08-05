package graph

// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
//
// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，
// 表示如果要学习课程 ai 则 必须 先学习课程  bi 。 [bi] -> [ai]
//
// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
// 核心拓扑排序
func canFinish(numCourses int, prerequisites [][]int) bool {
	//节点入度 表示课程的前置依赖课程数量
	inDegree := make([]int, numCourses)
	//节点出度 表示课程的下一门课程 多个
	outDegree := make([][]int, numCourses)
	//可以学习的课程 也就是课程的前置依赖课程数量为0
	canLearn := make([]int, 0, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		inDegree[prerequisites[i][0]]++
		outDegree[prerequisites[i][1]] = append(outDegree[prerequisites[i][1]], prerequisites[i][0])
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
	return len(canLearn) == numCourses
}
