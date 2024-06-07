package main

func canWinNim(n int) bool {
	if n%4==0 {
		return false
	}
	return true
}

// This solution belong to game theory. There is no point in creating a repo for single solution.