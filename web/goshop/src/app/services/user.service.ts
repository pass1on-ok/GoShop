import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { User } from '../models/user.model';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private apiUrl = 'http://localhost:8080/api';

  constructor(private http: HttpClient) {}

  getUserProfile(): Observable<User | undefined> {
    const token = localStorage.getItem('token');
    if (!token) {
      // Возвращаем пустой Observable, если токен отсутствует
      return of(undefined);
    }

    console.log(token)
    const headers = new HttpHeaders().set('Authorization', 'Bearer ' + token);
    return this.http.get<User>(`${this.apiUrl}/profile`, { headers });
  }
}
