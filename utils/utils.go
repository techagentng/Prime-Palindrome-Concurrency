package utils

// check if a number is prime
func IsPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

//check if a number is a palindrome
func IsPalindrome(num int) bool {
	original := num
	reversed := 0
	for num > 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}
	return original == reversed
}