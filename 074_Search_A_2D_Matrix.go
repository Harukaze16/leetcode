//o(m+n)
func searchMatrix(matrix [][]int, target int) bool {
   	if len(matrix)==0 || len(matrix[0])==0 {
   		return false
   	} 	

   	rows,cols := len(matrix),len(matrix[0])
   	i,j:=0,cols-1
   	for i<rows && j>=0 {
   		if matrix[i][j]==target {
   			return true
   		} else if matrix[i][j]<target {
   			i++
   		} else {
   			j--
   		}
   	}
   	return false
}

//O(lg(mn))
func searchMatrix(matrix [][]int, target int) bool {
      if len(matrix)==0 || len(matrix[0])==0 {
         return false
      }  

      rows,cols := len(matrix),len(matrix[0])

   i,j := 0,rows*cols
   for i<j {
        mid := (i+j)/2
      if matrix[mid/cols][mid%cols] == target {
         return true
      } else if matrix[mid/cols][mid%cols] <target {
         i=mid+1
      } else {
         j=mid
      }
   }
   return false
}