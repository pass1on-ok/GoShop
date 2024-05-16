// auth.service.ts

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private baseUrl = 'http://localhost:8080/api'; 

  constructor(private http: HttpClient) { }

  login(email: string, password: string): Observable<any> {
    return this.http.post<any>(`${this.baseUrl}/login`, { email, password }).pipe(
      tap(res => {
        
        localStorage.setItem('token', res.token);
      })
    );
  }

  logout(): void {
    
    localStorage.removeItem('token');
  }

  register(username: string, email: string, password: string): Observable<any> {
    return this.http.post<any>(`${this.baseUrl}/register`, { username, email, password });
  }
}
