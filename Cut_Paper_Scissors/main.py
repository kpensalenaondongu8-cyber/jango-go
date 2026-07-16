
import random


random_list = ["rock", "paper", "scissors"]
systemwins = 0
yourwins = 0
ties = 0

while True:
    user_input = (input("Enter Your Pick: "))
    system_guess = random.choice(random_list)
    converted_input  =  user_input.lower()
    if system_guess == converted_input:
          ties+= 1
          print(f"Tie: you picked {converted_input} and computer picked {system_guess}")
          continue
    elif system_guess == "rock" and converted_input == "scissors":
          systemwins+= 1
          print(f"system Wins: you picked {converted_input} and sytem picked {system_guess}")
          continue
    elif system_guess == "scissors" and converted_input == "paper":   
          systemwins+= 1     
          print(f"system Wins: you picked {converted_input} and system picked {system_guess}")
          continue
    elif system_guess == "paper" and converted_input == "scissors":
          yourwins+= 1
          print(f"You are the Winner: you picked {converted_input} and system picked {system_guess}")
          continue
    elif system_guess == "rock" and converted_input == "paper":
          yourwins+= 1
          print(f"You are the Winner: you picked {converted_input} and system picked {system_guess}")
          continue
    elif system_guess == "scissors" and converted_input == "rock":
          yourwins+= 1
          print(f"You are the winner: you picked {converted_input} and system picked {system_guess}")
          continue  
    elif system_guess == "paper" and converted_input == "rock":
          systemwins+= 1
          print(f"system wins: you picked {converted_input} and system picked {system_guess}")
    elif converted_input == "quit":
          print("Nice Try Mate. See You Later") 
          print(f"system won: {systemwins}, you won: {yourwins}, and there was {ties} ties")
          break  
    else :
          print(f"Invalid Pick: You Picked {converted_input} and system picked {system_guess}")
          continue