func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    if len(num1) < len(num2) {
        num1, num2 = num2, num1
    }
    
    res := ""
    n1 := len(num1)
    n2 := len(num2)
    for i := n2 - 1; i >= 0; i-- {
        d2, _ := strconv.Atoi(string(num2[i]))
        if d2 == 0 {
            continue
        }
        
        flag := 0
        product := ""
        for j := 0; j < n2 - 1 - i; j++ {
            product += "0"
        }
        for j := n1 - 1; j >= 0; j-- {
            d1, _ := strconv.Atoi(string(num1[j]))
            a := (d1 * d2 + flag) % 10
            flag = (d1 * d2 + flag) / 10
            product += fmt.Sprint(a)
        }
        if flag > 0 {
            product += fmt.Sprint(flag)
        }
        
        fmt.Println(res, product)
        if res == "" {
            res = product
        } else {
            res = sum(res, product)
        }
    }
    
    if res == "" {
        return "0"
    }
    return reverse(res)
}

func sum(num1, num2 string) string {
    res := ""
    flag := 0
    for i := 0; i < len(num1) || i < len(num2); i++ {
        if i >= len(num1) {
            d1, _ := strconv.Atoi(string(num2[i]))
            a := (d1 + flag) % 10
            flag = (d1 + flag) / 10
            res += fmt.Sprint(a)
            
        } else if i >= len(num2) {
            d1, _ := strconv.Atoi(string(num1[i]))
            a := (d1 + flag) % 10
            flag = (d1 + flag) / 10
            res += fmt.Sprint(a)
            
        } else {
            d1, _ := strconv.Atoi(string(num1[i]))
            d2, _ := strconv.Atoi(string(num2[i]))
            a := (d1 + d2 + flag) % 10
            flag = (d1 + d2 + flag) / 10
            res += fmt.Sprint(a)
        }
    }
    if flag > 0 {
        res += fmt.Sprint(flag)
    }
    return res
}

func reverse(num string) string {
    res := ""
    for i := len(num) - 1; i >= 0; i-- {
        res += string(num[i])
    }
    return res
}
