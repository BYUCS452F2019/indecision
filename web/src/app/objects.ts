export class User {
    id: string;
    username: string;
}

export class Choice {
    choiceID: string;
    pictureUrl: string;
    type: string;
}

export class Restaurant extends Choice {
    constructor() {
        super();
    }
    restaurantID: string;
    name: string;
    category: string;
    location: string;
    price: string;
}

export class Movie extends Choice {
    movieID: string;
    title: string;
    genre: string;
}
