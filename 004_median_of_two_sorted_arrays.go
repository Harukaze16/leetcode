
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    len1,len2 := len(nums1),len(nums2)
    if (len1+len2)%2 == 0 {
        return ((float64)(find_kth(nums1,nums2,(len1+len2)/2)+find_kth(nums1,nums2,(len1+len2)/2+1)))/2
    } else {
        return float64(find_kth(nums1,nums2,(len1+len2)/2+1))
    }
}

func find_kth(nums1 []int,nums2 []int, k int) int {
	beg1,end1,beg2,end2 := 0,len(nums1)-1,0,len(nums2)-1

	for k>1 && end1>=beg1 && end2>=beg2 {
		findth := k/2
		i1 := min(beg1+findth-1,end1) // donot forget the beg1
		i2 := min(beg2+findth-1,end2)
		if nums1[i1]<nums2[i2] {
            end2 = min(end2,i2+k%2)  //notice 应在k变动前得出end
            k-=(i1+1-beg1)
			beg1 = i1+1

		} else {
			end1 = min(end1,i1+k%2) //注意此处以防止end值溢出
			k-=(i2+1-beg2)
			beg2= i2+1
		}
	}
	if beg1>end1 {
		return nums2[beg2+k-1]
	} else if beg2 >end2 {
		return nums1[beg1+k-1]
	} else {
		return min(nums1[beg1+k-1],nums2[beg2+k-1]) //notice there beg1
	}

}

func min(i1 int,i2 int) int{
	if(i1<=i2) {
		return i1
	}
	return i2
}


