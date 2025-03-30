def square_sum(numbers):
    return sum(x**2 for x in numbers)

numbers = [1, 2, 3, 4]
result = square_sum(numbers)
print(result) # Output: 50
