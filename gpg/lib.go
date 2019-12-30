package gpg
type GenPoly struct{
	Poly map[int]int//Expo Coe
}
func (g *GenPoly) Init(p map[int]int) {
	g.Poly=p
}

func (g *GenPoly) Sort() {

	keys :=g.GetKeys()

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
	g.Sort()
	return g.Poly
}

func (g *GenPoly) MultiplyCoesBy(coe int){
	keys:=g.GetKeys()
	for i:=0;i<len(keys);i++{
		g.Poly[keys[i]]=g.Poly[keys[i]]*coe
	}
} 

func (g *GenPoly) SumExposBy(expo int){
	
	keys:=g.GetKeys()
	p:=make(map[int]int)
	for i:=0;i<len(keys);i++{
		p[(keys[i]+expo)%255]=g.Poly[keys[i]]
	}
	g.Poly=p
} 

