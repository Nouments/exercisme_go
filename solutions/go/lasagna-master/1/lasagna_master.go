package lasagna
import ("strings"
        "fmt")
// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string , averageTime int) int{
    if averageTime==0{
        return len(layers)*2
    }
     return len(layers)*averageTime
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int,float64){
    return (strings.Count(strings.Join(layers, " "),"noodles")*50 ), float64(strings.Count(strings.Join(layers, " "),"sauce"))*0.2
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(a []string , b []string ) {
    if (strings.Contains(strings.Join(a, " "), "?")){
        a[len(a)-1]=b[len(b)-1]
    }else{
        b[len(b)-1]=a[len(a)-1]
    }
    
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, n int) []float64{
   	quantities2:=make([]float64, len(quantities))
    for i:= range quantities{
        quantities2[i]= (quantities[i]*float64(n))/2
    }
    fmt.Println(quantities2)
    
    return quantities2
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
