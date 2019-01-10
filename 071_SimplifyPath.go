func simplifyPath(path string) string {
    strArr := strings.Split(path,"/")
    
    var strStack []string
    for _,str := range strArr {
        if str == ".." {
            if len(strStack) > 0 {
                strStack = strStack[:len(strStack)-1]
            }
        } else if str!="." && str != ""{
            strStack = append(strStack,str)
        }
    }  
    return "/" + strings.Join(strStack,"/")
}