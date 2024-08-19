# 933_最近的请求次数_Number_of_Recent_Calls

Write a class RecentCounter to count recent requests. It has only one method: ping(int t), where t represents some time in milliseconds. Return the number of pings that have been made from 3000 milliseconds ago until now. Any ping with time in [t - 3000, t] will count, including the current ping. It is guaranteed that every call to ping uses a strictly larger value of t than before.

Example:  
Input: inputs = ["RecentCounter","ping","ping","ping","ping"], inputs = [[],[1],[100],[3001],[3002]]  
Output: [null,1,2,3,3]  

Note:  
Each test case will have at most 10000 calls to ping.   
Each test case will call ping with strictly increasing values of t.   
Each call to ping will have 1 <= t <= 10^9.

```go
type RecentCounter struct {
    pings []int
}

func Constructor() RecentCounter {
    rc := RecentCounter{}
    return rc
}

func (this *RecentCounter) Ping(t int) int {
    this.pings = append(this.pings, t)
    l := len(this.pings)
    for i := l - 1; i >= 0; i-- {
        if this.pings[i] < t-3000 {
            this.pings = this.pings[i+1:]
            break
        }
    }
    return len(this.pings)
}
/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
```