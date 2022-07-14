package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
}

func numUniqueEmails(emails []string) int {
	valid_emails := make(map[string]bool)
	no_mailids := 0
	for _, email := range emails {
		split := strings.Index(email, "@")
		local := email[:split]
		pls := strings.Index(local, "+")
		if pls > 0 {
			local = local[:pls]
		}
		local = strings.ReplaceAll(local, ".", "")
		local = local + email[split:]
		if _, ok := valid_emails[local]; !ok {
			valid_emails[local] = true
			no_mailids++
		}
	}
	fmt.Println(valid_emails)
	return no_mailids
}
