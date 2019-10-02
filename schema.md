# Indecision Schema

#### The User table contains a record of all users that have used the application. 
User(UserID, UserName)

#### The Restaurant table contains records of various places to eat that will be used as options in the decision making process.
Restaurant(RestID, RestName, Category, Location, Price)

#### The Movie table contains records of various movies to watch that will be used as options in the decision making process.
Movie ( )

#### The RestaurantWinner table contains records of the winners from previous decisions made. The table contains a reference to the user and the restaurant that was chosen.
RestaurantWinner (WinnerID, UserID, RestID, Date)
 - Foreign Key UserID references User
 - Foreign Key RestID references Restaurant

#### The MovieWinner table contains records of the winners from previous decisions made. The table contains a reference to the user and the movie that was chosen.
MovieWinner (WinnerID, UserID, RestID, Date)
 - Foreign Key UserID references User
 - Foreign Key MovieID references Movie
