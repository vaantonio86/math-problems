def count_numbers_in_string(s: str) -> int:
    """
    Count the number of numbers in a given string.
    
    Args:
    s (str): The input string containing numbers and letters.
    
    Returns:
    int: The total count of numbers in the string.
    """
    return sum(1 for char in s if char.isdigit())

def main() -> None:
    print("The number of numbers in the given string:", count_numbers_in_string(input()))

if __name__ == "__main__":
    main()
