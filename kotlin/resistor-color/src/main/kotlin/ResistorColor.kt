object ResistorColor {

    private val colorMap: Map<String, Int> = colors().withIndex().associate { (i,s) -> Pair(s, i) }.toMap()

    fun colorCode(input: String): Int {
        return colorMap.getOrDefault(input, -1)
    }

    fun colors(): List<String> {
        return listOf<String>("black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white")
    }

}
