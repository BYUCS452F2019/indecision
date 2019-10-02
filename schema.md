Indecision Schema

# The Users table contains a record of all users that have used the application. 
User(UserID, UserName)

Restaurant(RestID, RestName, Category, Location, Price)

Movie ( )

# Like a history of the previous winners
Winner (WinnerID, UserID, RestID, Date)
 - Foreign Key UserID references User
 - Foreign Key RestID references Restaurant
