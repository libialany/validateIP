package main
import (
	"main"
	"net"
)
func main() {
	validIPV4 := "10.50.260.553"
	checkIPAddress(validIPV4)

	invalidIPV4 := "1000.450.210.253"
	checkIPAddress(invalidIPV4)

	validIPV6 := "2001:678jk:85a3:0000:0000:8a2e:0370:7334"
	checkIPAddress(validIPV6)

	invalidIPV6 := "2001:0db8:85a3:0000:00"
	checkIPAddress(invalidIPV6)	
}
