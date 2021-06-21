import kotlin.math.pow

object ResistorColorTrio {

    fun text(vararg input: Color): String {
        val n = input.take(2).map(Color::ordinal).reduce{acc, e -> acc *10 + e}
        val v = 10.toDouble().pow(input.get(2).ordinal).toInt()*n
        val (base, ohm) = when{
            v / 1000000000 != 0 -> Pair(v/10.pow(9), Unit.GIGAOHMS.name)
            v / 1000000 != 0 -> Pair(v/10.pow(6), Unit.MEGAOHMS.name)
            v / 1000 != 0 -> Pair(v/10.pow(3), Unit.KILOOHMS.name)
            else -> Pair(v, Unit.OHMS)
        }
        return "$base $ohm".toLowerCase()
    }

    private fun Int.pow(i :Int) = toDouble().pow(i).toInt()

}
