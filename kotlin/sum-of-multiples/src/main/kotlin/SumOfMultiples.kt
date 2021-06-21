object SumOfMultiples {
    fun sum(factors: Set<Int>, limit: Int): Int {
        return (0 until limit).filter { i -> factors
                .filter { it != 0 }
                .any { f -> i % f == 0 } }
            .sum()
    }
}
