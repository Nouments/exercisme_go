package diffsquares

func SquareOfSum(n int) int {
    sm:=0
    for i:=1;i<=n;i++{
        sm+=i
    }
    return sm*sm
}

func SumOfSquares(n int) int {
    sm:=0
    for i:=1;i<=n;i++{
        sm+= i*i
    }
	// implement the SumOfSquares function")
    return sm
    
}

func Difference(n int) int {
    a:=SquareOfSum(n)
    b:=SumOfSquares(n)
    return int(a-b)
	//panic("Please implement the Difference function")
}
