package main

type Terrain struct {
    squares [MAX_ROWS][MAX_COLS]Square
    waterNeighbors [MAX_ROWS][MAX_COLS]byte
}

func NewTerrain(input string) *Terrain {
    this := new(Terrain)

    rows = 1
    cols = 1
    p := Point{0, 0}

    for _, c := range input {
        switch {
        case c == '.':
            this.SeeLand(p)
        case c == '%':
            this.SeeWater(p)
        case c == '*':
            this.SeeFood(p)
        case c >= 'a' && c <= 'z':
            owner := Player(c - 'a')
            this.SeeAnt(p, owner)
        case c >= '0' && c <= '9':
            owner := Player(c - '0')
            this.SeeHill(p, owner)
        case c >= 'A' && c <= 'Z':
            owner := Player(c - 'A')
            this.SeeAnt(p, owner)
            this.SeeHill(p, owner)
        case c == '\n':
            p.row += 1
            p.col = 0
            continue
        }

        if rows <= p.row {
            rows = p.row + 1
            if rows >= MAX_ROWS {
                panic("Too many rows")
            }
        }
        if cols <= p.col {
            cols = p.col + 1
            if cols >= MAX_COLS {
                panic("Too many cols")
            }
        }
        p.col += 1
    }

    return this
}

func (this *Terrain) At(p Point) Square {
    return this.squares[p.row][p.col]
}

func (this *Terrain) SeeWater(p Point) {
    this.squares[p.row][p.col] = this.At(p).PlusVisible().PlusWater()
    ForEachNeighbor(p, func(p2 Point) {
        this.waterNeighbors[p2.row][p2.col] += 1
    })
}

func (this *Terrain) SeeLand(p Point) {
    this.squares[p.row][p.col] = this.At(p).PlusVisible().PlusLand()
}

func (this *Terrain) SeeFood(p Point) {
    this.squares[p.row][p.col] = this.At(p).PlusVisible().PlusLand().PlusFood()
}

func (this *Terrain) SeeAnt(p Point, owner Player) {
    this.squares[p.row][p.col] = this.At(p).PlusVisible().PlusLand().PlusAnt(owner)
}

func (this *Terrain) SeeHill(p Point, owner Player) {
    this.squares[p.row][p.col] = this.At(p).PlusVisible().PlusLand().PlusHill(owner)
}

func (this *Terrain) Update(terrain *Terrain) {
    ForEachPoint(func(p Point) {
        s := this.At(p).MinusVisible().MinusAnt()
        s2 := terrain.At(p)
        if s2.HasAnt() {
            s = s.PlusAnt(s2.owner)
        }
        this.squares[p.row][p.col] = s
    })

    ForEachPoint(func(p Point) {
        if this.At(p).HasFriendlyAnt() {
            ForEachPointWithinRadius2(p, viewradius2, func(p2 Point) {
                this.squares[p2.row][p2.col] = this.At(p2).PlusVisible()
            })
        }
    })

    ForEachPoint(func(p Point) {
        s := this.At(p)

        if s.IsVisible() {
            s2 := terrain.At(p)

            if s2.HasWater() {
                //s = s.PlusWater()
                this.SeeWater(p)
                return
            } else if !s.HasWater() && !s.HasLand() {
                s = s.PlusLand()
            }

            if s2.HasFood() {
                s = s.PlusFood()
            } else if s.HasFood() {
                s = s.MinusFood()
            }

            if s2.HasHill() {
                s = s.PlusHill(s2.owner)
            } else if s.HasHill() {
                s = s.MinusHill()
            }

            this.squares[p.row][p.col] = s
        }
    })
}

func (this *Terrain) String() string {
    return GridToString(func(p Point) byte {
        s := this.At(p)
        switch {
        case s.HasLand():
            switch {
            case s.HasFood():
                return '*'
            case s.HasAnt() && s.HasHill():
                return 'A' + byte(s.owner)
            case s.HasAnt():
                return 'a' + byte(s.owner)
            case s.HasHill():
                return '0' + byte(s.owner)
            }
            return '.'
        case s.HasWater():
            return '%'
        }
        return '?'
    })
}