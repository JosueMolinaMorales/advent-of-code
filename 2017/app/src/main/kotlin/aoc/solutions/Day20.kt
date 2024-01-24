package aoc.solutions

import aoc.utils.FileLoader
import kotlin.math.abs

fun solveDayTwenty() {
    val input = FileLoader("day20.txt").readLines()
    println("Day 20 part 1: ${partOne(input)}")
    println("Day 20 part 2: ${partTwo(input)}")
}

private fun partOne(input: List<String>): Int {
    val particles = input.map { Particle.fromLine(it) }
    var minAcc = Int.MAX_VALUE
    var minAccIndex = -1
    particles.forEachIndexed { index, particle ->
        val acc = particle.acceleration.manhattanDistance()
        if (acc < minAcc) {
            minAcc = acc
            minAccIndex = index
        }
    }
    return minAccIndex
}

private fun partTwo(input: List<String>): Int {
    val particles = input.map { Particle.fromLine(it) }.toMutableList()
    var particlesLeft = particles.size
    var lastSize = particlesLeft
    var attempts = 100
    while (true) {
        val positions = mutableMapOf<String, MutableList<Int>>()
        particles.forEachIndexed { idx, particle ->
            particle.tick()
            val key = "${particle.position.x},${particle.position.y},${particle.position.z}"
            positions.getOrPut(key) { mutableListOf() }.add(idx)
        }
        val toRemove = mutableListOf<Int>()
        positions.forEach { (_, value) ->
            if (value.size > 1) {
                toRemove.addAll(value)
            }
        }
        particlesLeft -= toRemove.size
        particles.removeAll { toRemove.contains(particles.indexOf(it)) }
        if (particlesLeft < lastSize) {
            attempts = 100
            lastSize = particlesLeft
        } else if (attempts == 0) {
            return particlesLeft
        } else {
            attempts -= 1
        }
    }
}

private data class Point(
    var x: Int,
    var y: Int,
    var z: Int,
) {
    fun manhattanDistance(): Int {
        return abs(x) + abs(y) + abs(z)
    }

    companion object {
        fun fromLine(line: String): Point {
            val parts = line.split(",")

            return Point(parts[0].toInt(), parts[1].toInt(), parts[2].toInt())
        }
    }
}

private data class Particle(
    val position: Point,
    val velocity: Point,
    val acceleration: Point,
) {
    companion object {
        fun fromLine(line: String): Particle {
            val parts = line.split(", ")
            val position = Point.fromLine(parts[0].substring(3, parts[0].length - 1))
            val velocity = Point.fromLine(parts[1].substring(3, parts[1].length - 1))
            val acceleration = Point.fromLine(parts[2].substring(3, parts[2].length - 1))
            return Particle(position, velocity, acceleration)
        }
    }

    fun tick() {
        velocity.x += acceleration.x
        velocity.y += acceleration.y
        velocity.z += acceleration.z
        position.x += velocity.x
        position.y += velocity.y
        position.z += velocity.z
    }
}
