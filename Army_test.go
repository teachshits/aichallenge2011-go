package main

import "testing"

func TestArmy(t *testing.T) {
    terrain := NewTerrain(
        "..aa........*.............\n" +
        "...aa..............a......\n" +
        "...........a..........a...\n" +
        "....abb.....a.........a...\n" +
        ".............a...aa...a...\n" +
        "..........................\n" +
        "....a..a.......aaaa.......\n" +
        "....aa.a..........a.......\n" +
        "................a.a.......\n" +
        "..................a.......\n" +
        "..........................")
    expected :=
        "..AA......................\n" +
        "...AA..............e......\n" +
        "...........F..........G...\n" +
        "....h.......F.........G...\n" +
        ".............F...ll...G...\n" +
        "..........................\n" +
        "....O..p.......QQQQ.......\n" +
        "....OO.p..........Q.......\n" +
        "................y.Q.......\n" +
        "..................Q.......\n" +
        ".........................."
    army := NewArmy(terrain)
    army.Calculate()
    if army.String() != expected {
        t.Error(army)
    }
}
