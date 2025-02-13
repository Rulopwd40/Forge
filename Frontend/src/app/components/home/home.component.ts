import { Component } from '@angular/core';
import { FormGroup, FormControl,ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { User } from '../../models/user';
import { UserService } from '../../services/user.service';
import { AbstractControl, ValidationErrors, ValidatorFn, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { MatSnackBar } from '@angular/material/snack-bar';
import { AuthService } from '../../services/auth.service';


@Component({
  selector: 'app-home',
  imports: [ReactiveFormsModule,CommonModule],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss',
  providers: [UserService]
})


export class HomeComponent {


 
  loginForm: FormGroup;
  registerForm: FormGroup<any>;

  showPasswordArray: { [key: string]: boolean } = 
    { 
      'login-password': false,
      'register-password': false,
      'confirm-password': false 
    }

  constructor(private userService: UserService, private authService: AuthService, private router:Router, private snackBar: MatSnackBar) {
    this.loginForm = new FormGroup({
      username: new FormControl('', [Validators.required, Validators.minLength(3)]),
      password: new FormControl('', [Validators.required, Validators.minLength(6)])
    });
    
    this.registerForm = new FormGroup(
      {
        username: new FormControl('', [Validators.required, Validators.minLength(3)]),
        password: new FormControl('', [Validators.required, Validators.minLength(6)]),
        confirm: new FormControl('', [Validators.required, Validators.minLength(6)]),
        name: new FormControl('', [Validators.required]),
        email: new FormControl('', [Validators.required, Validators.email])
      },
      
    );
  }

  login = true;

  slide() {
    this.login = !this.login;
  }

  loginUser(event : any) {
    if(!this.loginForm.valid) {
      this.snackBar.open('Invalid form', 'Close', { duration: 3000 });
      throw new Error('Invalid form');
    }
    const credentials = {
      username: this.loginForm.value.username,
      email: this.loginForm.value.email,
      password: this.loginForm.value.password
    }

    this.authService.login(credentials).subscribe({
      next:(response) => {
        this.snackBar.open("Authenticated correctly", 'Close', {duration: 3000})
      },
      error: (error) => {
        this.snackBar.open("Unauthorized", 'Close', {duration: 3000})
      }
    })
  }

  registerUser(event : any) {
    if(this.registerForm.value.password !== this.registerForm.value.confirm) {
      this.snackBar.open('Passwords do not match', 'Close', { duration: 3000 });
      throw new Error('Passwords do not match');
    }
    if(!this.registerForm.valid) {
      this.snackBar.open('Invalid form', 'Close', { duration: 3000 });
      throw new Error('Invalid form');
    }
    const newUser: User = {
      username: this.registerForm.value.username,
      password: this.registerForm.value.password,
      name: this.registerForm.value.name,
      level:0,
      email: this.registerForm.value.email
    };
    this.userService.registerUser(newUser).subscribe({
      next: (response) => {
        this.snackBar.open('User registered successfully', 'Close', { duration: 3000 });
        setTimeout(() => {
            location.reload();
        }, 2000);
      },
      error: (error) => {
        this.snackBar.open(`Error registering user: ${error.message}`, 'Close', { duration: 3000 });
        console.error(error);}
    });
  }
  toggleShowPassword(arg0: string) {
    this.showPasswordArray[arg0] = !this.showPasswordArray[arg0];
    
  }

  showPassword(arg0: string):boolean {
    return this.showPasswordArray[arg0];
  }
  
}
