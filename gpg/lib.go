package gpg

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type GenPoly struct {
	Poly         map[int]int           //Expo of X : Expo of α
	GroupBlock   map[string][][]string //string of 8 bit
	GroupBlockEC map[string][][]string //string of 8 bit
}

func New(p map[int]int) GenPoly {

	g := GenPoly{Poly: p}
	return g
}
func NewAntiLog(p map[int]int) GenPoly {

	g := GenPoly{Poly: p}
	return g
}
func (g *GenPoly) Init(p map[int]int) GenPoly {
	g.Poly = p
	return *g
}
func (g *GenPoly) ToAntilog() {
	keys := g.GetKeys()
	//p := make(map[int]int)
	for i := 0; i < len(keys); i++ {
		g.Poly[keys[i]] = AntiLog[g.Poly[keys[i]]]
	}
}
func (g *GenPoly) ToLog() {
	keys := g.GetKeys()
	//p := make(map[int]int)
	for i := 0; i < len(keys); i++ {
		g.Poly[keys[i]] = Log[g.Poly[keys[i]]]
	}
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
func (g *GenPoly) Serialize(n int) string {
	s := ""
	if n == 1 {
		//no need to interleaving

		for grp := range g.GroupBlock {
			// g.GroupBlock[grp]
			for _,block:= range g.GroupBlock[grp] {
				s=s+strings.Join(block[:], "")
			}
		}

		for grp := range g.GroupBlockEC {
			// g.GroupBlock[grp]
			for _,block := range g.GroupBlockEC[grp] {
				s=s+strings.Join(block[:], "")
			}
		}
	} else {
		len1:=len(g.GroupBlock["GROUP1"])
		len2:=len(g.GroupBlock["GROUP2"])
		if(len1!=0){
			len1=len(g.GroupBlock["GROUP1"][0])
		}
		if(len2!=0){
			len2=len(g.GroupBlock["GROUP2"][0])
		}

		lenMax:=Max(len1,len2)

		for i := 0; i < lenMax; i++ {

			for grp := range g.GroupBlock {
				for _,block := range g.GroupBlock[grp] {
					if  i<len(block) {
						s=s+block[i]
					}
				}
			}
		}
	}
	return s
}
func Max(a int,b int) int{
	if(a>b){
		return a
	}
	return b
}
func (g *GenPoly) SetGroupBlockECC(s string, BlocksInfo map[string][]int, CountECC int, Galois int) {

	g.GroupBlockEC = make(map[string][][]string, 2)
	begin := 0

	for Grp := range BlocksInfo {
		//fmt.Println(Grp)
		var BlocksCount = BlocksInfo[Grp][0]
		var CodeWordCounts = BlocksInfo[Grp][1]

		intArryas := make([][]int, BlocksCount)
		eccArryas := make([][]string, BlocksCount)
		//calc in every data codeword in block
		for i := 0; i < len(intArryas); i++ {
			intArryas[i] = make([]int, CodeWordCounts)
			for j := 0; j < len(intArryas[i]); j++ {
				v, _ := strconv.ParseInt(s[begin:begin+8], 2, 16)
				intArryas[i][j] = int(v)
				begin += 8
			}
			p := map[int]int{}
			for j := len(intArryas[i]) - 1; j >= 0; j-- {
				p[len(intArryas[i])-1-j] = intArryas[i][j]
			}
			gpg1 := New(p)
			gpg2 := GenGalois(Galois)
			div := gpg1.Divide(
				gpg2,
			)
			//div.Sort()

			eccArryas[i] = make([]string, len(div.Poly))

			keysD := div.GetSortedKeys()
			ecc := []string{}
			for k := range keysD {
				ecc = append(ecc, fmt.Sprintf("%08b", div.Poly[len(keysD)-1-k]))
			}
			eccArryas[i] = ecc

		}
		//fmt.Println(intArryas)
		g.GroupBlockEC[Grp] = eccArryas
	}
}
func (g *GenPoly) SetGroupBlock(s string, BlocksInfo map[string][]int) {

	g.GroupBlock = make(map[string][][]string, 2)
	begin := 0

	for Grp := range BlocksInfo {
		//fmt.Println(Grp)
		var BlocksCount = BlocksInfo[Grp][0]
		var CodeWordCounts = BlocksInfo[Grp][1]

		stringArrays := make([][]string, BlocksCount)
		for i := 0; i < len(stringArrays); i++ {
			stringArrays[i] = make([]string, CodeWordCounts)
			for j := 0; j < len(stringArrays[i]); j++ {
				stringArrays[i][j] = s[begin : begin+8]
				begin += 8
			}
		}
		//fmt.Println(stringArrays)
		g.GroupBlock[Grp] = stringArrays
	}

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
func (g *GenPoly) SumCeosBy(ceo int) {
	keys := g.GetKeys()
	for i := 0; i < len(keys); i++ {
		g.Poly[keys[i]] = (g.Poly[keys[i]] + ceo) % 255
	}
	//fmt.Println("::",g.Poly)
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
func (g *GenPoly) GetSortedKeys() []int {
	keysG := []int{}
	for key, _ := range g.Poly {
		keysG = append(keysG, key)
	}
	sort.Slice(keysG, func(i, j int) bool {
		return keysG[i] > keysG[j]
	})
	return keysG
}
func (gg GenPoly) Divide(ff GenPoly) GenPoly {

	f := ff
	g := gg
	keysG := g.GetSortedKeys()
	NTimes := len(keysG)
	keysF := f.GetSortedKeys()

	if g.Poly[keysG[0]] != f.Poly[keysF[0]] {
		g.SumExposBy(keysF[0])
		f.SumExposBy(keysG[0])
	}
	//now they have same expo
	//panic("stop")
	for n := 0; n < NTimes; n++ {

		keysG = g.GetSortedKeys()
		keysF = f.GetSortedKeys()

		f.SumCeosBy(AntiLog[g.Poly[keysG[0]]])
		f.ToLog()

		for i := 0; i < len(keysG); i++ {
			k := keysG[i]
			if fval, ok := f.Poly[k]; ok {
				g.Poly[k] = (g.Poly[k] ^ fval) & 255
				if i == 0 {
					delete(g.Poly, k)
					delete(f.Poly, k)
				}
			} else {
				g.Poly[k] = (g.Poly[k] ^ 0) & 255
			}
		}
		keysF = f.GetSortedKeys()
		for i := 0; i < len(keysF); i++ {
			if _, ok := g.Poly[keysF[i]]; !ok {
				g.Poly[keysF[i]] = f.Poly[keysF[i]]
			}
		}

		f = ff
		keysF = f.GetSortedKeys()
		keysG = g.GetSortedKeys()
		f.SumExposBy(keysG[0] - keysF[0])
	}

	return g
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

func GenGalois(n int) GenPoly {

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
		gr := GenGalois(n - 1)
		gg = gr.MultiplyBy(gg)
		return gg
	}
}
