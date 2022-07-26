package main

import (
	"ntools"
	"os"
	"os/exec"
	"regexp"
	"strings"
)


func main(){

	
	for { 
	lebok := dovizAl("currencies/eur-try") +
	"\n" + dovizAl("currencies/usd-try") + 
	"\n" + dovizAl("currencies/gbp-try") + 
	"\n" + dovizAl("commodities/brent-oil") +
	"\n" + dovizAl("crypto/bitcoin/btc-usd") + 
	"\n" + dovizAl("crypto/ethereum/eth-usd") + "\033[?25l"

	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	print(lebok)
	//time.Sleep(time.Second * 15)

}
}

func dovizAl(uzanti string) string{
	kur, err := ntools.WebToText("https://www.investing.com/" + uzanti)

	if err != nil {
		panic(err)
	}
	getCur := regexp.MustCompile(`instrument-price-last">(.*?)</span>`).FindAllStringSubmatch(kur,1)
	
	return(strings.ToUpper(uzanti) + " : " + string(getCur[0][1]))

}