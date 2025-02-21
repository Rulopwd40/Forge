import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { API_URL } from '../environment';
import { User } from '../models/user';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http : HttpClient) { }

  registerUser(user: User): Observable<any> {
    return this.http.post(`${API_URL}/user/register`, user);
  }

  getUserData(username: string): Observable<any>{
    return this.http.get(`${API_URL}/user?username=${username}`)
  }

}
