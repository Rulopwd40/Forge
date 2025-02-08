import { Component } from '@angular/core';
import { FormGroup, FormControl,ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
@Component({
  selector: 'app-home',
  imports: [ReactiveFormsModule,CommonModule],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss'
})
export class HomeComponent {


  loginForm: FormGroup;
  registerForm: FormGroup<any>;

  constructor() {
    this.loginForm = new FormGroup({
      username: new FormControl(''),
      password: new FormControl('')
    });
    this.registerForm = new FormGroup({
      username: new FormControl(''),
      password: new FormControl(''),
      confirm: new FormControl(''),
      email: new FormControl('')
    });
  }

  login = true;

  slide() {
    this.login = !this.login;
  }

  loginUser(Event : any) {
    console.log(this.loginForm.value);
    }

  registerUser(Event : any) {
    console.log(this.registerForm.value);
  }
}
