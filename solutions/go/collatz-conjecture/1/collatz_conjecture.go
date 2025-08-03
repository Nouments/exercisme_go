package collatzconjecture
import "errors"

func CollatzConjecture(n int) (int, error) {
    if n<=0{
        return 0,errors.New("erreur")
    }
    stp:=0
	for {
        if n==1{
            break
        }
        if n%2==0{
            n= n/2
        }else{
            n=(3*n)+1
        }
        stp++
    }
    return stp,nil
}
