# Indecision Schema

User Table
---
#### Description:
 - The User table contains a record of all users that have used the application. 
#### Schema:
User(UserID, UserName)

---

Choice Table
---
#### Description:
 - The choice table contains IDs for all the available choices and the information that is shared across all choice types.
#### Schema:
Choice (ChoiceID, Url, PictureURL)
  - Foreign Key UserID references User

---

Restaurant Table
---
#### Description:
 - The Restaurant table is a choice subtype that contains records of various places to eat that will be used as options in the decision making process. It is associated with a Choice.
#### Schema:
Restaurant(RestID, RestName, Category, Location, Price, ChoiceId) |
  - Foreign Key ChoiceId references Choice

---

Movie Table
---
#### Description:
 - The Movie table is a choice subtype and contains records of various movies to watch that will be used as options in the decision making process. It is associated with a Choice.
#### Schema:
Movie (MovieId, Title, Genre, ChoiceId)
  - Foreign Key ChoiceID references Choice

---
Round Table
---
#### Description:
 - The Round table contains records from each round a user has participated in. The table contains a reference to the user, the round type and the status of the round.
#### Schema:
Round (RoundID, UserID, Type, Status)
  - Foreign Key UserID references User
---

Match Table
---
#### Description:
 - The Match table contains records of the individual matches from each round. The table contains a reference to the round, the winning choice, the losing choice and where the match occured in the round.
#### Schema:
Match (MatchId, RoundId, LosingId, WinningId, MatchNumber)
  - Foreign Key RoundId references Round
  - Foreign Key LosingId references Choice
  - Foreign Key LosingId references Choice
---
