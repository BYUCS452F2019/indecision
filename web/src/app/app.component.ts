import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'web';
  activeLink: string;
  links = [
    "play",
    "tables"
  ]

  constructor(private route: Router) {
    this.activeLink = window.location.pathname.split("/", 2)[1];
  }

  toTitle(word: string): string {
    if (word === undefined || word == null) {
      return "title broken...";
    }

    return word.charAt(0).toUpperCase() + word.slice(1);
  }
}
