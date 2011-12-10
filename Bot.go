package main

import "fmt"
import "os"

type Bot struct {
    terrain, update *Terrain
    mystery *Mystery
    potentialEnemy *PotentialEnemy
    search *Search
    predictions *Predictions
    //reinforcement *Reinforcement
    command *Command
    hud *os.File
    hudCenter Point
}

func (this *Bot) Ready() {
    VerifySituationSize()

    this.terrain = new(Terrain)
    this.mystery = NewMystery(this.terrain)
    this.potentialEnemy = NewPotentialEnemy(this.terrain)
    this.search = NewSearch(this.terrain, this.mystery, this.potentialEnemy)
    this.predictions = NewPredictions(this.terrain)
    //this.reinforcement = NewReinforcement(this.terrain, this.army, this.distanceToTrouble)
    this.command = NewCommand(this.terrain, this.search, this.predictions)

    if debug {
        this.hud = NewLog("hud", "txt")
    }
}

func (this *Bot) Turn() {
    this.update = new(Terrain)
}

func (this *Bot) SeeWater(row, col int) {
    this.update.SeeWater(Point{row, col})
}

func (this *Bot) SeeFood(row, col int) {
    this.update.SeeFood(Point{row, col})
}

func (this *Bot) SeeAnt(row, col, player int) {
    this.update.SeeAnt(Point{row, col}, Player(player))
}

func (this *Bot) SeeHill(row, col, player int) {
    this.update.SeeHill(Point{row, col}, Player(player))
}

func (this *Bot) SeeDeadAnt(row, col, player int) {
}

func (this *Bot) Go(issueOrder func(int, int, byte), done func()) {
    startTime := now()

    this.terrain.Update(this.update)

    this.mystery.Calculate()
    this.potentialEnemy.Calculate()
    this.search.Calculate()
    this.predictions.Calculate()
    //this.reinforcement.Calculate()
    this.command.Calculate()

    this.command.ForEach(func(move Move) {
        issueOrder(move.from.row, move.from.col, move.dir.Char())
    })

    time := now() - startTime

    if debug {
        this.hud.WriteString(fmt.Sprintf("\n%v\n", this.ColorString()))
        this.hud.
        WriteString(fmt.Sprintf("turn %4v, time %3v (map %3v, myst %3v, potE %3v, search %3v, pred %3v, re %3v, comm %3v)",
                turn, time, this.terrain.time, this.mystery.time, this.potentialEnemy.time, this.search.time, this.predictions.time, 999, this.command.time))
        this.hud.Sync()

        //NewTurnLog("map", "txt").WriteString(this.terrain.String())
        //NewTurnLog("mystery", "txt").WriteString(this.mystery.String())
        //NewTurnLog("potentialEnemy", "txt").WriteString(this.potentialEnemy.String())
        //NewTurnLog("forageScent", "txt").WriteString(this.forageScent.String())
        //NewTurnLog("forageScent", "csv").WriteString(this.forageScent.Csv())
        //NewTurnLog("battleScent", "txt").WriteString(this.battleScent.String())
        //NewTurnLog("battleScent", "csv").WriteString(this.battleScent.Csv())
        //NewTurnLog("army", "txt").WriteString(this.army.String())
        //NewTurnLog("distanceToTrouble", "txt").WriteString(this.distanceToTrouble.String())
    }

    done()
}

func (this *Bot) End() {
    if debug {
        this.hud.WriteString(fmt.Sprintf("\nGame over.\n"))
    }
}

func (this *Bot) ColorString() string {
    ForEachPoint(func(p Point) {
        if this.terrain.At(p).HasFriendlyHill() {
            this.hudCenter = p
        }
    })

    viewportSize := Point{63, 170}
    topLeftCorner := this.hudCenter.Plus(Point{-viewportSize.row / 2, -viewportSize.col / 2})
    if rows < viewportSize.row {
        topLeftCorner.row -= (rows - viewportSize.row) / 2
    }
    if cols < viewportSize.col {
        topLeftCorner.col -= (cols - viewportSize.col) / 2
    }

    return GridToColorString(viewportSize, func(p1 Point) ColorChar {
        p := p1.Plus(topLeftCorner)
        s := this.terrain.At(p)

        var cc ColorChar

        cc.background = BLACK
        if !s.IsVisible() {
            cc.background += BRIGHT
        }

        switch {
        case s.HasLand():
            switch {
            case s.HasFood():
                cc.symbol = "*"
                cc.foreground = YELLOW
            case s.HasHill():
                if s.HasAnt() {
                    cc.symbol = string('a' + byte(s.owner))
                } else {
                    cc.symbol = " "
                }
                cc.foreground = BLACK
                if s.IsFriendly() {
                    cc.background = BRIGHT + GREEN
                } else {
                    cc.background = BRIGHT + RED
                }
            case s.HasAnt():
                cc.symbol = string('a' + byte(s.owner))
                if s.IsFriendly() {
                    //switch {
                    //case this.tier.At(p) == 1:
                        cc.foreground = GREEN
                    //case this.tier.At(p) == 2:
                    //    cc.foreground = CYAN
                    //case this.tier.At(p) == 3:
                    //    cc.foreground = BRIGHT + CYAN
                    //case this.tier.At(p) >= 4:
                    //    cc.foreground = MAGENTA
                    //default:
                    //    cc.foreground = WHITE
                    //}
                    //if this.reinforcement.At(p) {
                    //    cc.foreground += BRIGHT
                    //}
                } else {
                    cc.foreground = RED
                }
            default:
                cc.symbol = "‧"
                if this.potentialEnemy.At(p) {
                    cc.foreground = RED
                } else {
                    cc.foreground = BRIGHT + BLACK
                }
            }
        case s.HasWater():
            cc.symbol = "▒"
            cc.foreground = BLUE
        default:
            cc.symbol = " "
            cc.background = WHITE
        }
        return cc
    })
}
