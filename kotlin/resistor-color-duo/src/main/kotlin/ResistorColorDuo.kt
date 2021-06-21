object ResistorColorDuo {
    fun value(vararg colors: Color): Int = colors.take(2).map(Color::ordinal).reduce{acc, element -> acc*10 + element}
}
