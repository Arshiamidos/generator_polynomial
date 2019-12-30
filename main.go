package main
import (
	"fmt"
)

type GenPoly struct{
	Poly map[int]int//Expo Coe
}
func (g *GenPoly) Init(p map[int]int) {
	g.Poly=p
}

func (g *GenPoly) Sort() {

	keys := []int{}
	for k := range g.Poly {
		keys = append(keys, k)
	}

	for i:=0;i<len(keys);i++{
		for j:=i+1;j<len(keys);j++{
			if keys[i] > keys[j] {
				tmp:=keys[i]
				keys[i]=keys[j]
				keys[j]=tmp
			}
		}
	}

	p:=make(map[int]int)
	for i:=0;i<len(keys);i++{
		p[keys[i]]=g.Poly[keys[i]]
	}
	g.Poly=p

}

func (g *GenPoly) GetKeys() []int{

	keys := []int{}
	for k := range g.Poly {
		keys = append(keys, k)
	}
	return keys
}

func (g *GenPoly) GetAll() map[int]int{
	return g.Poly
}


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


	fmt.Println(CoeExpos)
}