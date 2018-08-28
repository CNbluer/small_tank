package main
//test github
import (
	"Clib"
	"fmt"
	"os"
	"math/rand"
	"time"
)
//
const Wide int=20
const High int=20
var key =make(chan int)
var t tank
var lasthead position
var lastbody [5]position
var z zidan
var maxin danjia
var newz zidan

type position struct {
	X int
	Y int
}
type danjia struct {
	danjia []position
} 
type tank struct {
	posbody [5]position
	poshead position
}
type zidan struct {
	 position
}
func (z *zidan)sendzidan()zidan  {
	z.X=t.poshead.X
	z.Y=t.poshead.Y-1
	return *z
}

func DrawUI(p position,ch byte)  {
	Clib.GotoPostion(p.X*2+4,p.Y+2)
	fmt.Fprintf(os.Stderr,"%c",ch)
}
func (t *tank)tankinit()  {
	t.poshead.X,t.poshead.Y=Wide/2,High-3
	t.posbody[0].X,t.posbody[0].Y=t.poshead.X-1,t.poshead.Y+1
	t.posbody[1].X,t.posbody[1].Y=t.poshead.X,t.poshead.Y+1
	t.posbody[2].X,t.posbody[2].Y=t.poshead.X+1,t.poshead.Y+1
	t.posbody[3].X,t.posbody[3].Y=t.poshead.X-1,t.poshead.Y+2
	t.posbody[4].X,t.posbody[4].Y=t.poshead.X+1,t.poshead.Y+2
	fmt.Fprintln(os.Stderr,
		`
  #-----------------------------------------#
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  #-----------------------------------------#
`)

	go func() {
		for {
			switch Clib.Direction() {
			case 87, 119:
				//方向上
				lasthead=t.poshead
				lastbody=t.posbody
				DrawUI(lasthead, ' ')
				for i := 0; i < len(lastbody); i++ {
				DrawUI(lastbody[i], ' ')
			}
				t.poshead.Y--
				for i := 0; i < len(t.posbody); i++ {
					t.posbody[i].Y--
				}
				key<-66
				//方向左
			case 65, 97:
				lasthead=t.poshead
				lastbody=t.posbody
				DrawUI(lasthead, ' ')

				for i := 0; i < len(lastbody); i++ {
					DrawUI(lastbody[i], ' ')
				}
				t.poshead.X--
				for i := 0; i < len(t.posbody); i++ {
					t.posbody[i].X--
				}
				key<-66
				//方向右
			case 68, 100:
				lasthead=t.poshead
				lastbody=t.posbody
				DrawUI(lasthead, ' ')
				for i := 0; i < len(lastbody); i++ {
					DrawUI(lastbody[i], ' ')
				}
				t.poshead.X++
				for i := 0; i < len(t.posbody); i++ {
				t.posbody[i].X++
			}
				key<-66
				//方向下
			case 83, 115:
				lasthead=t.poshead
				lastbody=t.posbody
				DrawUI(lasthead, ' ')
				for i := 0; i < len(lastbody); i++ {
					DrawUI(lastbody[i], ' ')
				}
				t.poshead.Y++
				for i := 0; i < len(t.posbody); i++ {
					t.posbody[i].Y++
				}
				key<-66
			}
		}
	}()
	go func() {
		for {
			for i:=0;i<len(maxin.danjia);i++ {
				DrawUI(maxin.danjia[i], '@')
			}

			time.Sleep(time.Millisecond*300)
		}
	}()
	go func() {
		for  {
			<-key
			time.Sleep(time.Millisecond*300)
			newz=z.sendzidan()
			maxin.danjia=append(maxin.danjia,newz.position)
		}
	}()

}
func (t *tank)playgame()  {
	for {
		time.Sleep(time.Millisecond*100)
		DrawUI(t.poshead, '|')
		for i := 0; i < len(t.posbody); i++ {
			DrawUI(t.posbody[i], '*')
		}
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	Clib.HideCursor()
	t.tankinit()
	z.sendzidan()
	t.playgame()
}
