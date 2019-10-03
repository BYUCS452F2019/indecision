# Indecision Schema

The first column is the key column in each table.

User Table
---
#### Description:
 - The User table contains a record of all users that have used the application. 
#### Schema:
User(UserID, UserName)

---

Restaurant Table
---
#### Description:
 - The Restaurant table contains records of various places to eat that will be used as options in the decision making process. Each column is an attribute of a restaurant that will be displayed to the user.
#### Schema:
Restaurant(RestID, RestName, Category, Location, Price) |

---

Movie Table
---
#### Description:
 - The Movie table contains records of various movies to watch that will be used as options in the decision making process.
#### Schema:
Movie (MovieId, Title, Genre )

---

RestaurantWinner Table
---
#### Description:
 - The RestaurantWinner table contains records of the winners from previous decisions made. The table contains a reference to the user and the restaurant that was chosen.
#### Schema:
RestaurantWinner (WinnerID, UserID, RestID, Date)
  - Foreign Key UserID references User
  - Foreign Key RestID references Restaurant
 
---

MovieWinner Table
---
#### Description:
- The MovieWinner table contains records of the winners from previous decisions made. The table contains a reference to the user and the movie that was chosen.
#### Schema:
MovieWinner (WinnerID, UserID, RestID, Date)
 - Foreign Key UserID references User
 - Foreign Key MovieID references Movie
