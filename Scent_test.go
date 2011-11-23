package main

import "testing"

func TestScentCalculateRuntime(t *testing.T) {
    rows = 200
    cols = 160
    terrain := new(Terrain)
    terrain.SeeHill(Point{180, 50}, Player(1))
    terrain.SeeHill(Point{180, 100}, Player(1))
    terrain.SeeWater(Point{179, 50})
    terrain.SeeWater(Point{179, 51})
    terrain.SeeWater(Point{179, 52})
    terrain.SeeWater(Point{179, 53})
    terrain.SeeWater(Point{179, 54})
    terrain.SeeWater(Point{179, 55})
    terrain.SeeWater(Point{179, 56})
    terrain.SeeWater(Point{179, 57})
    terrain.SeeWater(Point{179, 58})
    terrain.SeeWater(Point{179, 59})
    ForEachPoint(func(p Point) {
        if !terrain.At(p).IsVisible() {
            terrain.SeeLand(p)
        }
    })

    mystery := NewMystery(terrain)
    scent := NewScent(terrain, mystery)
    start := now()
    scent.Calculate()
    runtime := now() - start
    if runtime > 100 {
        t.Error(scent)
        t.Errorf("runtime=%v ms\n", runtime)
    }
}

func TestCalculateScentDissipation(t *testing.T) {
    terrain := NewTerrain(
    ".......................................\n" +
    ".......................................\n" +
    ".......................................\n" +
    ".......................................\n" +
    "................%%%%%..................\n" +
    "..................1....................\n" +
    ".......................................\n" +
    ".......................................\n" +
    ".......................................\n" +
    ".......................................")

    mystery := NewMystery(terrain)
    scent := NewScent(terrain, mystery)

    scent.Calculate()
/*
    t.Error(scent)
    scent.Calculate()
    t.Error(scent)
    scent.Calculate()
    t.Error(scent)
    scent.Calculate()
    t.Error(scent)
*/
}

func TestCalculate3(t *testing.T) {
    terrain := NewTerrain(
    "%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%\n" +
    "%1....................................%\n" +
    "%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%.%\n" +
    "%.....................................%\n" +
    "%.%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%\n" +
    "%.....................................%\n" +
    "%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%.%\n" +
    "%.....................................%\n" +
    "%.%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%\n" +
    "%.....................................%\n" +
    "%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%.%\n" +
    "%.....................................%\n" +
    "%.%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%\n" +
    "%.....................................%\n" +
    "%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%.%\n" +
    "%.....................................%\n" +
    "%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")

    mystery := NewMystery(terrain)
    scent := NewScent(terrain, mystery)

    scent.Calculate()
/*
    t.Errorf("%v %v %v %v %v", scent.At(Point{1, 1}), scent.At(Point{5, 1}), scent.At(Point{9, 1}), scent.At(Point{13, 1}), scent.At(Point{17, 1}))
    scent.Calculate()
    t.Errorf("%v %v %v %v %v", scent.At(Point{1, 1}), scent.At(Point{5, 1}), scent.At(Point{9, 1}), scent.At(Point{13, 1}), scent.At(Point{17, 1}))
    scent.Calculate()
    t.Errorf("%v %v %v %v %v", scent.At(Point{1, 1}), scent.At(Point{5, 1}), scent.At(Point{9, 1}), scent.At(Point{13, 1}), scent.At(Point{17, 1}))
    scent.Calculate()
    t.Errorf("%v %v %v %v %v", scent.At(Point{1, 1}), scent.At(Point{5, 1}), scent.At(Point{9, 1}), scent.At(Point{13, 1}), scent.At(Point{17, 1}))
    scent.Calculate()
    t.Errorf("%v %v %v %v %v", scent.At(Point{1, 1}), scent.At(Point{5, 1}), scent.At(Point{9, 1}), scent.At(Point{13, 1}), scent.At(Point{17, 1}))

    terrain.squares[1][1] = terrain.squares[1][1].MinusHill()
    t.Errorf("*disappear*")

    scent.Calculate()
    t.Errorf("%v %v %v %v %v", scent.At(Point{1, 1}), scent.At(Point{5, 1}), scent.At(Point{9, 1}), scent.At(Point{13, 1}), scent.At(Point{17, 1}))
    scent.Calculate()
    t.Errorf("%v %v %v %v %v", scent.At(Point{1, 1}), scent.At(Point{5, 1}), scent.At(Point{9, 1}), scent.At(Point{13, 1}), scent.At(Point{17, 1}))
    scent.Calculate()
    t.Errorf("%v %v %v %v %v", scent.At(Point{1, 1}), scent.At(Point{5, 1}), scent.At(Point{9, 1}), scent.At(Point{13, 1}), scent.At(Point{17, 1}))
    scent.Calculate()
    t.Errorf("%v %v %v %v %v", scent.At(Point{1, 1}), scent.At(Point{5, 1}), scent.At(Point{9, 1}), scent.At(Point{13, 1}), scent.At(Point{17, 1}))
    scent.Calculate()
    t.Errorf("%v %v %v %v %v", scent.At(Point{1, 1}), scent.At(Point{5, 1}), scent.At(Point{9, 1}), scent.At(Point{13, 1}), scent.At(Point{17, 1}))
*/
}
