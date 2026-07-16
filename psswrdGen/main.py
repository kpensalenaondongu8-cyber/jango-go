
import secrets
import string

try:
    user_input = int(input("How many charcters? "))
    word = string.ascii_letters + string.digits + string.punctuation
    password = ''.join(secrets.choice(word)for _ in range(user_input))
    print(f"This is Your Password: {password}")
except ValueError:
    print("Invalid Input")


