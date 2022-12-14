package main

import (
        "encoding/base64"
        "encoding/base32"
        "encoding/ascii85"
        "encoding/hex"
        "flag"
        "fmt"
        "github.com/btcsuite/btcutil/base58"
        "github.com/dasio/base45"
        "github.com/catinello/base62"
        "github.com/Max-Sum/base32768"
        "github.com/usk81/base65536"
        "github.com/mtraver/base91"
        "github.com/martinlindhe/base36"
        "github.com/fatih/color"
        "moul.io/banner"
)

func main () {

        fmt.Println(banner.Inline("base cracker"))
        fmt.Println()

        ct := flag.String("ct","","target ciphertext")
        flag.Parse()

        ans2,err := hex.DecodeString(*ct)
        if err != nil {
                fmt.Println("Base16 :",err)
        } else {
                color.Green("Base16 : " + string(ans2))
        }

        ans1,err := base32.StdEncoding.DecodeString(*ct)
        if err != nil {
                fmt.Println("Base32 :",err)
        } else {
                color.Green("Base32 : " + string(ans1))
        }

        ans9 := base36.Decode(*ct)
        color.Green("Base36 : " +  string(ans9))

        ans3,err := base45.DecodeString(*ct)
        if err != nil {
                fmt.Println("Base45 :",err)
        } else {
                color.Green("Base45 : " + string(ans3))
        }

        ans4 := base58.Decode(*ct)
        color.Green("Base58 : " + string(ans4))

        ans5,err := base62.Decode(*ct)
        if err != nil {
                fmt.Println("Base62 :",err)
        } else {
                color.Green("Base62 : " + string(ans5))
        }

        ans,err := base64.StdEncoding.DecodeString(*ct)
        if err != nil {
                fmt.Println("Base64 :", err)
        } else {
                color.Green("Base64 : " + string(ans))
        }

        newbuffer := make([]byte,len(*ct))
        d1, d2, err := ascii85.Decode(newbuffer,[]byte(*ct),true)
        if err != nil {
                fmt.Println("Base85 :",err,d1,d2)
        } else {
                color.Green("Base85 : " + string(newbuffer))
        }

        ans8,err := base91.StdEncoding.DecodeString(*ct)
        if err != nil {
                fmt.Println("Base91 :",err)
        } else {
                color.Green("Base91 : " + string(ans8))
        }

        ans6,err := base32768.SafeEncoding.DecodeString(*ct)
        if err != nil {
                fmt.Println("Base32768 :",err)
        } else {
                color.Green("Base32768 : " + string(ans6))
        }

        ans7 := base65536.Decode(*ct)
        color.Green("Base65536 : " + string(ans7))

}
