import { Component } from '@angular/core';
import { NgIcon } from '@ng-icons/core';
import { heroUserSolid } from '@ng-icons/heroicons/solid';
import { AuthService } from '../../services/auth.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-menu',
  imports: [NgIcon],
  templateUrl: './menu.component.html',
  styleUrl: './menu.component.scss'
})
export class MenuComponent {
  
  user= {
    username: '',
    name:'',
    level: 0,
  };
  
  constructor(private authService:AuthService, private snackBar: MatSnackBar){
    
  }
  
  ngOnInit(){
    this.authService.getProfile().subscribe({
      next:(response) => {
        this.user.username= response.username;
      },
      error: (error) => {
        this.snackBar.open("Error getting username", 'Close', {duration: 3000})
      }
    })
  }
  
  levelColor() {

  }
}
