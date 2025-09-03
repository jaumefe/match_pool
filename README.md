# Match Pool
A project for a simple pool for football competitions.

The game consists of a pool of teams and goal scorers. Teams are classified in 4 groups depending on their chances of winning the competition. Each team is given an amount of points and the sum of their points cannot exceed an predefined value.

The players are also divided into groups according to their teams.

Hereinafter, I will give some tips to configure the game as I usually do, because I think it is the most competitive way.

## Teams organisation in groups
The groups don't have a predefined amount of teams because depending on the competitions it may change. Normally, group 3 is the largest, followed by group 2, group 4 and finally, group 1.

I assign points from 10 to 1 to the teams depending on their chances to win, 10 to the team with the greatest and 1 to the lowest chances.

To have a competitive pool, the rule of thumb I use is to have an approximate ratio of `assigned points / teams = 5` or `max points /number pool teams = 4.7`

## Teams / Players selection
This Match Pool project is ready to choose: 1 team from Group 1, 2 from Group 2, 2 from Group 3 and 2 from Group 4. The total sum of points of the chosen team cannot exceed the maximum value defined.

Concerning goal scorers, the expected amount of players to be chosen are: 1 from Group 1, 2 from Group 2, 3 from Group 3 and 2 from Group 4.

## Configuration
There are some aspects that can be configured by the game administrator. It can be done directly by writing on the database (not recommended) or by using the website given:

- Points per round: Every round that the team goes on may have a different way of counting the points
- Points per goal: A goal in a different round may be worth different
- Max points: The maximum value of the sum of teams points
- If there is a pool stage, it may be worth to set points for being the first or second placed.
- Loading on a csv file both teams and goal scorers

By default, when launching the game the sqlite database will be created with a player whose username and password is `admin`. From the website, it is possible to change both user and password.
