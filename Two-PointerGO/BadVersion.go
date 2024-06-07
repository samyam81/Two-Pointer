package main

/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    ss81:=0
	kdb:=n
	var mid int

	for ss81<kdb{
		mid=ss81+(kdb-ss81)/2

		if isBadVersion(mid) {
			kdb=mid
		} else{
			ss81=mid+1
		}
	}
	 
	return ss81
}

func isBadVersion(mid int) bool {
	panic("unimplemented")
}