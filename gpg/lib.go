package gpg

type GenPoly struct {
	Poly map[int]int //Expo of X : Expo of α
}

func New(p map[int]int) GenPoly {

	g := GenPoly{Poly: p}
	return g

}
func (g *GenPoly) Init(p map[int]int) GenPoly {
	g.Poly = p
	return *g
}
func (g *GenPoly) InitByAntiLog(p map[int]int) GenPoly {

	for k := range p {
		if p[k] == 0 {
			panic("Zero value if not valid for antilog !!")
		}
		p[k] = AntiLog[p[k]]
	}

	g.Poly = p
	return *g
}

func (g *GenPoly) Sort() {

	keys := g.GetKeys()

	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			if keys[i] > keys[j] {
				tmp := keys[i]
				keys[i] = keys[j]
				keys[j] = tmp
			}
		}
	}

	p := make(map[int]int)
	for i := 0; i < len(keys); i++ {
		p[keys[i]] = g.Poly[keys[i]]
	}
	g.Poly = p
}

func (g *GenPoly) GetKeys() []int {

	keys := []int{}
	for k := range g.Poly {
		keys = append(keys, k)
	}
	return keys
}

func (g *GenPoly) GetAll() map[int]int {
	g.Sort()
	return g.Poly
}

func (g *GenPoly) MultiplyCoesBy(coe int) {
	keys := g.GetKeys()
	for i := 0; i < len(keys); i++ {
		g.Poly[keys[i]] = g.Poly[keys[i]] * coe
	}
}

func (g *GenPoly) SumExposBy(expo int) {

	keys := g.GetKeys()
	p := make(map[int]int)
	for i := 0; i < len(keys); i++ {
		p[(keys[i]+expo)%255] = g.Poly[keys[i]]
	}
	g.Poly = p
}

func (g *GenPoly) MultiplyBy(f GenPoly) GenPoly {

	keys := g.GetKeys()
	COEs := New(map[int]int{})
	for i := 0; i < len(keys); i++ {
		COEs.Poly[keys[i]] = (f.Poly[0] + g.Poly[keys[i]]) % 255
		//multiply of a
	}
	COEs.Sort()


	XEXPO := New(map[int]int{})
	for i := 0; i < len(keys); i++ {
		XEXPO.Poly[keys[i]+1] = g.Poly[keys[i]]
		//multiply of x
	}
	XEXPO.Sort()


	return XEXPO.SumBy(COEs)
}

func (g *GenPoly) SumBy(f GenPoly) GenPoly {

	h := New(map[int]int{})
	for k := range g.Poly {

		if fval, ok := f.Poly[k]; ok {
			h.Poly[k] = AntiLog[(Log[g.Poly[k]] ^ Log[fval])]
			delete(f.Poly, k)
		} else {
			h.Poly[k] = g.Poly[k]
		}
		delete(g.Poly, k)
	}
	for k := range f.Poly {
		h.Poly[k] = f.Poly[k]
	}

	return h

}

func (g *GenPoly) GenGalois(n int) GenPoly {

	if n == 1 {

		gg := New(
			map[int]int{
				0: 0,
				1: 0,
			},
		)

		return gg

	} else {

		gg := New(
			map[int]int{
				0: n - 1,
				1: 0,
			},
		)
		gr := g.GenGalois(n - 1)
		gg = gr.MultiplyBy(gg)
		return gg
	}
}
