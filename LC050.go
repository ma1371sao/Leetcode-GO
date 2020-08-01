func myPow(x float64, n int) float64 {
    if x == 0 {
        return 0
    }
    if n == 0 {
        return 1
    }
    if n == 1 {
        return x
    }
    if n == -1 {
        return 1.0 / x
    }
    sub := myPow(x, n/2)
    return sub * sub * myPow(x, n%2)
}
