package aoc.utils

import kotlin.math.abs

class Point(var x: Int, var y: Int) {
    fun manhattanDistance(other: Point): Int {
        return abs(this.x - other.x) + abs(this.y - other.y)
    }

    override fun toString(): String {
        return "($x, $y)"
    }

    override fun equals(other: Any?): Boolean {
        if (other !is Point) {
            return false
        }
        return this.x == other.x && this.y == other.y
    }

    override fun hashCode(): Int {
        var result = 17
        result = 31 * result + x
        result = 31 * result + y
        return result
    }
}
