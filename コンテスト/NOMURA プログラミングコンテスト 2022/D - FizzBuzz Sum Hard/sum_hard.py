def gcd(a, b):
	while b != 0:
		t = b
		b = a % b
		a = t

	return a

def lcm(a, b):
	result = a * b // gcd(a, b)
	return result


if __name__ == "__main__":
	n, a, b = map(int, input().split())
	c = lcm(a, b)

	sum = n * (n+1) / 2
	sumA = (a + n // a * a) * (n // a) / 2
	sumB = (b + n // b * b) * (n // b) / 2
	sumC = (c + n // c * c) * (n // c) / 2

	result = sum - (sumA + sumB) - sumC
	print(result)
