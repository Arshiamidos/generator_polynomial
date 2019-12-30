package main
import (
	"fmt"
)

type CoeExpo struct{
	Coe  int;
	Expo int;
}

func main(){
	fmt.Println("ya");
	CoeExpos := []CoeExpo{
		CoeExpo{3,4},
		CoeExpo{2,1},
		CoeExpo{1,2},
		CoeExpo{10,5},
		CoeExpo{6,3},
		CoeExpo{7,2},
		//3x^4+ 2x^1 +2x^2
	}

	for i:=0;i<len(CoeExpos);i++{
		for j:=i+1;j<len(CoeExpos);j++{
			if CoeExpos[i].Expo > CoeExpos[j].Expo {
				tmp:=CoeExpos[i]
				CoeExpos[i]=CoeExpos[j]
				CoeExpos[j]=tmp
			}
		}
	}
	for i:=0;i<len(CoeExpos);i++{
		if(i!=0){
			if CoeExpos[i-1].Expo==CoeExpos[i].Expo {

			}
		}
	}





	fmt.Println(CoeExpos)
}