import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Movie } from '../../objects';

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

  baseList: Movie[] = [];
  currentList: Movie[] = [];
  roundWinners: Movie[] = [];
  winner: Movie;

  private headers: HttpHeaders;

  constructor(private http: HttpClient) {
    
  }

  ngOnInit() {
    this.winner = null;
    this.headers = new HttpHeaders({
      'Content-Type': 'application/json',
      'Access-Control-Allow-Origin': '*'
    });
    this.http.get('http://localhost:5500/choices/movies', { headers: this.headers}).toPromise().then((answer) => {
      this.baseList = answer as Movie[];
      for (const r of this.baseList) {
        this.currentList.push(r);
      }
      console.log('Movie list is', this.baseList);
    });
  }

  selectWinner(r: Movie) {
    this.roundWinners.push(r);
    if (this.currentList.length === 3) {
      this.currentList.splice(0, 3);
    } else {
      this.currentList.splice(0, 2);
    }

    console.log('Selection:', r);
    console.log('Round winners:', this.roundWinners);
    console.log('What is left over:', this.currentList);

    if (this.currentList.length === 0) {
      this.currentList = [];
      this.currentList.push(...this.roundWinners);
      this.roundWinners = [];
    }
  }

  reset() {
    this.currentList = [];
    this.currentList.push(...this.baseList);
  }
}
