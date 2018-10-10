package main

import "fmt"

func main()  {
	//s := "U2FsdGVkX19Q3I//VCH0U3cVtITZ3ckILJnUcdPX3Gs5qjdFlUjZ3mAftGivtFYDN5ZCSkBynnVqBawl4p8wKO008zI6D0A1+VEVCUyEvEeNoUfGcS0El9d93vsPxbg7D5avufQsScgsk3QEtq9Do33AXmDzuuZ0qmZaI7re16FcXIrmPPiQDOHRc7wt0ng6qLiNz7VqESRTdxPOahKFRkWT8sT+Ur2y+2iZ2LEaxNM7UZqcPwYgm6FoK0Vjnqdeg30R27jc6AoFPyRZ2g8+EJMp3n/pf94oSCLEWkc0osjH9DqbM6DUptu3HJbAVwXQ=="
	s := "U2FsdGVkX19Q3I//VCH0U3cVtlTZ3ckILJnUcdPX3Gs5qjd1UjZ3mAftGivtDN5ZCSkBynnVqBawl4p8wKO0O8zI6D0A1+VEVCUyEvEeNoUfGcS0El9d93vb7D5avufQsScgsk3QEtq9/M4Do32OKFeq0/3NrxWOsMmh3AXmDzuuZ0qmZal7re16FcXIrmPPiQDOHRc7wtOng6qLiNz7VqESRTdxPOahKFRkWT8sT+Ur2y+2iZ2LEaNM7UZqcPwYgm6FoKOVjnqdeg3OR27jc6AoFPyRZ2g8+EJMp3n/pf94oSCLEWkc0osjH9DqbM6DUptu3HJbAVwXQ=="
	fmt.Println(len(s))
	r := ""
	for i,a := range s{
		//fmt.Println(string(a))
		if i%25 == 0{
			r = fmt.Sprintf("%s%s",r,string(a))
		}
	}
	fmt.Println(r)
}
