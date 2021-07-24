package leetcode

func intToRoman(num int) string {
    ge := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
    shi := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
    bai := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
    qian := []string{"", "M", "MM", "MMM"}
    return qian[num / 1000] + bai[num % 1000 / 100] + shi[num % 100 / 10] + ge[num % 10]
}
