import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Restaurant } from '../../objects';

export class ListItem {
  name: string;
  description: string;
  price: number;
  imageURL?: string;
}
@Component({
  selector: 'app-food',
  templateUrl: './food.component.html',
  styleUrls: ['./food.component.scss']
})
export class FoodComponent implements OnInit {
  redFlipped = false;
  blueFlipped = false;
  yellowFlipped = false;
  greenFlipped = false;

  baseList: Restaurant[] = [];
  winner: Restaurant;
  // baseList: ListItem[] = [
  //   {
  //     name: 'McDonalds',
  //     description: 'A greasy fast food place.',
  //     price: 1,
  //     imageURL: 'https://cdn.junglecreations.com/wp/junglecms/2019/04/77a4227c-mcdonalds-free-breakfast-feature.jpg'
  //   },
  //   {
  //     name: 'Wendy\'s',
  //     description: 'A greasy fast food place with square burgers.',
  //     price: 1,
  //     imageURL: 'https://130e178e8f8ba617604b-8aedd782b7d22cfe0d1146da69a52436.ssl.cf1.rackcdn.com/wendys-reaches-50-million-breach-settlement-banks-showcase_image-8-a-12032.jpg'
  //   },
  //   {
  //     name: 'Chik-Fil-A',
  //     description: 'An also greasy fast food place that mostly sells chicken.',
  //     price: 2,
  //     imageURL: 'https://assets.simpleviewinc.com/simpleview/image/fetch/c_limit,q_75,w_1200/https://assets.simpleviewinc.com/simpleview/image/upload/crm/sandysprings/Chick-fil-a-Northridge-DTO---logo-resized-95aaa9a05056b3a_95aaaa9b-5056-b3a8-493eb82fb0c5d336.jpg'
  //   }
  // ];
  // winner: ListItem;

  private headers: HttpHeaders;

  constructor(private http: HttpClient) {
    this.winner = null;
    this.headers = new HttpHeaders({
      'Content-Type': 'application/json',
      'Access-Control-Allow-Origin': '*'
    });
    this.http.get('http://localhost:5500/choices/food', { headers: this.headers}).toPromise().then((answer) => {
      this.baseList = answer as Restaurant[];
      console.log('restaurant list is', this.baseList);
    });
  }

  ngOnInit() {
  }

}
