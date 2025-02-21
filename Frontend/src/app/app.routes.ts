import { Routes } from '@angular/router';
import { HomeComponent } from './components/home/home.component';
import { MenuComponent } from './components/menu/menu.component';
import { authGuard } from './guards/auth.guard';


export const routes: Routes = [
    { path: '', redirectTo: 'home', pathMatch: 'full' }, 
    { path: 'home', component: HomeComponent },
    { path: 'menu', component: MenuComponent, canActivate: [authGuard] }, 
    { path: '**', redirectTo: 'home' } 
];