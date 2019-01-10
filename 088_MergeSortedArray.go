//O(m+n) time, O(1) space
func merge(nums1 []int, m int, nums2 []int, n int)  {
    cur := m+n-1
    for m,n = m-1,n-1;m>=0 && n>=0;cur-- {
        if nums1[m]>=nums2[n] {
            nums1[cur] = nums1[m]
            m--
        } else {
            nums1[cur] = nums2[n]
            n--
        }
    }
    for n>=0 {
        nums1[cur]=nums2[n]
        cur,n = cur-1,n-1
    }
}