import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { catchError, map } from 'rxjs/operators';
import { API_URL } from '../environment';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient) { }

  isAuthenticated(): Observable<boolean> {
    return this.http.get(`${API_URL}/profile`, { withCredentials: true }).pipe(
      map(() => true), 
      catchError(() => of(false)) 
    );
  }

  login(credentials: { username: string,email: string; password: string }) {
    return this.http.post(`${API_URL}/login`,credentials);
  }
}
