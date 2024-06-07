using System;
using System.Text.RegularExpressions;

public class ValidPalindrome {
    public bool IsPalindrome(string s) {
        if (string.IsNullOrEmpty(s)) {
            return true;
        }
        s = Regex.Replace(s.ToLower(), "[^a-z0-9]", "");

        int start = 0;
        int end = s.Length - 1;
        
        while (start < end) {
            if (s[start] != s[end]) {
                return false;
            }
            start++;
            end--;
        } 
        return true; 
    }
}
