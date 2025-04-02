def find_greatest_common_divisor(a: int, b: int) -> int:
    while b != 0:
        a, b = b, a % b
    return a

def is_prime(n: int) -> bool:
    if n < 2:
        return False
    for i in range(2, int(n ** 0.5) + 1):
        if n % i == 0:
            return False
    return True

# Example usage
if __name__ == "__main__":
    a = 30  # Replace with any number you want to find the greatest common divisor of
    print("GCD of", a, "and", a, "is:", find_greatest_common_divisor(a, a))
    print("Is", a, "prime?", is_prime(a))
