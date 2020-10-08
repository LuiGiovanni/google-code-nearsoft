package main
​
import( 
	"bufio"
	. "fmt" 
	. "os"
)
​
var B int
​
var reader *bufio.Reader = bufio.NewReader(Stdin)
var writer *bufio.Writer = bufio.NewWriter(Stdout)

func query(i int) (int){
	Println(i)
	defer writer.Flush()
	var r int
	Scanf("%d", &r)
	return r
}

func easb_atad(){
	var result = make([]int,B+1)
	var L = 1
	var R = B
	for nr = 1; true; nr += 2{
		if nr % 10 == 1 && nr != 1{
			var found = -1
			var found_diff = -1
			for i = 1; i < L; i++{
				if result[i] == result[B+1-i]{
					found = i;
				}else{
					found_diff = i
				}
			}
			if found == -1{
				new_value = query(1)
				query(1)
				if new_value != result[1] {
					for i = 1; i <= L; i++ {
						result[i] ^= 1
					}
					for i = R; i <= B; i++ {
						result[i] ^= 1
					}
				}
			}else{
				one = query(found)
				if one != result[found] {
					for i = 1; i <= L; i++ {
						result[i] ^= 1
					}
					for i = R; i <= B; i++ {
						result[i] ^= 1
					}
				}
				if found_diff == -1 {
					query(found)
				}else{
					if query(found_diff) != result[found_diff] {
						for i, j = 1, len(result)-1; i < j; i, j = i+1, j-1 {
							result[i], result[j] = result[j], result[i]
						}
					}
				}
			}
			nr += 2
		}
		result[L] = query(L)
		result[R] = query(R)
		L++
		R--
		if L > R {
			for i = 1; i <= B; i++ {
				Printf("%d",result[i])
			}
			Println()
			defer writer.Flush()
			var response string
			Scanf("%s", &response)
			if response != "Y"{
				Exit(0)
			}
			return
		}
	}
}
