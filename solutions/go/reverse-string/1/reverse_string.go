

package reverse

func Reverse(input string) string {
    a:=[]rune(input)
	rev:=[]rune{}
    for i:=len(a)-1;i>=0;i--{
        rev= append(rev,a[i])
    }
    return string(rev)
}
