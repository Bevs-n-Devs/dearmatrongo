package utils

/*
Checks to see if the user wants to make a claim

Returns true if the user wants to make a claim, false otherwise
*/
func MakeCklaimCheck(claim string) bool {
	return claim == "Yes"
}
