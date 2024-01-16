package aoc.utils

import java.net.URL

class FileLoader(private val fileName: String) {
    fun readFile(): String {
        val resource: URL =
            FileLoader::class.java.getResource("/$fileName")
                ?: throw RuntimeException("Failed to read file $fileName")

        return resource.readText()
    }
}
