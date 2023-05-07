package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
func init(){
	/*Generating a random number from math/rand package between 0-99 
	using seed and current time to make sure the 
	numbers are different any time it is run
	Below seed the random number generator
	*/
	rand.Seed(time.Now().UnixNano())
}

//Generate random number between min & max
func RandomInt(min, max int64) int64{
	/*Generates ramdom number between 0 & max-min inclusive*/
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string{
	/*
	variable s declared of type strings.Builder type, 
	this builder is used to concatenate strings together
	*/
	var sb strings.Builder
	k := len(alphabet)

	//loop
	for i := 0; i < n; i++ {
		/*
			generates random number btw 0 and k - 1 
			picks the character that fall in the index k - 1 
			and assign it to c 
		*/
		c := alphabet[rand.Intn(k)]

		/* Concatenate the character c to sb*/
		sb.WriteByte(c)
	}

	return "s"
}