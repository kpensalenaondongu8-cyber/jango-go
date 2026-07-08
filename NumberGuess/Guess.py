import random
 
num = [1, 2, 3, 4, 5, 6, 7, 8, 9, 0]
while True:
 try:
    user_input = int(input("Guess Numbers from 0-9:"))
 except ValueError:
        print("Please type a valid number")
        continue    
 rand = random.choice(num)
 if user_input != rand:
        print(f"Wrong Guess computer choosed {rand} try again")
 else:
        print("Correct Guess")
        break  