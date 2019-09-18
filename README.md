# indecision
An application to help decide where to eat, what movie to watch, etc.

## Project Description
People have a hard time choosing when there are dozens of options placed before them. However, in my experience people can make the decision much more quickly and easily if the decision is based on only 2 options. This is where this application comes into play:
 - A pool of choices is created of things of a certain type, ex: restaurants.
 - 2 of the possible choices are placed before the user, they make their decision, and then they move on to the next. 
 - "Winners" of each match up move on a la March Madness style tournament bracket until one option remains.
 
## Team
I am looking for as many people as would like to help, as the possibilities for what kind of choices this can help solve can only grow with more time and effort invested.

## SQL
Probably a MySQL database, but I am open to suggestions.

## NoSQL
Likely some kind of document based database like Couch or Mongo, or a key value store, but again I am open to suggestions.

## Business
I haven't considered much of the possibilites in the business realm to be honest, but I could see how if the application gained enough traction that certain parties could potentially pay to be featured more or something? Sounds kinda shady to me though...

## Legal
User data would be stored to be able to track potential preferences, but the plan is not to share that data for advertising purposes or anything. Restaurant information would come from the Yelp API, movies from something like Fandango, so we'd have to operate within the specifications of their APIs and data usage.

## Technical
Using various APIs to query data such as Yelp, Fandango, etc. I know at least the Yelp API has a limit on queries per day, so the use of the database is to cache already queried data so that those APIs would have to be queried less often. User information would also be stored in a database so as to keep track of user preferences and location to better match them up with information stored in the other databases.
 - Front end envisioned as a website, preferrably done as an Angular project as that is what I most familiar with.
 - Back end server would preferrably be done in Golang (again, it's what I am most familiar with) but I am open to other options if they seem to fit the scope better.
