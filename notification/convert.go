package notification

import (
	"fmt"
	"math"
	"strconv"
)

/**
 * USドルを日本円に変換する
 *
 * amount: 現在のAWSのコスト（返却値が文字列なのでstringになってる）
 * currentJpy: 現在の日本円のレート
 */
func ConvertUsDollarToJpy(amount string, currentJpy float64) string {
	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		fmt.Println(err)
	}

	// 小数点第一位で四捨五入
	currentAmount := math.Round(amountFloat * currentJpy)
	resultInt := int(currentAmount)

	jpy := strconv.Itoa(resultInt)

	return jpy
}
