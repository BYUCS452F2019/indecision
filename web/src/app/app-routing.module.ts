import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { FoodComponent } from './components/food/food.component';
import { MoviesComponent } from './components/movies/movies.component';
import { AboutComponent } from './components/about/about.component';
import { HomeComponent } from './components/home/home.component';


const routes: Routes = [
  {
    path: '',
    redirectTo: 'home',
    pathMatch: 'full'
  },
  {
    path: 'home',
    component: HomeComponent
  },
  {
    path: 'food',
    component: FoodComponent
  },
  {
    path: 'movies',
    component: MoviesComponent
  },
  {
    path: 'about',
    component: AboutComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
