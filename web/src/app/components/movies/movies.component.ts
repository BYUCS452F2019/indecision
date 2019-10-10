import { Component, OnInit } from '@angular/core';
import { ListItem } from '../food/food.component';

@Component({
  selector: 'app-movies',
  templateUrl: './movies.component.html',
  styleUrls: ['./movies.component.scss']
})
export class MoviesComponent implements OnInit {
  redFlipped = false;
  blueFlipped = false;
  yellowFlipped = false;
  greenFlipped = false;

  baseList: ListItem[] = [
    {
      name: 'Star Wars',
      description: 'A long time ago, in a galaxy far, far away...',
      price: 1,
      imageURL: 'https://vignette.wikia.nocookie.net/starwars/images/c/cc/Star-wars-logo-new-tall.jpg/revision/latest?cb=20190313021755'
    },
    {
      name: 'Toy Story',
      description: 'An animated film about toys that come to life when humans aren\'t around.',
      price: 1,
      imageURL: 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/0a/Toy_Story_logo.svg/1200px-Toy_Story_logo.svg.png'
    },
    {
      name: 'Won\'t You Be My Neighbor?',
      description: 'A documentary about the life of Mr. Rogers.',
      price: 2,
      imageURL: 'https://www.92y.org/92streety/media/img/talks/lg/wont-you-by-my-neighbor.jpg'
    }
  ];
  winner: ListItem;

  constructor() { }

  ngOnInit() {
  }
}
