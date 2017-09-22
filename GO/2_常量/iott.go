package main;
import(
	"fmt"
	
)
func main(){
const(
	a=iota //0
	b 	   //1
	c	   //2
	d="ha" //独立值 iota+=1
	e      //"ha" iota+=1
	f=100  //iota +=1
	g      //100 iota +=1
	h = iota //恢复计数
	i		//	8
	j = 2	//独立值 2
	k = 3	//独立值 2
	l = iota//恢复计数
	m		//iota+1
)
fmt.Println(a,b,c,d,e,f,g,h,i,j,k,l,m)
}