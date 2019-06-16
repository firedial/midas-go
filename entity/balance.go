package entity

import (
    "strings"
    "strconv"
)

type Balance struct {
    BalanceId int `json:"balance_id"`
    Amount int `json:"amount"`
    Item string `json:"item"`
    KindId int `json:"kind_id`
    PurposeId int `json:"purpose_id"`
    PlaceId int `json:place_id"`
    Date string `json:"date"`
}

type Balances []Balance

func IsSuitableBalances(balances Balances) bool {
    for _, balance := range balances {
        if !isSuitableBalance(balance) {
            return false
        }
    }

    return true
}

func isSuitableBalance(balance Balance) bool {
    if balance.Amount == 0 {
        return false
    }
    if !isSuitableString(balance.Item) {
        return false
    }
    if balance.KindId <= 0 {
        return false
    }
    if balance.PurposeId <= 0 {
        return false
    }
    if balance.PlaceId <= 0 {
        return false
    }
    if !isSuitableDate(balance.Date) {
        return false
    }

    return true
}

// 文字列の中に['":. ]の5種類の半角記号が入っていないことを見る
// @return bool true: それらが入っていない / false: それらが入っている
func isSuitableString(str string) bool {
    checkChars := []string{"'", "\"", ":", ".", " "}
    
    if str == "" {
        return false
    }

    for _, c := range checkChars {
        if strings.Contains(str, c) {
            return false
        }
    }

    return true
}

func isSuitableDate(str string) bool {
    splitStr := strings.Split(str, "/")

    // "/" で 3 つに分割できない場合は不適
    if len(splitStr) != 3 {
        return false
    }

    year, erry := strconv.Atoi(splitStr[0])
    month, errm := strconv.Atoi(splitStr[1])
    day, errd := strconv.Atoi(splitStr[2])

    if erry != nil || errm != nil || errd != nil {
        return false
    }

    // 年が 1000 から 9999 の間であるかどうか
    if !(1000 <= year && year <= 9999) {
        return false
    }

    // 月が 1 から 12 の間であるかどうか
    if !(1 <= month && month <= 12) {
        return false
    }

    // 日が 1 から 31 の間であるかどうか
    // 31日がない月もあるがそこまでは考えない
    if !(1 <= day && day <= 31) {
        return false
    }

    return true
}
