import { Component } from '@angular/core';
import { NavigationEnd, Router, RouterOutlet } from '@angular/router';
import { HomeComponent } from "./components/home/home.component";
import { CommonModule } from '@angular/common';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { AuthService } from './services/auth.service';
import { interval, Subscription } from 'rxjs';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, CommonModule, MatSnackBarModule],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'Frontend';
  loginStatus: boolean = false;  // Variable para almacenar el estado de autenticación
  private routerSubscription: Subscription | undefined;  // Para gestionar la suscripción al router
  private authSubscription: Subscription | undefined;  // Para gestionar la suscripción al servicio de autenticación

  constructor(private router: Router, private authService: AuthService) {}

  ngOnInit() {
    // Escuchar los eventos de navegación
    this.routerSubscription = this.router.events.subscribe(event => {
      if (event instanceof NavigationEnd) {
        this.updateLoginStatus();  // Actualiza el estado de login en cada navegación
      }
    });
  }

  ngOnDestroy() {
    // Limpiar las suscripciones cuando el componente se destruya
    if (this.routerSubscription) {
      this.routerSubscription.unsubscribe();
    }
    if (this.authSubscription) {
      this.authSubscription.unsubscribe();
    }
  }

  // Método para actualizar el estado de autenticación
  private updateLoginStatus() {
    // Cancelar la suscripción anterior antes de crear una nueva
    if (this.authSubscription) {
      this.authSubscription.unsubscribe();
    }

    // Hacer la llamada para verificar si el usuario está autenticado
    this.authSubscription = this.authService.isAuthenticated().subscribe(
      (isAuth) => {
        this.loginStatus = isAuth;  // Actualiza la variable según el estado de autenticación
        console.log('Autenticado:', this.loginStatus);
      },
      () => {
        this.loginStatus = false;  // Si hay error, el usuario no está autenticado
        console.log('No autenticado');
      }
    );
  }
}
