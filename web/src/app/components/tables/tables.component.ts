import { Component, OnInit, ViewChild } from '@angular/core';
import { Restaurant, Movie, User } from 'src/app/objects';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { MatTableDataSource, MatPaginator } from '@angular/material';

@Component({
  selector: 'app-tables',
  templateUrl: './tables.component.html',
  styleUrls: ['./tables.component.scss']
})
export class TablesComponent implements OnInit {
  chosenType: string;
  choiceTypes = [
    "Food",
    "Movies",
  ]

  tableColumns: string[] = [];

  restaurantList: Restaurant[] = [];
  movieList: Movie[] = [];
  userList: User[] = [];

  private headers: HttpHeaders;

  dataSource: MatTableDataSource<any>;
  @ViewChild(MatPaginator, { static: false }) paginator: MatPaginator;

  constructor(private http: HttpClient) {
    this.headers = new HttpHeaders({
      'Content-Type': 'application/json',
      'Access-Control-Allow-Origin': '*'
    });
  }

  ngOnInit() {
  }

  getTableInfo() {
    if (this.chosenType === "Food") {
      this.http.get('http://localhost:5500/choices/food', { headers: this.headers}).toPromise().then((answer) => {
        this.restaurantList = answer as Restaurant[];
        this.dataSource = new MatTableDataSource(this.restaurantList);
        this.dataSource.paginator = this.paginator;
        console.log(this.dataSource);

        if (this.restaurantList && this.restaurantList.length > 0) {
          this.tableColumns = Object.keys(this.restaurantList[0]);
        }
      });
    }
    if (this.chosenType === "Movies") {
      this.http.get('http://localhost:5500/choices/movies', { headers: this.headers}).toPromise().then((answer) => {
        this.movieList = answer as Movie[];
        this.dataSource = new MatTableDataSource(this.movieList);
        this.dataSource.paginator = this.paginator;

        if (this.movieList && this.movieList.length > 0) {
          this.tableColumns = Object.keys(this.movieList[0]);
        }
      });
    }
  }

}
