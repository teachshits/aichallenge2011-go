package main

//import "fmt"

const POINTQUEUE_CAPACITY = MAX_ROWS * MAX_COLS

type PointQueue struct {
    points [POINTQUEUE_CAPACITY]Point
    start, end int
}

func (this *PointQueue) Push(p Point) {
    this.points[this.end] = p
    this.end += 1
    this.end %= POINTQUEUE_CAPACITY
    if this.end == this.start {
        panic("PointQueue ran out of capacity")
    }
}

func (this *PointQueue) ForEach(f func(Point)) {
    for this.start != this.end {
        f(this.points[this.start])
        if this.start == this.end {
            break
        }
        this.start += 1
        this.start %= POINTQUEUE_CAPACITY
    }
}

func (this *PointQueue) Clear() {
    this.start = this.end
}

func (this *PointQueue) Empty() bool {
    return (this.end == this.start)
}

func (this *PointQueue) Size() int {
    return (this.end - this.start + POINTQUEUE_CAPACITY) % POINTQUEUE_CAPACITY
}
