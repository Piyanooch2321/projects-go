package prime
import "fmt"

func Prime() {

	for i:=0; i<=100; i++ {
	
		count := 0 
		for j:=i ; j>0; j--{
			if(i % j) == 0 {
				count ++
			}

		}

		if count == 2{
			fmt.Print(i,",")
		}

		
		// Println("hello ,world" ,i)
	}
	

}