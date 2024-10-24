package dns

import (
	"fmt"
	"net"
	"os"
	"sort"
)

func Search(domain string) {
	var errors []error

	fmt.Println()
	fmt.Printf("DNS for %s\n", domain)
	fmt.Println()

	// IPアドレス
	ips, err := net.LookupIP(domain)
	if err != nil {
		errors = append(errors, err)
	} else {
		for _, ip := range ips {
			if ip.To4() != nil {
				fmt.Printf("IPv4 : %s\n", ip)
			} else {
				fmt.Printf("IPv6 : %s\n", ip)
			}
		}
	}

	// Aレコード
	as, err := net.LookupHost(domain)
	if err != nil {
		errors = append(errors, err)
	} else {
		for _, a := range as {
			fmt.Printf("A    : %s\n", a)
		}
	}

	// NSレコード
	nss, err := net.LookupNS(domain)
	if err != nil {
		errors = append(errors, err)
	} else {
		sort.Slice(nss, func(i, j int) bool {
			return nss[i].Host < nss[j].Host
		})
		for i, ns := range nss {
			fmt.Printf("NS%d  : %s\n", i+1, ns.Host)
		}
	}

	// MXレコード
	mxs, err := net.LookupMX(domain)
	if err != nil {
		errors = append(errors, err)
	} else {
		for _, mx := range mxs {
			fmt.Printf("MX   : %s\n", mx.Host)
		}
	}

	// CNAMEレコード
	cnames, err := net.LookupCNAME(domain)
	if err != nil {
		errors = append(errors, err)
	} else {
		fmt.Printf("CNAME: %s\n", cnames)
	}

	// TXTレコード
	txts, err := net.LookupTXT(domain)
	if err != nil {
		errors = append(errors, err)
	} else {
		for _, txt := range txts {
			fmt.Printf("TXT  : %s\n", txt)
		}
	}

	// エラーチェック
	if len(errors) > 0 {
		for _, err := range errors {
			fmt.Println("❌", err)
		}
		os.Exit(1)
	}
}