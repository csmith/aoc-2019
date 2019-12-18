package main

import "strings"

func compress(parts []string) (main, a, b, c string) {
	for _, candidateA := range prefixes(parts) {
		remainderA := replace(parts, candidateA, "A")
		for _, candidateB := range prefixes(remainderA) {
			remainderB := replace(remainderA, candidateB, "B")
			for _, candidateC := range prefixes(remainderB) {
				remainderC := replace(remainderB, candidateC, "C")
				if acceptableRoutine(remainderC) {
					return strings.Join(remainderC, ","), strings.Join(candidateA, ","), strings.Join(candidateB, ","), strings.Join(candidateC, ",")
				}
			}
		}
	}

	return
}

func replace(parts []string, function []string, name string) []string {
	var res []string
	for start := 0; start < len(parts); start++ {
		if start+len(function) <= len(parts) && equals(parts[start:start+len(function)], function) {
			res = append(res, name)
			start += len(function) - 1
		} else {
			res = append(res, parts[start])
		}
	}
	return res
}

func equals(a, b []string) bool {
	for i, p := range a {
		if b[i] != p {
			return false
		}
	}
	return true
}

func prefixes(parts []string) [][]string {
	start := 0
	for start < len(parts) && (parts[start] == "A" || parts[start] == "B") {
		start++
	}
	if start == len(parts) {
		return nil
	}

	var res [][]string
	for end := start + 1; end <= len(parts) && acceptableFunction(parts[start:end]); end++ {
		res = append(res, parts[start:end])
	}

	// Reverse the slice so the longest element is first
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

func acceptableRoutine(parts []string) bool {
	for _, p := range parts {
		if p != "A" && p != "B" && p != "C" {
			return false
		}
	}
	return true
}

func acceptableFunction(parts []string) bool {
	length := 0
	for i, part := range parts {
		if i > 0 {
			length++
		}
		length += len(part)

		if strings.ContainsAny(part, "ABC") || length > 20 {
			return false
		}
	}

	return true
}
