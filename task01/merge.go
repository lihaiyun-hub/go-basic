package main
import "sort"

func main(){

}
func merge(intervals [][]int) [][]int {
    if len(intervals) <= 1 {
        return intervals
    }

    // 1. 按起始点排序
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] < intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })

    merged := make([][]int, 0)
    merged = append(merged, intervals[0]) // 第一个区间直接加入

    for i := 1; i < len(intervals); i++ {
        last := merged[len(merged)-1] // 上一个合并后的区间
        curr := intervals[i]          // 当前区间

        // 2. 判断是否重叠或相接
        if curr[0] <= last[1] {
            // 合并：更新 end 为较大值
            if last[1] < curr[1] {
                last[1] = curr[1]
            }
        } else {
            // 不重叠，加入新区间
            merged = append(merged, curr)
        }
    }

    return merged
}